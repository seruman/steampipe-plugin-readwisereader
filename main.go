package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	"github.com/seruman/steampipe-plugin-readwisereader/readwisereader"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: readwisereader.Plugin})
}
