package server

import  (
    "context"
    "fmt"
    "log"
    "net"
    "os"
    "github.com/myk4040okothogodo/bikerenting/db"
    renteesv1 "github.com/myk4040okothogodo/bikerenting/gen/go/proto/rentees"
    "github.com/arangodb/go-driver"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"

)


const (
    renteesCollectionName = "rentees"
    defaultPort           = "60002"
)


type Server struct {
    database           driver.Database
    renteesCollection  driver.Collection
}

func NewServer(ctx context.Context, database driver.Database)(*Server, error){
    collection, err := db.AttachCollection(ctx, database, renteesCollectionName)
    if err != nil {
        return nil, err
    }

    return &Server {
        database: database,
        renteesCollection: collection,
    }, nil
}

func (s *Server) Run() {
    port := os.Getenv("APP_PORT")
    if port == "" {
        port = defaultPort
    }
    listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0: %s", port))
    if err != nil {
        log.Print("net.Listen failed")
        return
    }
    grpcServer := grpc.NewServer()
    renteesv1. RegisterRenteesAPIServer(grpcServer, s) // use autogenerated code to register the server
    reflection.Register(grpcServer)

    log.Printf("Starting Rentees server on port %s", port)
    go func() {
        grpcServer.Serve(listener)
    }()
}

func (s *Server) ListRentees(ctx context.Context, in *renteesv1.ListRenteesRequest)(*renteesv1.ListRenteesResponse, error){
    if in == nil {
        return nil, fmt.Errorf("Request is empty")
    }

    cursor, err := s.database.Query(ctx, db.ListRecords(s.renteesCollection.Name()), nil)
    if err != nil {
        return nil, fmt.Errorf("failed to iterate over documents: %s", err)
    }

    defer cursor.Close()
    allRentees := []*renteesv1.Rentee{}
    for {
        rentee := new(renteesv1.Rentee)
        var meta driver.DocumentMeta
        meta, err := cursor.ReadDocument(ctx, holder)
        if driver.IsNoMoreDocuments(err) {
            break
        } else if err != nil {
            return nil, fmt.Errorf("Failed to read rentee document: %s", err)
        }
        rentee.Id = meta.Key
        allRentees = append(allRentees, rentee)
    }
    return &renteesv1.ListRenteesResponse{Rentees: allRentees}, nil
}


func (s *Server) GetRentee(ctx context.Context, in *renteesv1.GetRenteeRequest)(*renteesv1.GetRenteeResponse, error) {
    if in == nil || in.Id == "" {
        return nil, fmt.Errorf("Rentee id is not provided")
    }

    rentee := new(renteesv1.Rentee)
    meta, err := s.renteesCollection.ReadDocument(ctx, in.Id, rentee)
    if err != nil {
        if driver.IsNotFound(err) {
             err = fmt.Errorf("Rentee with id '%s' not found", in.Id)
        } else {
             err = fmt.Errorf("Failed to get rentee with id '%s':'%s'", in.Id, err)
        }
        return nil, err
    }
    rentee.Id = meta.Key

    return &renteesv1.GetRenteeResponse{Rentee: rentee}, nil
}


func (s *Server) GetRenteeByBikeId(ctx context.Context, in *renteesv1.GetRenteeByBikeIdRequest)(*renteesv1.GetRenteeByBikeIdResponse, error){
    if in == nil || in.Id == "" {
        return nil, fmt.Errorf("Bike id is not provided")
    }

    const queryHolderByBikeId = `
    FOR rentee IN %s
        FOR bikeId IN rentee.held_bikes
           FILTER bikeId == @bikeId
               RETURN rentee`
    query := fmt.Sprintf(queryRenteeByBikeId, renteesCollectionName)
    bindVars := map[string]interface{}{"bikeId": in.Id}

    cursor, err := s.database.Query(ctx, query, bindVars)
    if err != nil {
        return nil, fmt.Errorf("Failed to iterate over rentee documents with query '%s': %s", queryRenteeByBikeId, err)
    }
    defer cursor.Close()

    r := new(renteesv1.Rentee)
    meta, err := cursor.ReadDocument(ctx, r)
    if driver.IsNoMoreDocuments(err){
        return nil, fmt.Errorf("Rentee that held bike with id %s not found: %s", in.Id, err)
    } else if err != nil {
        return nil, fmt.Errorf("Failed to read rentee document: %s", err)
    }
    r.Id = meta.Key
    return &renteesv1.GetRenteeByBikeIdResponse{Rentee: r}, nil
}


func (s *Server) GetRenteesByBikeTYPE(ctx context.Context, in *renteesv1.GetRenteesByBikeTYPERequest)(*renteesv1.GetRenteesByBikeTYPEResponse, error){
    if in == nil || in.Type == " " {
        return nil, fmt.Errorf("Request is empty")
    }

    const queryRenteeByBikeTYPE = `
    FOR rentee IN %s
        FOR bikeType IN rentee.held_bikes
           FILTER bikeType == @type
               RETURN rentee`
    query := fmt.Sprintf(queryRenteeByBikeTYPE, renteesCollectionName)
    bindVars := map[string]interface{}{"bikeType": in.Type}

    cursor, err := s.database.Query(ctx, query, bindVars)
    if err != nil {
        return nil, fmt.Errorf("Failed to iterate over rentee documents with query '%s': %s", queryRenteeByBikeTYPE, err)
    }
    defer cursor.Close()

    rentees := []*renteesv1.Rentee{}

    for {
        rentee := new(renteesv1.Rentee)
        meta, err := cursor.ReadDocument(ctx, rentee)
        if driver.IsNoMoreDocuments(err){
            break
        } else if err != nil {
            log.Print(err)
            return nil, fmt.Errorf("failed to read rentees document: %s", err)
        }
        rentee.Id = meta.Key
        rentees = append(rentees, rentee)
    }
    return &renteesv1.GetRenteesByBikeTYPEResponse{Rentees: rentees}, nil


func (s *Server) GetRenteesByBikeMAKE(ctx context.Context, in *renteesv1.GetRenteesByBikeMAKERequest)(*renteesv1.GetRenteesByBikeMAKEResponse, error){
    if in == nil || in.Make == " " {
        return nil, fmt.Errorf("Request is empty")
    }

    const queryHolderByBikeMAKE = `
    FOR rentee IN %s
        FOR bikeMake IN rentee.held_bikes
           FILTER bikeMake == @make
               RETURN rentee`
    query := fmt.Sprintf(queryRenteeByBikeMAKE, renteesCollectionName)
    bindVars := map[string]interface{}{"bikeMake": in.Make}

    cursor, err := s.database.Query(ctx, query, bindVars)
    if err != nil {
        return nil, fmt.Errorf("Failed to iterate over rentee documents with query '%s': %s", queryRenteeByBikeMAKE, err)
    }
    defer cursor.Close()

    rentees := []*renteesv1.Rentee{}

    for {
        rentee := new(renteesv1.Rentee)
        meta, err := cursor.ReadDocument(ctx, rentee)
        if driver.IsNoMoreDocuments(err){
            break
        } else if err != nil {
            log.Print(err)
            return nil, fmt.Errorf("failed to read rentees document: %s", err)
        }
        rentee.Id = meta.Key
        rentees = append(rentees, rentee)
    }
    return &renteesv1.GetRenteesByBikeMAKEResponse{Rentees: rentees}, nil


func (s *Server) GetRenteesByBikeOWNER(ctx context.Context, in *renteesv1.GetRenteesByBikeOWNERRequest)(*renteesv1.GetRenteesByBikeOWNERResponse, error){
    if in == nil || in.OwnerName == " " {
        return nil, fmt.Errorf("Request is empty")
    }

    const queryRenteeByBikeOWNER = `
    FOR rentee IN %s
        FOR bikeOwner IN rentee.held_bikes
           FILTER bikeOwner == @owner
               RETURN rentee`
    query := fmt.Sprintf(queryRenteeByBikeOWNER, renteesCollectionName)
    bindVars := map[string]interface{}{"bikeOwner": in.OwnerName}

    cursor, err := s.database.Query(ctx, query, bindVars)
    if err != nil {
        return nil, fmt.Errorf("Failed to iterate over rentee documents with query '%s': %s", queryRenteeByBikeOWNER, err)
    }
    defer cursor.Close()

    rentees := []*renteesv1.Rentee{}

    for {
        rentee := new(renteesv1.Rentee)
        meta, err := cursor.ReadDocument(ctx, rentee)
        if driver.IsNoMoreDocuments(err){
            break
        } else if err != nil {
            log.Print(err)
            return nil, fmt.Errorf("failed to read rentees document: %s", err)
        }
        rentee.Id = meta.Key
        rentees = append(rentees, rentee)
    }
    return &renteesv1.GetRenteesByBikeOWNERResponse{Rentees: rentees}, nil


func (s *Server) AddRentee(ctx  context.Context, in *renteesv1.AddRenteeRequest) (*renteesv1.AddRenteeResponse, error) {
    if in == nil || in.Rentee == nil {
        return nil, fmt.Errorf("Rentee is not provided")
    }

    meta, err := s.renteesCollection.CreateDocument(ctx, in.Rentee)
    if err != nil {
        return nil, fmt.Errorf("Failed to create rentee: %s", err)
    }

    in.Rentee.Id = meta.Key
    return &renteesv1.AddRenteeResponse{Rentee: in.Rentee}, nil
}


func (s *Server) UpdateRentee(ctx context.Context, in *renteesv1.UpdateRenteeRequest)(*renteesv1.UpdateRenteeResponse, error){
    if in == nil || in.Rentee == nil || in.Rentee.Id == "" {
          return nil, fmt.Errorf("Existing rentee is provided")
    }

    _, err := s.renteesCollection.ReplaceDocument(ctx, in.Rentee.Id, in.Rentee)
    if err != nil {
        return nil, fmt.Errorf("Failed to update with id %s", in.Rentee.Id, err)
    }
    return &renteesv1.UpdateRenteeResponse{Rentee: in.Rentee}, nil
}
