# zaif-go
Zaif API wrapper for Golang
Test is not completed.Please be careful to use.

# Install
```
go get github.com/naoki-maeda/zaif-go
go get github.com/google/go-querystring
```

# Example
```
publicApi := public.NewApiClient()
currency, err := publicApi.Currencies("btc")
if err != nil {
    fmt.Println(err)
}
fmt.Println(currency)

privateApi := private.NewApiClient("API_KEY", "SECRET_KEY")
getInfo, err := privateApi.GetInfo()
if err != nil {
    fmt.Println(err)
}
fmt.Println(getInfo)
```

# LICENCE
This software is released under the MIT License, see LICENSE.