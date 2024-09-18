package client

import (
	"encoding/json"
	"fmt"
	"github.com/fichca/junior-astrologer-service/internal/config"
	"github.com/fichca/junior-astrologer-service/internal/model"
	"io"
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

func (c *client) GetAPOD() (*model.APODClientResponse, error) {
	url := fmt.Sprintf("%s?api_key=%s", c.cfg.APODBaseURL, c.cfg.ApiKey)
	resp, err := c.c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apod model.APODClientResponse

	err = json.NewDecoder(resp.Body).Decode(&apod)
	if err != nil {
		return nil, err
	}
	return &apod, nil
}

func (c *client) DownloadImage(imageURL string) (io.Reader, error) {
	resp, err := c.c.Get(imageURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download image: %s", resp.Status)
	}

	return resp.Body, nil
}
