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
    booksCollectionName = "Bikes"
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

    booksv1.RegisterBikesAPIServer(grpcServer, s)
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

func (s *Server) GetBikeByISBN(ctx context.Context, in *bikesv1.GetBikeByISBNRequest)(*bikesv1.GetBikeByISBNResponse, error){
    if in == nil || in.Isbn == "" {
        return nil, fmt.Errorf("Bike isbn is not provided")
    }

    const queryBikeByISBN = `
    FOR bike IN %s
        FILTER bike.isbn == @isbn
            RETURN bike`
    query := fmt.Sprintf(queryBikeByISBN, bikesCollectionName)
    bindVars := map[string]interface{}{"isbn": in.Isbn}

    cursor, err := s.database.Query(ctx, query, bindVars)
    if err != nil {
        return nil, fmt.Errorf("Failed to iterate over bike documents with query '%s': %s", queryBikeByISBN, err)
    }
    defer cursor.Close()

    b := new(bikesv1.Bike)
    meta, err := cursor.ReadDocument(ctx, b)
    if driver.IsNoMoreDocuments(err){
        return nil, fmt.Errorf("Bike with ISBN '%s' not found: %s", in.Isbn, err)
    } else if err != nil {
        return nil, fmt.Errorf("Failed to read book document: %s", err)
    }
    b.Id = meta.Key

    return &bikesv1.GetBikeByISBNResponse{Bike: b}, nil
}

func (s *Server) AddBike(ctx context.Context, in *bikesv1.AddBikeRequest)(*bikesv1.AddBikeResponse, error){
    if in == nil || in.Bike == nil {
        return nil, fmt.Errorf("Book is not provided")
    }

    meta, err := s.bikesCollection.CreateDocument(ctx, in.Bike)
    if err != nil {
        return nil, fmt.Errorf("failed to create bike: %s", err)
    }

    in.Book.Id = meta.Key
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