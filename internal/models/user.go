package models

type User struct {
	ID       uint16 `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
	Login    string `json:"login" bson:"login"`
	Password string `json:"password" bson:"password"`
}
