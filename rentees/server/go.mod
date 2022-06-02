module github.com/myk4040okothogodo/bikerenting/rentees/server

go 1.18

replace github.com/myk4040okothogodo/bikerenting/db => ../../db

replace github.com/myk4040okothogodo/bikerenting/gen/go/proto/rentees => ../../gen/go/proto/rentees

require (
	github.com/arangodb/go-driver v1.3.2
	github.com/myk4040okothogodo/bikerenting/db v0.0.0-00010101000000-000000000000
	github.com/myk4040okothogodo/bikerenting/gen/go/proto/rentees v0.0.0-20220602115133-471c7e03d352
	google.golang.org/grpc v1.47.0
)

require (
	github.com/arangodb/go-velocypack v0.0.0-20200318135517-5af53c29c67e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.0.0-20220531201128-c960675eff93 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
