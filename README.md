## **Go: criando uma API Rest**

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

## Interpoladores de Formatação no Go

A tabela abaixo apresenta os principais interpoladores utilizados em Go para formatar strings com `fmt.Sprintf` e outras funções de formatação:

| Interpolador | Descrição                                      |
|-------------|-----------------------------------------------|
| `%d`        | Inteiros decimais                            |
| `%f`        | Números de ponto flutuante                   |
| `%s`        | Strings                                      |
| `%q`        | Strings com aspas                           |
| `%x` e `%X` | Hexadecimal (minúsculo e maiúsculo)         |
| `%b`        | Binário                                     |
| `%p`        | Ponteiros                                   |
| `%t`        | Booleanos                                   |
| `%v`        | Valor padrão                                |
| `%+v`       | Valor padrão com campos                     |

💡 **Exemplo de uso em Go:**
```go
package main

import "fmt"

func main() {
    nome := "João"
    idade := 30
    preco := 99.99
    fmt.Printf("Nome: %s, Idade: %d, Preço: %.2f\n", nome, idade, preco)
}
```
### Conteúdo para leitura

## [1. A Linguagem de Programação Go](https://go.dev/tour/welcome/1)

Uma introdução à linguagem de programação Go pelos criadores, detalhando as características e motivações por trás da linguagem.

## [2. Go no Google: Design de Linguagem a Serviço da Engenharia de Software](https://go.dev/talks/2012/splash.article)

Um artigo de Rob Pike discutindo as escolhas de design e práticas de engenharia por trás do desenvolvimento do Go no Google.

## [3. Padrões de Concorrência em Go](https://go.dev/blog/waza-talk)

Uma palestra de Rob Pike sobre padrões de concorrência em Go, discutindo como Go lida com concorrência e as filosofias de design por trás disso.

## [4. Sintaxe de Declaração do Go](https://go.dev/blog/declaration-syntax)

Um artigo de Rob Pike explicando o raciocínio por trás da sintaxe de declaração do Go e como isso melhora a legibilidade e a manutenção do código.

## [5. A Evolução do Go](https://go.dev/blog/randv2)

Um artigo discutindo as mudanças e melhorias introduzidas na biblioteca Go Standard com math/rand/v2.

## [6. O Caminho para o Go 2](https://go.dev/blog/go2-here-we-come)

Uma postagem no blog de Russ Cox delineando o roadmap e os objetivos para o Go 2, focando em possíveis mudanças e melhorias na linguagem.

## [7. Erros são Valores](https://go.dev/blog/errors-are-values)

Um artigo de Rob Pike explicando a filosofia de tratamento de erros do Go e como tratar erros como valores leva a um melhor design de código.

## [8. Estruturas de Dados em Go: Interfaces](https://research.swtch.com/interfaces)

Um mergulho profundo de Russ Cox nos tipos de interface do Go, explicando como eles funcionam e seu papel no sistema de tipos do Go.
