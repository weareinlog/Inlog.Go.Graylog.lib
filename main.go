package main

import (
	"io"
	"os"

	"github.com/weareinlog/Inlog.Go.Graylog.lib/graylog"
)

func main() {
	aaa, _ := graylog.NewUDPWriter("", map[string]interface{}{})

	graylog.SetOutput(io.MultiWriter(os.Stderr, aaa))

	graylog.Information()

}
