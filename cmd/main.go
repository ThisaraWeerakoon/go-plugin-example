package main

import (
	"fmt"
	"log/slog"
	"plugin"

	"github.com/profusion/http-redirect/protocol"
)

func main() {
	fmt.Println("Hello World! Log Plugin Demo")

	// Load plugins directly by path (Can use config.json also)
	simpleLogPluginPath := "simple-log-plugin/plugin.so"
	fancyLogPluginPath := "fancy-log-plugin/plugin.so"

	// Load and use log plugin
	fmt.Println("\n=== Using Log Plugin ===")
	logPlugin := loadPlugin(simpleLogPluginPath)
	if logPlugin != nil {
		logPlugin.PrintMessage("Hello from log plugin!")
	}

	// Load and use fancy plugin
	fmt.Println("\n=== Using Fancy Plugin ===")
	fancyPlugin := loadPlugin(fancyLogPluginPath)
	if fancyPlugin != nil {
		fancyPlugin.PrintMessage("Hello from fancy plugin!")
	}
}

// loadPlugin loads a plugin from the given path and returns the SimplePlugin interface
func loadPlugin(pluginPath string) protocol.LogPlugin {
	// Load the plugin
	p, err := plugin.Open(pluginPath)
	if err != nil {
		slog.Error("Failed to load plugin", "path", pluginPath, "error", err)
		return nil
	}

	// Look up the Plugin symbol
	pluginSymbol, err := p.Lookup("Plugin")
	if err != nil {
		slog.Error("Failed to find Plugin symbol", "path", pluginPath, "error", err)
		return nil
	}

	// Cast to SimplePlugin interface
	simplePlugin, ok := pluginSymbol.(protocol.LogPlugin)
	if !ok {
		slog.Error("Plugin does not implement SimplePlugin interface", "path", pluginPath)
		return nil
	}

	return simplePlugin
}
