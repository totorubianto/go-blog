package users

type UsersSchema struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
	Email    string `bson:"email"`
	Fullname string `bson:"fullName"`
}
