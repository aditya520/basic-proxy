package main

import (
	"encoding/json"
	"fmt"
	"math/big"

	"log"
)

type TenderlyClient struct {
	URL    string
	apiKey string
}

var _ EthereumClient = &TenderlyClient{}

func NewTenderlyClient(URL string, apiKey string) *TenderlyClient {
	return &TenderlyClient{
		URL:    URL,
		apiKey: apiKey,
	}
}

func (t *TenderlyClient) GetBalance(address string) (*big.Int, error) {
	url := t.URL + t.apiKey
	fmt.Println(url)
	body := createBalanceRequest(address)
	response, err := doRequest(url, body)
	if err != nil {
		log.Printf("Error doing request: %e", err)
		return nil, err
	}

	var resp Response
	json.Unmarshal(response, &resp)
	log.Printf("Response: %v", resp)

	balance := new(big.Int)
	balance, ok := balance.SetString(resp.Result[2:], 16)
	if !ok {
		log.Printf("Error converting balance: %e", err)
		return nil, err
	}

	return balance, nil
}
