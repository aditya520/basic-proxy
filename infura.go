package main

import (
	"encoding/json"
	"fmt"
	"math/big"

	"log"
)

type InfuraClient struct {
	URL    string
	apiKey string
}

var _ EthereumClient = &InfuraClient{}

func NewInfuraClient(URL string, apiKey string) *InfuraClient {
	return &InfuraClient{
		URL:    URL,
		apiKey: apiKey,
	}
}

func (ic *InfuraClient) GetBalance(address string) (*big.Int, error) {
	url := ic.URL + ic.apiKey
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
