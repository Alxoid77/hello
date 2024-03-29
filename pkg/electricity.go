package pkg

//
import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	ELECTRICITY_ON_MESSAGE  = "світло є :)"
	ELECTRICITY_OFF_MESSAGE = "ми во тьмє :("
	SERVER_ERROR_MESSAGE    = "бабка надвоє гадала, або є або не є..."
)

type electricity struct {
	TimeDiff          int    `json:"TimeDiff"`
	PowerStatus       string `json:"PowerStatus"`
	LastObserverdTime string `json:"LastObserverdTime"`
	CurrentTime       string `json:"CurrentTime"`
}

func GetElectricityStatus() string {
	var url = os.Getenv("ELECTRICITY_URL")
	var status = ELECTRICITY_OFF_MESSAGE

	client := http.Client{
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Println(err)

		return SERVER_ERROR_MESSAGE
	}

	res, getErr := client.Do(req)

	if getErr != nil {
		log.Println(getErr)

		return SERVER_ERROR_MESSAGE
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)

	if readErr != nil {
		log.Println(readErr)

		return SERVER_ERROR_MESSAGE
	}

	electricity := electricity{}
	jsonErr := json.Unmarshal(body, &electricity)

	if jsonErr != nil {
		log.Println(jsonErr)

		return SERVER_ERROR_MESSAGE
	}

	if electricity.PowerStatus == "PowerON" {
		status = ELECTRICITY_ON_MESSAGE
	}

	return status
}
