package git

type GitInterface interface {
	IsRepoExists() bool
	GitClone(bare bool) error
	RemoveRemote(name string) error
	SetRemote(name string, url string) error
	GitAddSubmodule() error
	SetHead(branch string) error
	GitUpdateSubmodules() error
	GitUpdate() error
	CloneFromIpfs() error
	//UpdateServerInfo() error
}
