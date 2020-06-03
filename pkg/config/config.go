package config

// API config
type API struct {
	Port int
}

// MongoDB config
type MongoDB struct {
	Addresses                   string
	Username                    string
	Password                    string
	Database                    string
	NewsItemCollection          string
	FileHistoryCollection       string
	CustomLinkCollection        string
	ControlledKeywordCollection string
	PubHistoryCollection        string
	EnableSharding              bool
}
