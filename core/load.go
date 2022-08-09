package core

import "flag"

const (
	DEFAULT_CONFIG_FILE_NAME = "app.yaml"
	CONFIG                   = "config"
)

//CmdMap  The user set command will save in here.
var CmdMap = make(map[string]interface{})

type Load interface {
	LoadCtx()
}

type ConfigLoad struct {
	CfgPath string
}

func (cfgLoad *ConfigLoad) LoadCtx() {
	path := CmdMap[CONFIG].(string)
	if len(path) > 0 {
		cfgLoad.CfgPath = path
	} else {
		cfgLoad.CfgPath = DEFAULT_CONFIG_FILE_NAME
	}
}

func init() {
	cfgPath := flag.String(CONFIG, "", "description")
	flag.Parse()
	// User through the command to tell application operation model.
	CmdMap[CONFIG] = cfgPath
}
