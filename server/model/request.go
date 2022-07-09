package model

type UserReq struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Passwd string `json:"passwd"`
}
