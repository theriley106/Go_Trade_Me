// This is my first Go program - Please don't judge it too harshly xD

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
)

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
        return string(err.Error())
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


func main() {
	ticker := "AAPL"
	// Stock ticker that the price will return
	valTest := createURL(ticker)
	// This is the url
	apiResponse := grabSite(valTest)
	// This contains the actual network response

    var stockInfo apiStruct
	err := json.Unmarshal([]byte(apiResponse), &stockInfo)
	if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Printf("%+v", stockInfo)
    c := extractPrice(apiResponse)
    fmt.Println("%q\n", c)
	//
	return
}
