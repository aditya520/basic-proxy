# Ethereum Proxy Service



## System Diagram


## Key Components


### 1. ECS (Elastic Container Service)

Elast Container Service to deploy and manage application
### 2. Redis For Cache

### 3. Virtual Private networks and subnets

### 4. Content Delivery Network like Cloud Front

### 5. IAM Roles and Policies

## Data Flow

1. The user sends the request to Amazon Domain Route 53, which sends the request to the nearest CloudFront point.

2. It will check the cache, if it's fresh, it will revert the cached data. If it misses, it will send the request to the load balancer.

3. Elastic Load Balancer will distribute the container ensuring even load distribution.

4. The server in the container will process the request and update the cache with the latest block.


## Deployment

Describe the deployment strategy for the architecture, including any cloud providers or infrastructure used.


## Security

IAM Roles and Policies.

SSL Encryption

