package handlers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"server-api/internal/entity"
	"server-api/internal/infra/database"
	"time"
)

type AccountingHandler struct {
	Database database.AccountingInterface
}

func NewAccountingHandler(db database.AccountingInterface) *AccountingHandler {
	return &AccountingHandler{
		Database: db,
	}
}

func (a *AccountingHandler) GetDollarPrice(w http.ResponseWriter, r *http.Request) {
	//inicio requisição
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Printf("Erro na construção da requisição: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Erro ao efetuar a requisição: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Erro ao ler a resposta: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var dollarAccounting entity.DollarAccounting
	err = json.Unmarshal(body, &dollarAccounting)
	if err != nil {
		log.Printf("Erro ao converter o body da resposta: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//fim requisição

	//inicio persistencia
	ctx = context.Background()
	ctx, cancel = context.WithTimeout(ctx, time.Millisecond*10)
	defer cancel()
	err = a.Database.Create(ctx, &dollarAccounting.Usdbrl)
	if err != nil {
		log.Printf("Erro ao persistir a cotação: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//fim persistencia

	output := struct {
		Bid string `json:"bid"`
	}{
		dollarAccounting.Usdbrl.Bid,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
