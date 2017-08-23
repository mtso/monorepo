package gateway

type Config struct {
	Routes Routes `json:"routes"`
}

var ConfigDir string
