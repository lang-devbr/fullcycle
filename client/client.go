package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/lang-devbr/fullcycle/cotacao"
)

func ProcessarCotacao() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var c cotacao.Cotacao
	err = json.Unmarshal(body, &c)
	if err != nil {
		panic(err)
	}

	err = salvarArquivo(c.Bid)
	if err != nil {
		panic(err)
	}
}

func salvarArquivo(bid string) error {
	f, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if os.IsNotExist(err) {
		f, err = os.Create("cotacao.txt")
		if err != nil {
			return err
		}
	}
	defer f.Close()

	data := []byte(fmt.Sprintf("Dólar: %s\n", bid))
	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}
