// This is my first Go program - Please don't judge it too harshly xD

// go get github.com/bmuller/arrow/lib

package main


import (
	// This is similar to package imports in Python
	"encoding/json"
	// Allows you to decode json from the REST api call
	"fmt"
	// This allows you to input/output values
	// "log"
	// Logging info
	"net/http"
	// Allows get and post requests to be made
	"io/ioutil"
	// Adds the ability to interact with file IO
	"regexp"
	// Implements the ability to use regex
	// "time"
	// This allows you to interact with datetime
	"github.com/aws/aws-lambda-go/lambda"
	// This is for interaction through lambda
	// IDK why this needs to be defined in main.go instead of alexaHelper.go?
	// ^ I guess because of the line in main()?
	"log"
	// For logging with lambda
	"context"
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

type apiStruct struct {
		// Golang strucutre surrounding the api response
        MetaData struct {
        // Golang json structuring is weird...
		Symbol          string `json:"2. Symbol"`
		// .Symbol returns the stock ticker
		Refreshed string `json:"3. Last Refreshed"`
		// .Refreshed returns the last refresh time
		TimeZone string `json:"6. Time Zone"`
		// .Refreshed returns the last refresh time
		} `json:"Meta Data"`
		// NO SPACE BETWEEN COLON WHEN SPACE IN JSON KEY
    }

func createURL(tickerVal string) string {
	// There needs to be two string because you have to define return type
	var urlVal string
	urlVal = "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol="
	urlVal += tickerVal
	urlVal += "&interval=1min&apikey=3ZC52BSRXYLK76YY"
	return urlVal
}

func grabSite(urlVal string) string {
	// This grabs the site - similar to requests module in python
	response, err := http.Get(urlVal)
	// This gets the response
	if err != nil {
		// This means there was an error
        return string(urlVal + " Error")
        // Returns the error string
    } else {
    	// There was not an error
        data, _ := ioutil.ReadAll(response.Body)
        // Reads the response of the http get function
        return string(data)
        // Returns the response
    }
}

func extractPrice(apiResponse string) string {
	// This extracts the most recent trading price from the api call
	re := regexp.MustCompile("open\\S:\\s.((\\d+).(\\d+))")
	// Matches all opening prices using Regex
	a := re.FindAllString(apiResponse, -1)
	// Extracts all opening prices from the string
	re = regexp.MustCompile("[0-9]+")
	// Extracts digits from the opening prices
	a = re.FindAllString(a[0], -1)
	// Finds all strings that contain only digits
	output := fmt.Sprintf("$%s.%s", a[0], a[1][:2])
	// Sets output to the matched strings in proper $x.xx format
	return output
}

func extractRefresh(apiResponse string) string {
	// This extracts the refresh date from the API Call
	var stockInfo apiStruct
	// Defines a variable called stockInfo that uses that apiStruct
	err := json.Unmarshal([]byte(apiResponse), &stockInfo)
	// Unmarshall basically decodes the json
	if err != nil {
		// There was an error
        fmt.Println("error:", err)
    }
    return string(stockInfo.MetaData.Refreshed)
}

func extractTimeZone(apiResponse string) string {
	// This extracts the refresh date from the API Call
	var stockInfo apiStruct
	// Defines a variable called stockInfo that uses that apiStruct
	err := json.Unmarshal([]byte(apiResponse), &stockInfo)
	// Unmarshall basically decodes the json
	if err != nil {
		// There was an error
        fmt.Println("error:", err)
    }
    return string(stockInfo.MetaData.TimeZone)
}

func generateResponse(tickerVal string, priceVal string) string {
	// This generates a string that contains speech a human can interperet
	responseVal := tickerVal + " is currently trading at " + priceVal
	return responseVal
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
				resp.Say("It looks like there was an error with this stock quote.  Please try again later")
			} else {
		    stockPrice := extractPrice(apiResponse)
		    // This is a string that contains the stock price
		    responseVal := stockName + " is currently trading at " + stockPrice
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

func main() {
	lambda.Start(HandleRequest)
}

