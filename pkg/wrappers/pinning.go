package wrappers

import (
	"bytes"
	"encoding/json"
	"errors"
	iio "io"
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

func (p *Pinning) SmartPin(cid string, labels map[string]string, ipnsName string) error {
	hasPin, err := p.CheckPin(cid)
	if err != nil {
		return err
	}

	if hasPin {
		_, err := p.SimpleUpdateLabels(cid, labels)
		if err != nil {
			return err
		}
		println("pin exists label updated", cid)
	} else {
		pin, err := p.SimplePin(cid, labels)
		if err != nil {
			return err
		}
		println("pin", pin)
	}

	hasName, err := p.CheckName(ipnsName)
	if err != nil {
		return err
	}

	ipnsId, err := p.SetName(cid, ipnsName, !hasName)
	if err != nil {
		return err
	}
	println("ipnsId", ipnsId)

	return nil
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
	println(string(body))
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

type PackInfo struct {
	Cid string
	To  string
}

func (p *Pinning) FindOne(packageName string) (*PackInfo, error) {
	repo, err := p.FindRepo(packageName)
	if err != nil {
		return nil, err
	}
	for _, v := range *repo {
		return &v, nil
	}
	return nil, errors.New("not found")
}

func (p *Pinning) FindRepo(repoName string) (*map[string]PackInfo, error) {
	namePattern := "code*"
	url := p.Host + "/select/names?name=" + namePattern + "&value=" + repoName
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	body, err := p.execRequestBytes(req)

	var resp map[string]map[string]string

	err = json.Unmarshal(body, &resp)

	mapping := make(map[string]PackInfo)

	if err != nil {
		io.Fatal(err)
	}
	for ipnsCid, mp := range resp {
		info := PackInfo{}
		info.Cid = ipnsCid
		info.To = mp["clone.to"]

		mapping[mp["code.source"]] = info

	}
	return &mapping, nil
}

func NewPinning() *Pinning {
	return &Pinning{
		Host:    "http://pinning.solenopsys.org", // todo remove it
		UserKey: "alexstorm",                     // todo remove it
	}
}
