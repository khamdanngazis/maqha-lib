module maqhaa/library/middleware

go 1.19

require (
	github.com/google/uuid v1.6.0
	github.com/sirupsen/logrus v1.9.3
	google.golang.org/grpc v1.61.0
	google.golang.org/protobuf v1.32.0
)

replace maqha/library/logging => ../logging

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231106174013-bbf56f31fb17 // indirect
)
