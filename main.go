// This is my first Go program - Please don't judge it too harshly xD

package main


import (
	// This is similar to package imports in Python
	// "encoding/json"
	// Allows you to decode json from the REST api call
	"fmt"
	// This allows you to input/output values
	// "log"
	// Logging info
	"net/http"
	// Allows get and post requests to be made
	"io/ioutil"
	// Adds the ability to interact with file IO
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


func main() {
	ticker := "AAPL"
	valTest := createURL(ticker)
	fmt.Println(valTest)
	fmt.Println(grabSite(valTest))
	return
}
