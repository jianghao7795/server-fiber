package request

type LoginForm struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" from:"password"`
}

type RegisterUser struct {
	Name         string `json:"name" form:"name"`
	Password     string `json:"password" from:"password"`
	RePassword   string `json:"re_password" from:"re_password"`
	Introduction string `json:"introduction" form:"introduction"`
	Content      string `json:"content" form:"content"`
	Header       string `json:"header" form:"header"`
}
