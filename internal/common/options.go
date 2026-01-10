package commons

import (
	"proxy-adapter/config"

	"go.uber.org/zap"
)

// Options common option for all object that needed
type Options struct {
	Config config.ConfigObject
	Logger *zap.Logger
	Errors []error
}

func InitCommonOptions(options ...func(*Options)) *Options {
	opt := &Options{}
	for _, o := range options {
		o(opt)
		if opt.Errors != nil {
			return opt
		}
	}
	return opt
}

func WithConfig(cfg config.ConfigObject) func(*Options) {
	return func(opt *Options) {
		opt.Config = cfg
	}
}

func WithLogger(logger *zap.Logger) func(*Options) {
	return func(opt *Options) {
		opt.Logger = logger
	}
}
