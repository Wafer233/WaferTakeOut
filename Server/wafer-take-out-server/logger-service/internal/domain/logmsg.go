package domain

import "time"

type LogMsg struct {
	Id      int64     `gorm:"primaryKey;type:autoIncrement;column:id;"`
	Level   string    `gorm:"column:level"`
	Message string    `gorm:"column:message"`
	Time    time.Time `gorm:"column:times"`
	//Fields  map[string]interface{} `json:"fields"`
}

func (LogMsg) TableName() string {
	return "logmsg"
}
