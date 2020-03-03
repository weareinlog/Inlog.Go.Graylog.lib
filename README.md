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