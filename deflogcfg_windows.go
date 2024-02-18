//go:build windows

package gitmon

var DefLogCfg = `{
      "level": "debug",
      "encoding": "json",
      "outputPaths": ["stdout", "c:/gitmon/gitmon.log"],
      "errorOutputPaths": ["stderr"],
      "encoderConfig": {
        "messageKey": "message",
        "levelKey": "level",
        "levelEncoder": "lowercase"
      }
    }`
