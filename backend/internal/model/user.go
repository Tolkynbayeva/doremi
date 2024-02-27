package model

type User struct {
	Id         int    `bson:"_id,omitempty"`
	Username   string `bson:"username,omitempty"`
	Email      string `bson:"email,omitempty"`
	Password   string `bson:"password,omitempty"`
	AuthMethod string
}
