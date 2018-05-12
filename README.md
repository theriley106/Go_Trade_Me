# Go_Trade_Me
Golang Alexa Skill that outputs current NYSE Stock Prices for a given stock ticker

## Sample Utterances

### "Open Go Trade Me"

```
"Thanks for checking out Go Trade Me, an open sourced stock price tracker written in Go!"
```

### "How do I use this skill?"

```
"You can ask me for the current trading price of any publicly traded company!"
```

### "What is Tesla currently trading at?"

```
"Tesla is currently trading at $301.13"
```

### "What is the trading price of Western Digital?"

```
"Western Digital is currently trading at $78.77"
```

### "Tell me about the Developer"

```
"Created in May 2018 by Christopher Lambert."
"This alexa skill is completely open sourced."
"Please check out the skill on Github or contact me for more information."
```


## Getting Stock Tickers from Company Name

Each company name is saved as a slot value, and each slot value has an ID that corresponds to the stock ticker for that company.

```go
stockName := i.Request.Intent.Slots.StockVals.Value
// ie: Google, Tesla, Home Depot, etc.
stockTicker := i.Request.Intent.Slots.StockVals.Resolutions.ResolutionsPerAuthority[0].Values[0].Value.ID
// ie: GOOG, TSLA, HD, etc.
```

## Request Structure

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

## Response Structure

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

### Slot Value Format

```javascript
{
    "id": "WYN",
"name": {
    "value": "Wyndham Worldwide"
    }
},
{
    "id": "WYNN",
"name": {
    "value": "Wynn Resorts Ltd"
    }
},
{
    "id": "XEL",
"name": {
    "value": "Xcel Energy"
    }
},
{
    "id": "XRX",
"name": {
    "value": "Xerox Corp"
    }
}
```

<b>or</b>

<p>
<img src ="src/IDVals.png">
</p>
<p>View from the Alexa Skill Kit</p>

## Building

```console
foo@bar:~$ GOOS=linux go build -o main *.go
foo@bar:~$ zip deployment.zip main
```
or...

```console
foo@bar:~$ ./deploy.sh
```

After building, upload "deployment.zip" to AWS Lambda
