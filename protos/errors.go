package ankr_default

import "errors"

// List execution errors
var (
	ErrDataCenterNotExist        = errors.New("dataCenter does not exist")
	ErrUserNotExist              = errors.New("token error, can not find user")
	ErrTaskNotExist              = errors.New("task does not exist")
	ErrUserNotOwn                = errors.New("user does not own this task")
	ErrUpdateFailed              = errors.New("task can not be updated")
	ErrUserAlreadyExist          = errors.New("user already existed")
	ErrPasswordError             = errors.New("password does not match")
	ErrHashPassword              = errors.New("hash password failed")
	ErrNamePasswordEmpty         = errors.New("name or Password is empty")
	ErrStatusNotSupportOperation = errors.New("current status not support operation")
	ErrTaskStatusCanNotUpdate    = errors.New("task status is done or cancelled, can not update")
	ErrReplicaTooMany            = errors.New("replica too many")
	ErrUnknown                   = errors.New("unknown operation or code")
	ErrSyncTaskInfo              = errors.New("sync task info error")
	ErrPublish                   = errors.New("mq publish message error")
	ErrConnection                = errors.New("connection error")
	ErrNoAvailableDataCenter     = errors.New("no available data center")
	ErrEmailFormat               = errors.New("email invalid format")
	ErrEmailShouldNotSame        = errors.New("email should not same as before")
	ErrPasswordFormat            = errors.New("password invalid format")
	ErrUserNameFormat            = errors.New("user name invalid format")
	ErrPasswordLength            = errors.New("password must be at least 6 characters long")
	ErrCronJobScheduleFormat     = errors.New("cronjob schedule invalid format")
	ErrPassword                  = errors.New("invalid password")
	ErrPasswordShouldNotSame     = errors.New("password should not same as before")
	ErrEmailExit                 = errors.New("email exist")
	ErrTokenNeedRefresh          = errors.New("token is unavailable, need call refresh token")
	ErrTokenPassedMax            = errors.New("tokens number reaches max limit(10)")
	ErrTokenParseFailed          = errors.New("tokens parse failed")
	ErrRefreshToken              = errors.New("refresh_token error, need login")
	ErrAccessTokenExpired        = errors.New("access_token expired")
	ErrCanceledTwice             = errors.New("can not cancel twice")
	ErrPurgedTwice               = errors.New("can not purge twice")
	ErrAuthNotAllowed            = errors.New("auth not allow")
	ErrUnexpectedChar            = errors.New("unexpected char")
	ErrPasswordSame              = errors.New("password must be not same as before")
	ErrOldPassword               = errors.New("old password does not match")
	ErrEmailSame                 = errors.New("email must be not same as before")
	ErrUserNotVariyEmail         = errors.New("user's email has not been varified, please varify email first")
	ErrUserDeactive              = errors.New("login failed, account has been locked, please contact admin")
	ErrEmailNoExit               = errors.New("email does not exist
)
