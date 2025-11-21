package domain

type LogMsgRepository interface {
	Save(*LogMsg) error
}
