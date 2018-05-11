package main

import (
	"context"
	"fmt"
	"log"
	"github.com/davecgh/go-spew/spew"
)

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

func HandleRequest(ctx context.Context, i AlexaRequest) (AlexaResponse, error) {
	// Use Spew to output the request for debugging purposes:
	fmt.Println("---- Dumping Input Map: ----")
	spew.Dump(i)
	fmt.Println("---- Done. ----")

	// Example of accessing map value via index:
	log.Printf("Request type is ", i.Request.Intent.Name)
	log.Printf("Request slot is ", i.Request.Intent.Slots.StockVals.Value)

	// Create a response object
	resp := CreateResponse()

	// Customize the response for each Alexa Intent
	switch i.Request.Intent.Name {
	case "officetemp":
		resp.Say("The current temperature is 68 degrees.")
	case "hello":
		resp.Say("Hello there, Lambda appears to be working properly.")
	case "AMAZON.HelpIntent":
		resp.Say("This app is easy to use, just say: ask the office how warm it is")
	default:
		resp.Say("I'm sorry, the input does not look like something I understand.")
	}

	return *resp, nil
}

