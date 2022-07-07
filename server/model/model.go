package model

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Uname    string `json:"uname"`
	Passwd   string `json:"passwd"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Dept     string `json:"dept"`
	Group    string `json:"group"`
	Reviewer string `json:"reviewer"`
}

type Project struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	Dept       string `json:"dept"`
	Group      string `json:"group"`
	Level      string `json:"level"`
	User       string `json:"user"`
	Risk       string `json:"risk"`
	Note       string `json:"note"`
	Developer  string `json:"developer"`
	Maintainer string `json:"maintainer"`
	Operator   string `json:"operator"`
}
