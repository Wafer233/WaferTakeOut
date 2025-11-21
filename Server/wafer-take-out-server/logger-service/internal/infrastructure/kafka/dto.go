package kafka

type LogMsg struct {
	Id      int64  `json:"id"`
	Level   string `json:"level"`
	Message string `json:"message"`
	Time    string `json:"time"`
	//Fields  map[string]interface{} `json:"fields"`
}
