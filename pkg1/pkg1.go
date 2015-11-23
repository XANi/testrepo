package pkg1

import (
	"fmt"
	"github.com/XANi/testrepo/plugin_api"
)

func PrintVersion() {
	plugin_api.Version()
	fmt.Printf("pkg1: v0.0.2\n")
}
