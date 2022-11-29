package main

import (
	"net/http"
	"server-api/internal/infra/webserver/handlers"

	"server-api/internal/infra/database"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"server-api/internal/entity"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	db, err := gorm.Open(sqlite.Open("accounting.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.UsdbrlAccounting{})

	accountingDB := database.NewAccounting(db)
	accountingHandler := handlers.NewAccountingHandler(accountingDB)

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/cotacao/", accountingHandler.GetDollarPrice)

	http.ListenAndServe(":8080", r)
}
