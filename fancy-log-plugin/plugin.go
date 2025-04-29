package main

import (
	"fmt"

	"github.com/profusion/http-redirect/protocol"
)

type FancyLogPluginStr struct{}

// Compile time check that FancyPluginStr implements protocol.LogPlugin
var _ protocol.LogPlugin = FancyLogPluginStr{}

// PrintMessage implements protocol.LogPlugin
func (p FancyLogPluginStr) PrintMessage(message string) {
	fmt.Printf("✨ FANCY LOG PLUGIN ✨: %s\n", message)
}

var Plugin = FancyLogPluginStr{}

func main() { /*empty because it does nothing*/ }
