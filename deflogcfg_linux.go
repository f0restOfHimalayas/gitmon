//go:build !windows

package gitmon

var DefLogCfg = `{
      "level": "debug",
      "encoding": "json",
      "outputPaths": ["stdout", "/tmp/gitmon.log"],
      "errorOutputPaths": ["stderr"],
      "encoderConfig": {
        "messageKey": "message",
        "levelKey": "level",
        "levelEncoder": "lowercase"
      }
    }`
