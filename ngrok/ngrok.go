package ngrok

import (
	"fmt"
	"github.com/ViRb3/sling/v2"
	"github.com/pkg/errors"
	"ios-signer-service/util"
	"strings"
)

func GetPublicUrl(ngrokPort uint64, proto string) (string, error) {
	var tunnels Tunnels
	response, err := sling.New().Get(fmt.Sprintf("http://localhost:%d/api/tunnels", ngrokPort)).ReceiveSuccess(&tunnels)
	if err != nil {
		return "", err
	}
	if err := util.Check2xxCode(response.StatusCode); err != nil {
		return "", err
	}
	for _, tunnel := range tunnels.Tunnels {
		if strings.EqualFold(tunnel.Proto, proto) {
			return tunnel.PublicURL, nil
		}
	}
	return "", errors.New("tunnel not found")
}

type Tunnels struct {
	Tunnels []Tunnel `json:"tunnels"`
	URI     string   `json:"uri"`
}
type Config struct {
	Addr    string `json:"addr"`
	Inspect bool   `json:"inspect"`
}
type Conns struct {
	Count  int     `json:"count"`
	Gauge  float64 `json:"gauge"`
	Rate1  float64 `json:"rate1"`
	Rate5  float64 `json:"rate5"`
	Rate15 float64 `json:"rate15"`
	P50    float64 `json:"p50"`
	P90    float64 `json:"p90"`
	P95    float64 `json:"p95"`
	P99    float64 `json:"p99"`
}
type HTTP struct {
	Count  int     `json:"count"`
	Rate1  float64 `json:"rate1"`
	Rate5  float64 `json:"rate5"`
	Rate15 float64 `json:"rate15"`
	P50    float64 `json:"p50"`
	P90    float64 `json:"p90"`
	P95    float64 `json:"p95"`
	P99    float64 `json:"p99"`
}
type Metrics struct {
	Conns Conns `json:"conns"`
	HTTP  HTTP  `json:"http"`
}
type Tunnel struct {
	Name      string  `json:"name"`
	URI       string  `json:"uri"`
	PublicURL string  `json:"public_url"`
	Proto     string  `json:"proto"`
	Config    Config  `json:"config"`
	Metrics   Metrics `json:"metrics"`
}
