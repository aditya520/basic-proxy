package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Response struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
	Id      int    `json:"id"`
}

func createBalanceRequest(address string) string {
	return fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getBalance","params":["%s", "latest"],"id":1}`, address)
}

func doRequest(url string, body string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Printf("Error creating request: %e", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error doing request: %e", err)
		return nil, err
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading body: %e", err)
		return nil, err
	}

	return responseBody, nil
}
