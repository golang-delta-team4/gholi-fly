package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/http/port"
	"github.com/golang-delta-team4/gholi-fly/transportCompany/pkg/adapters/clients/http/presenter"
	"github.com/google/uuid"
)

type httpPathClient struct {
	port int
	// host string
}

func NewHttpPathClient(port int) port.HttpPathClient {
	return &httpPathClient{ port: port}
}

var ErrPathNotFound = errors.New("path not found")

func (h *httpPathClient) GetPathDetail(pathID uuid.UUID) (*presenter.GetPathByIDResponse, error) {
	url := fmt.Sprintf("http://:%d/api/v1/paths/filter?id=%s", h.port, pathID)
	req, err := http.NewRequest("GET", url, nil)
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
		return nil, fmt.Errorf("response is not ok: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}
	var pathDetail []presenter.GetPathByIDResponse
	if err := json.Unmarshal(body, &pathDetail); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}
	if len(pathDetail) == 0 {
		return nil, ErrPathNotFound
	}
	return &pathDetail[0], nil

}