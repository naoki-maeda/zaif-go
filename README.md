# zaif-go
Zaif API wrapper for Golang.

Please use at your own risk.

Zaif
https://zaif.jp/

Zaif API Document
https://techbureau-api-document.readthedocs.io/ja/latest/index.html

# Install
```
go get github.com/naoki-maeda/zaif-go
go get github.com/stretchr/testify
```

# Example
```go
const (
	publicEndpoint = "https://api.zaif.jp/api/1"
	privateEndpoint = "https://api.zaif.jp/tapi"
)

publicApi := public.NewApiClient(publicEndpoint)
currency, err := publicApi.Currencies("btc")
if err != nil {
    fmt.Println(err)
}
fmt.Println(currency)

privateApi := private.NewApiClient("API_KEY", "SECRET_KEY", privateEndpoint)
getInfo, err := privateApi.GetInfo()
if err != nil {
    fmt.Println(err)
}
fmt.Println(getInfo)
```

# LICENCE
This software is released under the MIT License, see LICENSE.