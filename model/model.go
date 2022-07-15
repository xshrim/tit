package model

type User struct {
	Uid    string `json:"uid"`
	Name   string `json:"name"`
	Uname  string `json:"uname"`
	Passwd string `json:"passwd"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Dept   string `json:"dept"`
	Group  string `json:"group"`
	Role   string `json:"role"`
}

type Project struct {
	Pid        string `json:"pid"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	Dept       string `json:"dept"`
	Group      string `json:"group"`
	Business   string `json:"business"`
	Level      string `json:"level"`
	User       string `json:"user"`
	Risk       string `json:"risk"`
	Developer  string `json:"developer"`
	Maintainer string `json:"maintainer"`
	Operator   string `json:"operator"`
	Note       string `json:"note"`
}

type Rvitem struct {
	Rid      string `json:"rid"`
	Category string `json:"category"`
	Subject  string `json:"subject"`
	Require  string `json:"require"`
	Level    string `json:"level"`
	Content  string `json:"content"`
	Tips     string `json:"tips"`
	Result   string `json:"result"`
	Note     string `json:"note"`
}
