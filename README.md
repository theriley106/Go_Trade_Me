# Go_Trade_Me
Alexa Skill made in Go that outputs current NYSE Stock Prices for a given stock ticker

### This is my first Go program - Please don't judge it too harshly xD

## References

https://github.com/alco/gostart

https://medium.com/@edwardpie/parsing-json-request-body-return-json-response-with-golang-c4f862bbb19b

https://www.youtube.com/watch?v=V-wE4SLZ9q4

https://github.com/benr/alexa_go_prototype/blob/master/alexa_go_prototype.go


## Building

GOOS=linux go build -o main main.go

zip deployment.zip main

## Struct

```go
type AlexaRequest struct {
	// This is the structure for the JSON input
	Version string `json:"version"`
	Request struct {
		Type   string `json:"type"`
		Time   string `json:"timestamp"`
		Intent struct {
			Name               string `json:"name"`
			ConfirmationStatus string `json:"confirmationStatus"`
			Slots              struct {
				StockVals struct {
					Name        string `json:"name"`
					Value       string `json:"value"`
				} `json:"stockVals"`
			} `json:"slots"`
		} `json:"intent"`
	} `json:"request"`
}
```
