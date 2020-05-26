package main

import (
	"github.com/weareinlog/Inlog.Graylog.Lib/graylog"
)

func main() {
	graylog.ConfigurationLog("graylog.inlog.in:5144", "INLOG", "0.1.0", "github.com/weareinlog/Inlog.Graylog.Lib/graylog", "development")
}
