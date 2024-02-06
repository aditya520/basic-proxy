package main

import (
	"encoding/json"
	"fmt"
	"math/big"

	"log"
)

type ChainStackClient struct {
	URL    string
	apiKey string
}

var _ EthereumClient = &ChainStackClient{}

func NewChainStackClient(URL string, apiKey string) *ChainStackClient {
	return &ChainStackClient{
		URL:    URL,
		apiKey: apiKey,
	}
}

func (c *ChainStackClient) GetBalance(address string) (*big.Int, error) {
	url := c.URL + c.apiKey
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
