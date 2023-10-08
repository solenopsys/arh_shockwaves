package services

import (
	"encoding/json"
	"errors"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

type PinningRequests struct {
	pinning *wrappers.Pinning
}

type PackInfo struct {
	Cid string
	To  string
	Src string
}

func (p *PinningRequests) FindFontLib(fileName string) (string, error) {
	namePattern := "front.static.library" // todo move to const
	body, err := p.pinning.FindResource(namePattern, fileName)

	var resp map[string]map[string]string

	err = json.Unmarshal(body, &resp)

	if err != nil {
		io.Fatal(err)
	}

	for ipfsCid, _ := range resp {
		return ipfsCid, nil
	}

	return "", errors.New("front static library not found")
}

func (p *PinningRequests) FindRepo(repoName string) (*map[string]PackInfo, error) {
	namePattern := "code*"

	body, err := p.pinning.FindName(namePattern, repoName)

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
		info.Src = mp["source.url"]
		mapping[mp["code.source"]] = info
	}
	return &mapping, nil
}

func (p *PinningRequests) FindOne(packageName string) (*PackInfo, error) {
	repo, err := p.FindRepo(packageName)
	if err != nil {
		return nil, err
	}
	for _, v := range *repo {
		return &v, nil
	}
	return nil, errors.New("not found")
}

func NewPinningRequests() *PinningRequests {

	return &PinningRequests{wrappers.NewPinning()}

}
