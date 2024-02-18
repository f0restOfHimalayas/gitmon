package gitmon

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	var cfg zap.Config
	content, err := os.ReadFile("config.json")
	if err != nil || content == nil {
		content = []byte(DefLogCfg)
	}

	if err := json.Unmarshal(content, &cfg); err != nil {
		panic(err)
	}

	if cfg.OutputPaths == nil || len(cfg.OutputPaths) == 0 {
		content = []byte(DefLogCfg)
	}

	logger = zap.Must(cfg.Build())
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Default().Fatalf("error %v", err)
		}
	}(logger)
}

func Log(message string) {
	logger.Info(message)
}

func Error(err error, message string) {
	logger.Error(fmt.Sprintf("Error: %v, Message: %s", err, message))
}
