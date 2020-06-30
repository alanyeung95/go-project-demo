package items

type Item struct {
	ID          string      `json:"id"    bson:"_id"`
	Name        string      `json:"name"    bson:"name"`
	Components  []Component `json:"components"    bson:"components"`
	Price       int         `json:"price"    bson:"price"`
	Description string      `json:"description" bson:"description"`
	ModelNumber string      `json:"modelNumber" bson:"modelNumber"`
}

type Component struct {
	Name        string `json:"name"    bson:"name"`
	Description string `json:"description"    bson:"description"`
}
