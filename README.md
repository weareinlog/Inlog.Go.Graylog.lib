# Inlog.Go.Graylog.lib

# Graylog Lib
- .lib para integração com o Graylog

# Procedimentos
- rodar o comando: go get github.com/weareinlog/Inlog.Go.Graylog.lib
- Adicionar o seguinte código no arquivo main.go

```bash
import "github.com/weareinlog/Inlog.Go.Graylog.lib/graylog"

//VERSION versão do sistema
const VERSION string = "0.2.0"

func init() {
	os.Setenv("VERSION", VERSION)
	os.Setenv("SOLUTION", "Inlog.Go.Services.Bus.AreaAlarme.Back")
}

func configurationGraylog() {
	url := os.Getenv("URL_GRAYLOG")
	company := os.Getenv("COMPANY")
	version := os.Getenv("VERSION")
	solution := os.Getenv("SOLUTION")
	graylog.ConfigurationLog(url, company, version, solution)
}

func main() {
    configurationGraylog()
    ...
```
## Debug
- No arquivo launch.json adicionar as variáveis de ambiente

```bash
"env": {
    "URL_GRAYLOG": "graylog.inlog.in:5144", // url do graylog porta udp
    "COMPANY": "INLOG", // nome do cliente
    "ENVIRONMENT": "development" // ambiente publicado [development | staging | production]
}
```
## Docker
- No arquivo docker-compose.yml adicionar as variáveis de ambiente

```bash
 environment:
    URL_GRAYLOG: "graylog.inlog.in:5144" # url do graylog porta udp
    COMPANY: "INLOG" # nome do cliente
    ENVIRONMENT: "development" # ambiente publicado [development | staging | production]
```

## Exemplo de Uso

```bash
package main

import (
	"os"

	"github.com/weareinlog/Inlog.Go.Graylog.lib/graylog"
)

//VERSION versão
const VERSION string = "0.1.0"

func init() {
	os.Setenv("VERSION", VERSION)
	os.Setenv("SOLUTION", "TESTE")
}

func configurationGraylog() {
	url := os.Getenv("URL_GRAYLOG")
	company := os.Getenv("COMPANY")
	version := os.Getenv("VERSION")
	solution := os.Getenv("SOLUTION")
	graylog.ConfigurationLog(url, company, version, solution)
}

func main() {
	configurationGraylog()
	graylog.Emergency("Emergency")
	graylog.Alert("Alert")
	graylog.Critical("Critical")
	graylog.Error("Error")
	graylog.Warning("Warning")
	graylog.Notice("Notice")
	graylog.Information("Information")
	graylog.Debug("Debug")
}
```
