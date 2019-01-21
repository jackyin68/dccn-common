package cli

type ReplyStatus int32

// CLI reply status
const (
	Success ReplyStatus = iota
	Failure
)

var replyStatusString = []string{
	"available",
	"unavailable",
}

func (r ReplyStatus) String() string {
	return replyStatusString[int(r)]
}

type ErrorReason int32

const (
	DataCenterNotExist ErrorReason = iota
	UserNotExist
	TaskNotExist
	UserNotOwn
	UpdateFailed
)

var errorReasonString = []string{
	"DataCenter does not exist",
	"User does not exist",
	"Task does not exist",
	"User does not own this task",
	"Task can not be updated",
}

func (e ErrorReason) Error() string {
	return errorReasonString[int(e)]
}
