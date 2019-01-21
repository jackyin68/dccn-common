package hub

//go:generate stringer -type TaskStatus
type TaskStatus int32

const (
	New TaskStatus = iota
	Running
	StartFailed
	Updating
	UpdateFailed
	Cancelled
	Done
)

//go:generate stringer -type TaskEvent
type TaskEvent int32

const (
	NewEvent TaskEvent = iota
	UpdateEvent
	CancelEvent
)

//go:generate stringer -type RequestType
type RequestType int32

const (
	AddTask RequestType = iota
	TaskList
	CancelTask
)

type QueueName int32

const (
	TaskManager QueueName = iota
	K8sAdopter
)

var queueNameString = []string{
	"task_manager",
	"k8s_adopter",
}

func (q QueueName) String() string {
	return queueNameString[int(q)]
}
