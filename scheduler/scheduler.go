// scheduler.go
package scheduler

import (
	"bytes"
	"encoding/json"
	"integrasi-gbk-online/controller"
	"integrasi-gbk-online/models"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/robfig/cron/v3"
)

func StartScheduler() {
	c := cron.New()
	c.AddFunc("@every 1m", func() {
		log.Println("Menarik data dari API...")

		// Get token from API
		token, err := GetAPIToken("https://reservation.gbk.id/api/info/token")
		if err != nil {
			log.Printf("Error saat mendapatkan token dari API: %v", err)
			return
		}

		// Fetch data from API with authorization header
		venues, err := FetchDataFromVenue("https://reservation.gbk.id/api/info/venue", token)
		if err != nil {
			log.Printf("Error saat menarik data dari API: %v", err)
			return
		}

		log.Println("Menyimpan data ke database...")
		err = controller.SaveVenuesToDB(venues)
		if err != nil {
			log.Printf("Error saat menyimpan data ke database: %v", err)
		}
	})

	c.Start()
	log.Println("Scheduler berjalan...")
}

func GetAPIToken(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(`{
		"PartnerID": 11983,
		"PartnerKey": "53fdd7781054c10b94c1cbf63baaa74f"
	}`)))
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var apiResponse models.TokenResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return "", err
	}
	return apiResponse.Token, nil
}

func FetchDataFromVenue(url, token string) ([]models.Venue, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// Set request body
	reqBody := map[string]interface{}{
		"PartnerID": 11983,
		"Name":      "",
		"IsActive":  1,
		"Status":    0,
	}
	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(reqJSON))

	resp, err := client.Do(req)
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
