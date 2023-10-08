package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/beruang43221/assignment3-015/models"
)

func HitAPI() {
	for {
		water, wind := generateRandomData()

		// Mengirim permintaan PUT ke API
		if err := sendDataToAPI(water, wind); err != nil {
			fmt.Printf("Error: %s\n", err)
		}

		// Menentukan status dan mencetak log
		waterStatus := getStatus(water, 6, 8)
		windStatus := getStatus(wind, 7, 15)
		logData := models.Microservice{Water: water, Wind: wind}
		printLog(logData, waterStatus, windStatus)

		time.Sleep(15 * time.Second)
	}
}

func generateRandomData() (int, int) {
	return rand.Intn(100) + 1, rand.Intn(100) + 1
}

func sendDataToAPI(water, wind int) error {
	dataToUpdate := models.Microservice{Water: water, Wind: wind}
	apiURL := "http://localhost:8083/update/1"
	payload, err := json.Marshal(dataToUpdate)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", apiURL, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned non-OK status: %s", resp.Status)
	}

	return nil
}

func getStatus(value, safeThreshold, dangerThreshold int) string {
	if value < safeThreshold {
		return "aman"
	} else if value >= safeThreshold && value <= dangerThreshold {
		return "siaga"
	}
	return "bahaya"
}

func printLog(data models.Microservice, waterStatus, windStatus string) {
	logJSON, _ := json.MarshalIndent(data, "", "  ")
	fmt.Printf("%s\nstatus water: %s\nstatus wind: %s\n", string(logJSON), waterStatus, windStatus)
}
