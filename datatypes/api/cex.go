package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sofiengwin/cryptomasters/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	fmt.Println(fmt.Sprintf(apiUrl, strings.ToUpper(currency)))
	response, err := http.Get(fmt.Sprintf(apiUrl, strings.ToUpper(currency)))
	var cryptoRate CEXResponse

	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusOK {
		body, err := io.ReadAll(response.Body)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &cryptoRate)
	} else {
		return nil, fmt.Errorf("Status code received: %v", response.StatusCode)
	}

	rate := datatypes.Rate{Currency: cryptoRate.Pair, Price: cryptoRate.Ask}

	return &rate, nil
}
