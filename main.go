package main

import (
	"io"
	"os"

	"github.com/weareinlog/Inlog.Go.Graylog.lib/graylog"
)

func main() {

	graylog.ConfigurationLog("graylog.inlog.in:5144", "INLOG", "0.1.0", "github.com/weareinlog/Inlog.Go.Graylog.lib/graylog")
	aaa, _ := graylog.NewUDPWriter("", map[string]interface{}{})
	graylog.SetOutput(io.MultiWriter(os.Stderr, aaa))
	graylog.Information()

}
