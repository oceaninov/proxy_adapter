package appcontext

import "proxy-adapter/config"

type AppContext struct {
	config config.ConfigObject
}

// NewAppContext initiate appcontext object
func NewAppContext(config config.ConfigObject) *AppContext {
	return &AppContext{
		config: config,
	}
}
