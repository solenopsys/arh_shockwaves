package wrappers

import (
	"bytes"
	"encoding/json"
	"io"
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
		println("pin exists skip this step", cid)
	} else {
		pin, err := p.SimplePin(cid, labels)
		if err != nil {
			return err
		}
		println("pin", pin)
	}

	//ipnsName:=strings.ReplaceAll(name,"/","_")

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
	return p.Check(name, "name", "name")
}

func (p *Pinning) CheckPin(name string) (bool, error) {
	return p.Check(name, "pin", "cid")
}

func (p *Pinning) Check(cid string, sub string, paramName string) (bool, error) {

	req, err := http.NewRequest("GET", p.Host+"/check/"+sub+"?"+paramName+"="+cid, nil)

	// set header Authorization
	req.Header.Set("Authorization", p.UserKey)
	if err != nil {
		return false, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	body, err := io.ReadAll(resp.Body)

	var boolResponse BoolResponse
	err = json.Unmarshal(body, &boolResponse)
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

func (p *Pinning) SetName(cid string, name string, new bool) (string, error) {
	var method string = "create"
	if !new {
		method = "update"
	}

	url := p.Host + "/name/" + method + "?cid=" + cid + "&name=" + name
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", p.UserKey)
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
	req.Header.Set("Authorization", p.UserKey)
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
