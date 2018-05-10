// This is my first Go program - Please don't judge it too harshly xD

package main


import (
	// This is similar to package imports in Python
	// "encoding/json"
	// Allows you to decode json from the REST api call
	// "fmt"
	// This allows you to input/output values
	// "log"
	// Logging info
	// "net/http"
	// Allows get and post requests to be made
)

func createURL(tickerVal string) {
	return ("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=" + tickerVal + "&interval=1min&apikey=3ZC52BSRXYLK76YY")
}


func main() {
	return createURL("AAPL")
}
