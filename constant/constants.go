package constant

// To do: Remove this line when usrmgr is ready
const DefaultUserToken = "ed1605e17374bde6c68864d072c9f5c9"

const DefaultPort = 50051      // Default port for gRPC connection
const HeartBeatInterval = 30   // Default interval for heartbeating
const ClientTimeOut = 60       // Default timeout for client connection
const MicroServiceTimeOut = 10 // Default timeout for internal micro service connection

//DataCenter Metrics
type Metrics struct {
	TotalCPU      int64
	UsedCPU       int64
	TotalMemory   int64
	UsedMemory    int64
	TotalStorage  int64
	UsedStorage   int64
	ImageCount    int64
	EndPointCount int64
	NetworkIO     int64
}
