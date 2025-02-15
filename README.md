**Go: criando uma API Rest**

**Comandos.:**

```go mod init pizzaria```

🔍 Explicação detalhada: </br>
✅ Inicializa um projeto Go com módulos </br>
✅ Cria o arquivo go.mod para controle de dependências </br>
✅ Facilita importação de pacotes internos e externos </br>

```go run .\main.go```

🔍 Explicação detalhada: </br>
`go run` → Comando que compila e executa o código Go sem criar um executável. </br>
`.\main.go` → Caminho para o arquivo Go que contém a função main(), ponto de entrada do programa. </br>

```go build .```

🔍 Explicação detalhada:</br>
`go build` → Compila o código-fonte Go.</br>
`.` → Refere-se ao diretório atual, indicando que o Go deve compilar todos os arquivos .go ali presentes.</br>
Se o diretório contiver um arquivo com package main, o comando gerará um executável.</br>

**Var Declaration**
var nomePizzaria string

**Short Assignment Statement**
nomePizzaria := "Pizzaria Go"