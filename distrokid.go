package distrogo

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type DistroKid struct {
	AuthToken string // Basically the Bearer token for the API
}

func NewDistroKid(authToken string) *DistroKid {
	return &DistroKid{
		AuthToken: authToken,
	}
}

func (d *DistroKid) DoRequest(method, url string, data io.Reader) ([]byte, error) {
	method = strings.ToUpper(method)
	req, err := http.NewRequest(method, url, data)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", d.AuthToken))
	req.Header.Set("User-Agent", "DistroKid/369 CFNetwork/1492.0.1 Darwin/23.3.0") // Not required, but I set it to mimic the official client
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(resp.Status, "4") || strings.HasPrefix(resp.Status, "5") {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	defer resp.Body.Close()

	var body []byte

	body, err = io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
