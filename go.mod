module github.com/tk42/golang-echo

go 1.16

require (
	golang.org/x/net v0.7.0
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/tk42/golang-echo/proto/autogen/helloworld => ./proto/autogen/helloworld
