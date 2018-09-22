package croc

import (
	"time"

	"github.com/schollz/croc/src/logger"
	"github.com/schollz/croc/src/recipient"
	"github.com/schollz/croc/src/relay"
	"github.com/schollz/croc/src/sender"
)

// Croc options
type Croc struct {
	// Options for all
	Debug bool

	// Options for relay
	ServerPort string
	CurveType  string

	// Options for connecting to server
	WebsocketAddress string
	Timeout          time.Duration
	LocalOnly        bool
	NoLocal          bool

	// Options for file transfering
	UseEncryption       bool
	UseCompression      bool
	AllowLocalDiscovery bool
	Yes                 bool
	Stdout              bool

	// Parameters for file transfer
	Filename   string
	Codephrase string

	// private variables

	// localIP address
	localIP string
	// is using local relay
	isLocal      bool
	normalFinish bool
}

// Init will initiate with the default parameters
func Init(debug bool) (c *Croc) {
	c = new(Croc)
	c.ServerPort = "8152"
	c.CurveType = "siec"
	c.UseCompression = true
	c.UseEncryption = true
	c.AllowLocalDiscovery = true
	debugLevel := "info"
	if debug {
		debugLevel = "debug"
		c.Debug = true
	}
	logger.SetLogLevel(debugLevel)
	sender.DebugLevel = debugLevel
	recipient.DebugLevel = debugLevel
	relay.DebugLevel = debugLevel
	return
}
