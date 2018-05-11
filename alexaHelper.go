package main

import (
	"context"
	"log"
)

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

type AlexaResponse struct {
	// This is the structure for the response object
	Version  string `json:"version"`
	Response struct {
		OutputSpeech struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"outputSpeech"`
	} `json:"response"`
}

func CreateResponse() *AlexaResponse {
	var resp AlexaResponse
	resp.Version = "1.0"
	resp.Response.OutputSpeech.Type = "PlainText"
	resp.Response.OutputSpeech.Text = "Hello.  Please override this default output."
	return &resp
}

func (resp *AlexaResponse) Say(text string) {
	resp.Response.OutputSpeech.Text = text
}

func HandleRequest(ctx context.Context, i GoTradeMeRequestStruct) (AlexaResponse, error) {

	// Create a response object
	resp := CreateResponse()

	// Customize the response for each Alexa Intent
	switch i.Request.Intent.Name {
	case "officetemp":
		resp.Say("The current temperature is 68 degrees.")
	case "getPrice":
		if len(i.Request.Intent.Slots.StockVals.Resolutions.ResolutionsPerAuthority) == 0 {
			resp.Say("There is an issue")
		} else {
			idVal := string(i.Request.Intent.Slots.StockVals.Resolutions.ResolutionsPerAuthority[0].Values[0].Value.ID)
			stockName := string(i.Request.Intent.Slots.StockVals.Value)
			log.Printf("Request type is ", i.Request.Intent.Name)
			log.Printf("Request slot is ", stockName)
			log.Printf("Request ID is ", idVal)

			valTest := createURL(idVal)
			// This is the url
			apiResponse := grabSite(valTest)
			// This contains the actual network response
		    // Prints out the time
			// Prints out the time zone
			if len(apiResponse) < 2000 {
				resp.Say("Error")
			} else {
		    stockPrice := extractPrice(apiResponse)
		    // This is a string that contains the stock price
		    responseVal := stockName
		    responseVal += " is trading at "
		    responseVal += stockPrice
		    log.Printf(responseVal)
			resp.Say(responseVal)}
		}
	case "AMAZON.HelpIntent":
		resp.Say("This app is easy to use, just say: ask the office how warm it is")
	default:
		resp.Say("I'm sorry, the input does not look like something I understand.")
	}

	return *resp, nil
}

