module github.com/RadisaMilovcevic/bsep_xws/microservices_demo/api_gateway

go 1.17

replace github.com/RadisaMilovcevic/bsep_xws/microservices_demo/common => ../common

require (
	github.com/RadisaMilovcevic/bsep_xws/microservices_demo/common v0.0.0-20220505153930-930b36706fea
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	google.golang.org/grpc v1.46.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	golang.org/x/sys v0.0.0-20220503163025-988cb79eb6c6 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220505152158-f39f71e6c8f3 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
