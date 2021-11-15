package web_app_go

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Create a new instance of the logger. You can have any number of instances.
var logger = logrus.New()

// Level is the just struct.
var logLevelMap = map[string]logrus.Level{
	// TraceLevel is finer than Debug.
	"trace": logrus.TraceLevel,
	// very verbose logging.
	"debug": logrus.DebugLevel,
	"info": logrus.InfoLevel,
	"warn": logrus.WarnLevel,
	"error": logrus.ErrorLevel,
}

type arguments struct {
	LogLevel string
	BindAddress string
	BindPort int
	StaticContents string
}

func runServer(args arguments) error {
	level, ok := logLevelMap[args.LogLevel]
	if !ok {
		return fmt.Errorf("invalid log level: %s", args.LogLevel)
	}
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"args": args,
	}).Info("Given options")

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile(args.StaticContents, false)))
	r.GET("/api/v1/hello", func(c *gin.Context) {
		c.String(200, `{"message": "hello, hello, hello"}`)
	})
	if err := r.Run(fmt.Sprintf("%s:%d", args.BindAddress, args.BindPort)); err != nil {
		return err
	}
	return nil
}
