package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/lang-devbr/fullcycle/cotacao"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	port int = 8080
)

func Start() {
	http.HandleFunc("/cotacao", getCotacaoHandler)
	fmt.Printf("Server is running at port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func getCotacaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cotacao, err := Get(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = insert(r.Context(), cotacao)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao)
}

func insert(ctx context.Context, c *cotacao.Cotacao) error {
	db, err := gorm.Open(sqlite.Open("cotacao.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&cotacao.Cotacao{})
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*10)
	defer cancel()

	tx := db.WithContext(ctx).Create(&c)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func Get(ctx context.Context) (*cotacao.Cotacao, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cotacao cotacao.Cotacao
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		return nil, err
	}
	return &cotacao, nil
}
