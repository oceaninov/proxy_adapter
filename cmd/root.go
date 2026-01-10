package cmd

import (
	"fmt"
	"os"
	"proxy-adapter/config"
	commons "proxy-adapter/internal/common"
	"proxy-adapter/internal/server"
	"proxy-adapter/internal/service"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var levelMapper = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func initLogger(cfg config.ConfigObject) *zap.Logger {
	var level zapcore.Level
	if lvl, ok := levelMapper[cfg.AppLogLevel]; ok {
		level = lvl
	} else {
		level = zapcore.InfoLevel
	}

	loggerCfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.RFC3339NanoTimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger, _ := loggerCfg.Build()
	return logger
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "api",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}

func initService(serviceOption service.Option) *service.Services {
	return &service.Services{
		AzureAD: service.NewAzureADService(serviceOption),
	}
}

func start() {
	cfg := config.Config()
	logger := initLogger(cfg)

	//app := appcontext.NewAppContext(cfg)

	opt := commons.InitCommonOptions(
		commons.WithConfig(cfg),
		commons.WithLogger(logger),
	)

	if len(opt.Errors) > 0 {
		logger.Fatal("Init common options error",
			zap.Any("context", opt.Errors),
		)
		return
	}

	services := initService(service.Option{
		Options: *opt,
	})

	srv := server.NewServer(*opt, services)

	// run app
	srv.StartApp()
}
