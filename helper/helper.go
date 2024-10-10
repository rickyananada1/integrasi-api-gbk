// api_client.go
package helper

import (
	"encoding/json"
	"integrasi-gbk-online/models"
	"io/ioutil"
	"net/http"
)

func FetchDataFromAPI(url string) ([]models.Venue, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse models.ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}

	return apiResponse.Data, nil
}
