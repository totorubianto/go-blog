package types

type Register struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	Email    string `form:"email" json:"email" xml:"email" binding:"required"`
	Fullname string `form:"fullName" json:"fullName" xml:"fullName" binding:"required"`
}
