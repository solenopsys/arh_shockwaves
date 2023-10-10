package wrappers

import (
	"bytes"
	"encoding/json"
	"github.com/spf13/viper"
	iio "io"
	"xs/internal/configs"
	"xs/pkg/io"

	"net/http"
)

type BoolResponse struct {
	Result bool `json:"result"`
}

type IdResponse struct {
	Id string `json:"id"`
}

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
	Host    string
	UserKey string
}

func (p *Pinning) SmartPin(cid string, labels map[string]string) (string, error) {
	hasPin, err := p.CheckPin(cid)
	if err != nil {
		return "", err
	}

	if hasPin {
		_, err := p.SimpleUpdateLabels(cid, labels)
		io.Debug("pin exists label updated", cid)
		if err != nil {
			return "", err
		}

	}

	return p.SimplePin(cid, labels)
}

func (p *Pinning) SmartPinAndName(cid string, labels map[string]string, ipnsName string) error {
	pin, err := p.SmartPin(cid, labels)

	if err == nil {
		io.Debug("pin", pin)
	}

	ipnsCid, pinErr := p.SmartName(ipnsName, cid)
	io.Debug("ipnsCID", ipnsCid)
	return pinErr
}

func (p *Pinning) SmartName(ipnsName string, cid string) (string, error) {
	hasName, err := p.CheckName(ipnsName)
	if err != nil {
		return "", err
	}

	ipnsId, err := p.SetName(cid, ipnsName, !hasName)
	if err != nil {
		return "", err
	}

	return ipnsId, nil
}

func (p *Pinning) CheckName(name string) (bool, error) {
	return p.Check(name, "name", "value")
}

func (p *Pinning) CheckPin(name string) (bool, error) {
	return p.Check(name, "pin", "cid")
}

func (p *Pinning) execRequestBytes(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", p.UserKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return iio.ReadAll(resp.Body)
}

func (p *Pinning) execRequest(req *http.Request, resp interface{}) error {
	req.Header.Set("Content-Type", "application/json")
	body, err := p.execRequestBytes(req)
	if err != nil {
		return err
	}
	io.Debug(string(body))
	return json.Unmarshal(body, &resp)
}

func (p *Pinning) Check(cid string, sub string, paramName string) (bool, error) {
	url := p.Host + "/check/" + sub + "?" + paramName + "=" + cid
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}
	var boolResponse BoolResponse
	err = p.execRequest(req, &boolResponse)

	if err != nil {
		return false, err
	} else {
		return boolResponse.Result, nil
	}
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

func (p *Pinning) SimpleUpdateLabels(cid string, labels map[string]string) (string, error) {

	pin := Pin{
		CID:    cid,
		Labels: labels,
	}

	conf := Configuration{
		Pins: []Pin{pin},
	}

	return p.UpdateLabels(&conf)

}

func (p *Pinning) SetName(cid string, name string, new bool) (string, error) {
	var method string = "create"
	if !new {
		method = "update"
	}

	url := p.Host + "/name/" + method + "?cid=" + cid + "&name=" + name
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	body, err := p.execRequestBytes(req)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (p *Pinning) Pin(conf *Configuration) (string, error) {
	jsonData, err := json.Marshal(conf)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", p.Host+"/pin", bytes.NewBuffer(jsonData))

	body, err := p.execRequestBytes(req)

	return string(body), nil
}

func (p *Pinning) UpdateLabels(conf *Configuration) (string, error) {
	jsonData, err := json.Marshal(conf)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("PUT", p.Host+"/labels", bytes.NewBuffer(jsonData))

	body, err := p.execRequestBytes(req)

	return string(body), nil
}

func (p *Pinning) FindName(namePattern string, valuePattern string) ([]byte, error) {
	url := p.Host + "/select/names?name=" + namePattern + "&value=" + valuePattern
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return p.execRequestBytes(req)
}

func (p *Pinning) FindResource(namePattern string, valuePattern string) ([]byte, error) {
	url := p.Host + "/select/pins?name=" + namePattern + "&value=" + valuePattern
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return p.execRequestBytes(req)
}

func NewPinning() *Pinning {
	return &Pinning{
		Host:    viper.GetString("hosts.pinningService"),
		UserKey: configs.GetAuthManager().PublicKey,
	}
}
