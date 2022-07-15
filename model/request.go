package model

type UserReq struct {
	Uid    string `json:"uid"`
	Email  string `json:"email"`
	Passwd string `json:"passwd"`
	Token  string `json:"token"`
}
