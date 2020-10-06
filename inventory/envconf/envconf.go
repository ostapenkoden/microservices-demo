package envconf

import (
	"os"
)

type RPC struct {
	Address string
}

type MongoDB struct {
	ConnectionURI string
}

type Config struct {
	RPC     RPC
	MongoDB MongoDB
}

func New() *Config {
	return &Config{
		RPC: RPC{
			Address: os.Getenv("RPC_ADDRESS"),
		},
		MongoDB: MongoDB{
			ConnectionURI: os.Getenv("MONGODB_CONNECTION_URI"),
		},
	}
}
