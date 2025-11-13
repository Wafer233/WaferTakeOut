package application

type LoginVO struct {
	Id     int64  `json:"id"`
	OpenId string `json:"openid"`
	Token  string `json:"token"`
}
