# Go_Trade_Me
Alexa Skill made in Go that outputs current NYSE Stock Prices for a given stock ticker

### This is my first Go program - Please don't judge it too harshly xD


### Building

<b>GOOS=linux go build -o main *.go</b>

<b>zip deployment.zip main</b>

<b>Upload as ZIP in AWS Lambda</b>

or...

<b>./deploy.sh</b>

### Getting Stock Tickers from Company Name

<p align="center">
<img src ="src/IDVals.png">
</p>
<p align="center">Setting the Valid Ticker as an ID Value in the ASK</p>

Each company name is saved as a slot value, and each slot value has an ID that corresponds to the stock ticker for that company.

```go
stockName := i.Request.Intent.Slots.StockVals.Value
// ie: Google, Tesla, Home Depot, etc.
stockTicker := i.Request.Intent.Slots.StockVals.Resolutions.ResolutionsPerAuthority[0].Values[0].Value.ID
// ie: GOOG, TSLA, HD, etc.
```

### Request Structure

```go
type GoTradeMeRequestStruct struct {
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
					Resolutions struct {
						ResolutionsPerAuthority []struct {
							Authority string `json:"authority"`
							Status    struct {
								Code string `json:"code"`
							} `json:"status"`
							Values []struct {
								Value struct {
									Name string `json:"name"`
									ID   string `json:"id"`
								} `json:"value"`
							} `json:"values"`
						} `json:"resolutionsPerAuthority"`
					} `json:"resolutions"`
					ConfirmationStatus string `json:"confirmationStatus"`
				} `json:"stockVals"`
			} `json:"slots"`
		} `json:"intent"`
	} `json:"request"`
}
```

### Response Structure

```go
type AlexaResponse struct {
	// This is the structure for the response object
	Version  string `json:"version"`
	Response struct {
		OutputSpeech struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"outputSpeech"`
		EndSession  bool `json:"shouldEndSession"`
	} `json:"response"`
}
```
