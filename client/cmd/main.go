package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type DollarAccounting struct {
	Bid string `json:"bid"`
}

func main() {
	//inicio requisição
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*300)
	defer cancel()

	url := "http://localhost:8080/cotacao/"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Printf("Erro na construção da requisição: %s\n", err.Error())
		panic(err)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Erro ao efetuar a requisição: %s\n", err.Error())
		panic(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Erro ao ler a resposta: %s\n", err.Error())
		panic(err)
	}

	var dollarAccounting DollarAccounting
	err = json.Unmarshal(body, &dollarAccounting)
	if err != nil {
		log.Printf("Erro ao converter o body da resposta: %s\n", err.Error())
		panic(err)
	}
	//fim requisição

	//inicio arquivo
	f, err := os.Create("cotacao.txt")
	if err != nil {
		log.Printf("Erro ao criar arquivo: %s\n", err.Error())
		panic(err)
	}

	text := fmt.Sprintf("Dólar: %s",dollarAccounting.Bid)
	tamanho, err := f.WriteString(text)
	if err != nil {
		log.Printf("Erro ao escrever no arquivo: %s\n", err.Error())
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
	f.Close()
	//fim arquivo
}
