package main

import (
	"log/slog"

	"github.com/profusion/http-redirect/protocol"
)

type SimpleLogPluginStr struct{}

// Compile time check that PluginStr implements protocol.LogPlugin
var _ protocol.LogPlugin = SimpleLogPluginStr{}

// PrintMessage implements protocol.LogPlugin
func (p SimpleLogPluginStr) PrintMessage(message string) {
	slog.Info("Simple Log Plugin:", "message", message)
}

var Plugin = SimpleLogPluginStr{}

func main() { /*empty because it does nothing*/ }
