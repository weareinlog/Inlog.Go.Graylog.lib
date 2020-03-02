package graylog

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

//ConfigurationLog configuração do Graylog
func ConfigurationLog(url, company, softwareVersion, solution string) error {
	var graylogAddr string
	flag.StringVar(&graylogAddr, "GrayLog", url, "")
	if graylogAddr != "" {
		gelfWriter, err := NewUDPWriter(graylogAddr, map[string]interface{}{
			"Company":  company,
			"_version": softwareVersion,
			"Solution": solution,
		})
		if err != nil {
			return err
		}
		// log to both stderr and graylog2
		SetOutput(io.MultiWriter(os.Stderr, gelfWriter))
		msg := fmt.Sprintf("Graylog configurado: %s", graylogAddr)
		Information(msg)
		return nil
	} else {
		return errors.New("Erro na url do Graylog")
	}
}
