package users

type User struct {
	ID       string `json:"id"    bson:"_id"`
	Name     string `json:"name"    bson:"name"`
	Password string `json:"password"    bson:"password"`
}
