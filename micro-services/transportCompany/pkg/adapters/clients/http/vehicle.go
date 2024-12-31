package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/http/port"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/http/presenter"
)

type httpVehicleClient struct {
	port int
	host string
}

func NewHttpVehicleClient( port int) port.HttpVehicleClient {
	return &httpVehicleClient{port: port}
}

var (
	ErrVehicleNotFound = fmt.Errorf("vehicle not found")
)

func (h *httpVehicleClient) GetMatchedVehicle(matchMakerRequest *presenter.MatchMakerRequest) (*presenter.MatchMakerResponse, error) {
	url := fmt.Sprintf("http://127.0.0.1:%d/api/v1/vehicles/match", h.port)
	requestBody, err := json.Marshal(matchMakerRequest)
    if err != nil {
        return nil, fmt.Errorf("error marshalling request: %v", err)
    }

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, ErrVehicleNotFound
		}
		return nil, fmt.Errorf("response is not ok: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var pathDetail presenter.MatchMakerResponse
	if err := json.Unmarshal(body, &pathDetail); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return &pathDetail, nil

}