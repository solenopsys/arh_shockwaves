package git

type GitInterface interface {
	IsRepoExists() bool
	GitClone() error
	RemoveRemote(name string) error
	SetRemote(name string, url string) error
	GitAddSubmodule() error
	GitUpdateSubmodules() error
	GitUpdate() error
	//UpdateServerInfo() error
}
