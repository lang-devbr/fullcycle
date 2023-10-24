package client

import (
	"io"
	"net/http"
)

func GetCotacao() {
	c := http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
