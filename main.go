package main

import (
	"io"
	"os"
	"teste/graylog"
)

func main() {
	aaa, _ := graylog.NewUDPWriter("", map[string]interface{}{})

	graylog.SetOutput(io.MultiWriter(os.Stderr, aaa))

	graylog.Information()

}
