# EthProxy

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Description

A simple proxy server to fetch your ethereum balance

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Contact](#contact)

## Installation

Run `go mod download` and build the project using `go build ./...` and run the generated binary

Don't forget to copy `example.env` file to `.env` file with your priate keys

## Usage
Pass the address to the following curl command

```curl --location 'http://0.0.0.0:3333/eth/balance/{ADDRESS}'```


## Improvements and TODO

1. Usage of cache Layer. [cacheStore](https://github.com/goware/cachestore) can easily be integrated.
2. Use of SingleFlight for concurrent similar requests.
3. Implement a priority system for clients. Priorities can be set using average response time, failure frequency, cost of query and manual override.
4. Another way to have consistent data is to query multiple clients to verify response data.
5. In case of inconsistent data, we can chose the client with maximum priority.
6. Implement Unit testing and load testing.


## Contact

If you have any questions or suggestions, feel free to reach out to me at [arora.aditya520@gmail.com](mailto:arora.aditya520@gmail.com).
