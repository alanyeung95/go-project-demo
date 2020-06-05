package config

// API config
type API struct {
	Port int
}

// MongoDB config
type MongoDB struct {
	Addresses      string
	Username       string
	Password       string
	Database       string
	ItemCollection string
	EnableSharding bool
}