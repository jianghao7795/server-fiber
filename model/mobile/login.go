package mobile

type Login struct {
	Username string `json:"username" form:"username" gorm:"column:username;comment:用户名;size:50;"`
	Password string `json:"password" form:"password" gorm:"column:password;comment:密码;size:100;"`
	Realname string `json:"realname" form:"realname" gorm:"column:realname;comment:真实姓名;size:50;"`
}
