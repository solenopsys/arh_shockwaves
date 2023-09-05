package wrappers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Pin struct {
	CID    string            `json:"cid"`
	Labels map[string]string `json:"labels"`
}

type Configuration struct {
	Pins   []Pin `json:"pins"`
	RepMin int   `json:"rep_min"`
	RepMax int   `json:"rep_max"`
}

type Pinning struct {
	Host string `json:"pinning"`
}

func (p *Pinning) SimplePin(cid string, labels map[string]string) (string, error) {

	pin := Pin{
		CID:    cid,
		Labels: labels,
	}

	conf := Configuration{
		Pins:   []Pin{pin},
		RepMin: 2,
		RepMax: 3,
	}

	return p.Pin(&conf)

}

func (p *Pinning) SetName(cid string, name string) (string, error) {
	req, err := http.NewRequest("GET", p.Host+"/ipns/create?cid="+cid+"&name="+name, nil)

	if err != nil {
		return "", err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)

	return string(body), nil
}

func (p *Pinning) Pin(conf *Configuration) (string, error) {

	jsonData, err := json.Marshal(conf)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", p.Host+"/pin", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	return string(body), nil

}
