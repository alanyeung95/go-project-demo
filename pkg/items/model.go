package items

type Item struct {
	ID         string      `json:"id"    bson:"id"`
	Name       string      `json:"name"    bson:"name"`
	Components []Component `json:"components"    bson:"components"`
	Price      int         `json:"price"    bson:"price"`
}

type Component struct {
	Name string `json:"name"    bson:"name"`
}
