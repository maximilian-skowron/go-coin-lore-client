# Go coinlore client

A very simple client to access the coinlore api.
https://www.coinlore.com/cryptocurrency-data-api

## Usage

Create a client object:
```go
import "coinlore"

client := coinlore.NewClient(coinlore.BaseURL)
```

Use the client to query the api with the provided functions.

```go
cm, _ := client.GetCryptoMarket(context.TODO())
```