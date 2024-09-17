package client

import (
	"encoding/json"
	"fmt"
	"github.com/fichca/junior-astrologer-service/internal/config"
	"github.com/fichca/junior-astrologer-service/internal/model"
	"net/http"
	"time"
)

type client struct {
	cfg *config.Client
	c   *http.Client
}

func NewClient(clientCfg *config.Client) *client {
	return &client{
		cfg: clientCfg,
		c: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (as *client) GetAPOD() (*model.APODResponse, error) {
	url := fmt.Sprintf("%s?api_key=%s", as.cfg.APODBaseURL, as.cfg.ApiKey)
	resp, err := as.c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apod model.APODResponse

	err = json.NewDecoder(resp.Body).Decode(&apod)
	if err != nil {
		return nil, err
	}
	return &apod, nil
}
