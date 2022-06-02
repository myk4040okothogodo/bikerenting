package server


import (
    "context"
    "fmt"
    "log"
    "net"
    "os"
    "github.com/myk4040okothogodo/bikerenting/db"
    bikesv1 "github.com/myk4040okothogodo/bikerenting/gen/go/proto/bikes"
    "github.com/arangodb/go-driver"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

const (
    bikesCollectionName = "Bikes"
    defaultPort         = "60001"
)

type Server struct {
    database             driver.Database
    bikesCollection      driver.Collection
}

func NewServer(ctx context.Context, database driver.Database)(*Server, error) {
    collection, err := db.AttachCollection(ctx, database, bikesCollectionName)
    if err != nil {
        return nil, err
    }

    _, _, err = collection.EnsurePersistentIndex(ctx, []string{"serial"}, &driver.EnsurePersistentIndexOptions{Unique: true})
    if err != nil {
        return nil, err
    }

    return &Server{
        database:        database,
        bikesCollection: collection,
    }, nil
}


func (s *Server) Run() {
    port := os.Getenv("APP_PORT")
    if port == "" {
        port = defaultPort
    }

    listener,err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
    if err != nil {
        log.Fatal("net.Listen failed")
        return
    }

    grpcServer := grpc.NewServer()

    bikesv1.RegisterBikesAPIServer(grpcServer, s)
    reflection.Register(grpcServer)

    log.Printf("Starting Rental Bikes server on port %s", port)

    go func() {
        grpcServer.Serve(listener)
    }()
}


func (s *Server) ListBikes(ctx context.Context, in *bikesv1.ListBikesRequest)(*bikesv1.ListBikesResponse, error){
    if in == nil {
        return nil,  fmt.Errorf("Request is empty")
    }

    cursor, err := s.database.Query(ctx, db.ListRecords(s.bikesCollection.Name()), nil)
    if err != nil {
        return nil, fmt.Errorf("Failed to iterate over documents: %s", err)
    }
    defer cursor.Close()

    allBikes := []*bikesv1.Bike{}
    for {
        bike := new(bikesv1.Bike)
        var meta driver.DocumentMeta
        meta, err := cursor.ReadDocument(ctx, bike)
        if driver.IsNoMoreDocuments(err){
            break
        } else if err != nil {
            return nil, fmt.Errorf("Failed to read bike document: %s", err)
        }
        bike.Id = meta.Key
        allBikes = append(allBikes, bike)
    }
    return &bikesv1.ListBikesResponse{Bikes: allBikes}, nil
}


func (s *Server) GetBike(ctx context.Context, in *bikesv1.GetBikeRequest)(*bikesv1.GetBikeResponse, error){
    if in == nil || in.Id == "" {
        return nil, fmt.Errorf("Bike id is not provided")
    }

    bike := new(bikesv1.Bike)
    meta, err := s.bikesCollection.ReadDocument(ctx, in.Id, bike)
    if err != nil {
        if driver.IsNotFound(err){
            err = fmt.Errorf("Bike with id '%s' not found", in.Id)

        } else {
            err = fmt.Errorf("Failed to get bike with id '%s'", in.Id, err)
        }
        return nil, err
    }
    bike.Id = meta.Key
    return &bikesv1.GetBikeResponse{Bike: bike}, nil
}


func (s *Server) GetBikes(ctx context.Context, in *bikesv1.GetBikesRequest) (*bikesv1.GetBikesResponse, error){
    if in == nil || len(in.Ids) == 0 {
        return nil, fmt.Errorf("Bikes ids are not provided")
    }

    const queryBikesByIds = `
    FOR bike IN %s
        FILTER bike._key in @ids
            RETURN bike`

    query := fmt.Sprintf(queryBikesByIds, bikesCollectionName)
    bindVars := map[string]interface{}{"ids":in.Ids}

    cursor, err := s.database.Query(ctx, query, bindVars)
    if err != nil {
        return nil, fmt.Errorf("Failed to iterate over bike docments with query '%s': %s", queryBikesByIds, err)
    }

    defer cursor.Close()

    bikes := []*bikesv1.Bike{}

    for {
        bike := new(bikesv1.Bike)
        meta, err := cursor.ReadDocument(ctx, bike)
        if driver.IsNoMoreDocuments(err) {
            break
        } else if err != nil {
            log.Print(err)
            return nil, fmt.Errorf("Failed to read bike document: %s", err)
        }

        bike.Id = meta.Key
        bikes = append(bikes, bike)
    }
    return &bikesv1.GetBikesResponse{Bikes: bikes}, nil
}

func (s *Server) GetBikesByTYPE(ctx context.Context, in *bikesv1.GetBikesByTYPERequest)(*bikesv1.GetBikesByTYPEResponse, error){
    if in == nil || in.Type == "" {
        return nil, fmt.Errorf("Bike type is not provided")
    }

    const queryBikeByTYPE = `
    FOR bike IN %s
        FILTER bike.type == @type
            RETURN bike`
    query := fmt.Sprintf(queryBikeByTYPE, bikesCollectionName)
    bindVars := map[string]interface{}{"type": in.Type}

    cursor, err := s.database.Query(ctx, query, bindVars)
    if err != nil {
        return nil, fmt.Errorf("Failed to iterate over bike documents with query '%s': %s", queryBikeByTYPE, err)
    }
    defer cursor.Close()

   bikes := []*bikesv1.Bike{}
   for {
     bike := new(bikesv1.Bike)
     meta, err := cursor.ReadDocument(ctx, bike)
     if driver.IsNoMoreDocuments(err){
       break
     } else if err != nil {
         log.Print(err)
         return nil, fmt.Errorf("failed to read rentees document: %s", err)
     }
     bike.Id = meta.Key
     bikes = append(bikes, bike)
   }
   return &bikesv1.GetBikesByTYPEResponse{Bikes: bikes}, nil
}
    

func (s *Server) GetBikesByOWNER(ctx context.Context, in *bikesv1.GetBikesByOWNERRequest)(*bikesv1.GetBikesByOWNERResponse, error){
    if in == nil || in.OwnerName == "" {
        return nil, fmt.Errorf("Bike owner is not provided")
    }

    const queryBikeByOWNER = `
    FOR bike IN %s
        FILTER bike.owner_name == @ownername
            RETURN bike`
    query := fmt.Sprintf(queryBikeByOWNER, bikesCollectionName)
    bindVars := map[string]interface{}{"owner_name": in.OwnerName}

    cursor, err := s.database.Query(ctx, query, bindVars)
    if err != nil {
        return nil, fmt.Errorf("Failed to iterate over bike documents with query '%s': %s", queryBikeByOWNER, err)
    }
    defer cursor.Close()

     bikes := []*bikesv1.Bike{}
   for {
     bike := new(bikesv1.Bike)
     meta, err := cursor.ReadDocument(ctx, bike)
     if driver.IsNoMoreDocuments(err){
       break
     } else if err != nil {
         log.Print(err)
         return nil, fmt.Errorf("failed to read rentees document: %s", err)
     }
     bike.Id = meta.Key
     bikes = append(bikes, bike)
   }
   return &bikesv1.GetBikesByOWNERResponse{Bikes: bikes}, nil
}


func (s *Server) GetBikesByMAKE(ctx context.Context, in *bikesv1.GetBikesByMAKERequest)(*bikesv1.GetBikesByMAKEResponse, error){
    if in == nil || in.Make == "" {
        return nil, fmt.Errorf("Bike make is not provided")
    }

    const queryBikeByMAKE = `
    FOR bike IN %s
        FILTER bike.make == @make
            RETURN bike`
    query := fmt.Sprintf(queryBikeByMAKE, bikesCollectionName)
    bindVars := map[string]interface{}{"make": in.Make}

    cursor, err := s.database.Query(ctx, query, bindVars)
    if err != nil {
        return nil, fmt.Errorf("Failed to iterate over bike documents with query '%s': %s", queryBikeByMAKE, err)
    }
    defer cursor.Close()

     bikes := []*bikesv1.Bike{}
     for {
     bike := new(bikesv1.Bike)
     meta, err := cursor.ReadDocument(ctx, bike)
     if driver.IsNoMoreDocuments(err){
       break
     } else if err != nil {
         log.Print(err)
         return nil, fmt.Errorf("failed to read rentees document: %s", err)
     }
     bike.Id = meta.Key
     bikes = append(bikes, bike)
   }
   return &bikesv1.GetBikesByMAKEResponse{Bikes: bikes}, nil
}

//func (s *Server) GetBikeBySERIAL(ctx context.Context, in *bikesv1.GetBikesSERIALRequest)(*bikesv1.GetBikesBySERIALResponse, error){
//    if in == nil || in.Serial == "" {
//        return nil, fmt.Errorf("Bike serial is not provided")
//    }

//    const queryBikeBySERIAL = `
//    FOR bike IN %s
//        FILTER bike.serial == @serial
//            RETURN bike`
//    query := fmt.Sprintf(queryBikeBySERIAL, bikesCollectionName)
//    bindVars := map[string]interface{}{"serial": in.Serial}

//    cursor, err := s.database.Query(ctx, query, bindVars)
//    if err != nil {
//        return nil, fmt.Errorf("Failed to iterate over bike documents with query '%s': %s", queryBikeBySERIAL, err)
//    }
//    defer cursor.Close()

//    b := new(bikesv1.Bike)
//    meta, err := cursor.ReadDocument(ctx, b)
//    if driver.IsNoMoreDocuments(err){
//        return nil, fmt.Errorf("Bike with SERIAL '%s' not found: %s", in.Serial, err)
//    } else if err != nil {
//        return nil, fmt.Errorf("Failed to read bike document: %s", err)
//    }
//    b.Id = meta.Key

//    return &bikesv1.GetBikesBySERIALResponse{Bike: b}, nil
//}



func (s *Server) AddBike(ctx context.Context, in *bikesv1.AddBikeRequest)(*bikesv1.AddBikeResponse, error){
    if in == nil || in.Bike == nil {
        return nil, fmt.Errorf("Bike is not provided")
    }

    meta, err := s.bikesCollection.CreateDocument(ctx, in.Bike)
    if err != nil {
        return nil, fmt.Errorf("failed to create bike: %s", err)
    }

    in.Bike.Id = meta.Key
    return &bikesv1.AddBikeResponse{Bike: in.Bike}, nil
}

func (s *Server) DeleteBike(ctx context.Context, in *bikesv1.DeleteBikeRequest)(*bikesv1.DeleteBikeResponse, error) {
    if in == nil || in.Id == "" {
        return nil, fmt.Errorf("Bike id is not provided")
    }

    _, err := s.bikesCollection.RemoveDocument(ctx, in.Id)
    if err != nil {
        return nil, fmt.Errorf("Failed to remove existing bike: %s", err)
    }

    return &bikesv1.DeleteBikeResponse{}, nil
}
