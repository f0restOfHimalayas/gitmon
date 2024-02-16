package gitmon

import (
	"encoding/json"
	"os"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	var cfg zap.Config
	content, err := os.ReadFile("config.json")
	if err != nil || content == nil {
		content = []byte(`{
      "level": "debug",
      "encoding": "json",
      "outputPaths": ["stdout", "/tmp/logs"],
      "errorOutputPaths": ["stderr"],
      "initialFields": {"foo": "bar"},
      "encoderConfig": {
        "messageKey": "message",
        "levelKey": "level",
        "levelEncoder": "lowercase"
      }
    }`)
	}

	if err := json.Unmarshal(content, &cfg); err != nil {
		panic(err)
	}

	logger = zap.Must(cfg.Build())
	defer logger.Sync()
}

func Log(message string) {
	logger.Info(message)
}
