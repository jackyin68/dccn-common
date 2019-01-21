package dc

//go:generate stringer -type taskType
type taskType int32

const (
	NewTask taskType = iota
	UpdateTask
	CancelTask
	ListTask
	HeartBeat
)

//go:generate stringer -type DataCennterTaskStatus
type DataCennterTaskStatus int32

const (
	StartSuccess DataCennterTaskStatus = iota
	StartFailure
	UpdateSuccess
	UpdateFailure
	Done
	Cancelled
)

//go:generate stringer -type DataCenteStatus
type DataCenteStatus int32

const (
	Online DataCenteStatus = iota
	Offline
)
