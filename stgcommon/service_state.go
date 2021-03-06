package stgcommon

// ServiceState: 服务状态枚举
// Author: yintongqiang
// Since:  2017/8/10
type ServiceState int

const (
	CREATE_JUST      ServiceState = iota // Service just created,not start
	RUNNING                              // Service Running
	SHUTDOWN_ALREADY                     // Service shutdown
	START_FAILED                         // Service Start failure
)

func (state ServiceState) String() string {
	switch state {
	case CREATE_JUST:
		return "CREATE_JUST"
	case RUNNING:
		return "RUNNING"
	case SHUTDOWN_ALREADY:
		return "SHUTDOWN_ALREADY"
	case START_FAILED:
		return "START_FAILED"
	default:
		return "Unknow"
	}
}
