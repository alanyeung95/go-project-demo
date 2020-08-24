package users

type User struct {
	ID       string `json:"id"    bson:"_id"`
	Name     string `json:"name"    bson:"name"`
	Email    string `json:"email"    bson:"email"`
	Password string `json:"password"    bson:"password"`
}

type UserLoginParam struct {
	Email    string `json:"email"    bson:"email"`
	Password string `json:"password"    bson:"password"`
}

type LoginResponse struct {
	Status bool   `json:"status"    bson:"status"`
	Token  string `json:"token"    bson:"token"`
}
