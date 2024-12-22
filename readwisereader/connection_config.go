package readwisereader

import "github.com/turbot/steampipe-plugin-sdk/v5/plugin"

type readwisereaderConfig struct {
	Token *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &readwisereaderConfig{}
}

func GetConfig(connection *plugin.Connection) readwisereaderConfig {
	if connection == nil || connection.Config == nil {
		return readwisereaderConfig{}
	}
	config, _ := connection.Config.(readwisereaderConfig)
	return config
}
