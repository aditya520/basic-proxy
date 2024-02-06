# Instagram Application on Ethereum


## Overview
The biggest difference between the traditional Instagram and Instagram like application on Ethtereum would how we store the data on the chain. 

## Components

### 1. Authentication
Authentication can be done by signing in with any smart contract wallet out there and signing in the message to verify the wallet address. 

### 2. Media Storage service.
Best solution would be to store the images to the IPFS, and it returns the fixed length hash which can be stored into a smart contract along with the address map. 

The datatype will look like a `map (address => bytes)`.

We can have an offchain service to store the metadata in a SQL database to ease out the queries. 

### 3. Smart Contract 
The Smart contract will serve as the storage as well as proof of verification/existence service. 



## Conclusion 

This service can easily be scaled using traditional AWS services as it will 75% similar to how instagram works right now.