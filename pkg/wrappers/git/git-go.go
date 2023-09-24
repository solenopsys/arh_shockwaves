package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
) // with go modules enabled (GO111MODULE=on or outside GOPATH)

type GoGit struct {
	BasePath string
	Remote   string
}

// OpenRepo opens an existing Git repository at the specified path.
func (g *GoGit) OpenRepo() (*git.Repository, error) {
	return git.PlainOpen(g.BasePath)
}

// IsRepoExists checks if a Git repository exists at the specified path.
func (g *GoGit) IsRepoExists() bool {
	_, err := g.OpenRepo()
	return err == nil
}

// GitClone clones a Git repository from the remote URL to the specified path.
func (g *GoGit) GitClone() error {
	_, err := git.PlainClone(g.BasePath, false, &git.CloneOptions{
		URL: g.Remote,
	})
	return err
}

// RemoveRemote removes a remote from the Git repository.
func (g *GoGit) RemoveRemote(name string) error {
	repo, err := g.OpenRepo()
	if err != nil {
		return err
	}

	return repo.DeleteRemote(name)
}

// SetRemote adds a new remote to the Git repository.
func (g *GoGit) SetRemote(name string, url string) error {
	repo, err := g.OpenRepo()
	if err != nil {
		return err
	}

	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name: name,
		URLs: []string{url},
	})
	return err
}

func (g *GoGit) GitAddSubmodule() error {
	repo, err := g.OpenRepo()
	if err != nil {
		return err
	}

	w, err := repo.Worktree()

	if err != nil {
		return err
	}
	sub, err := w.Submodule(g.Remote)
	if err != nil {
		return err
	}
	err = sub.Init()

	if err != nil {
		return err
	}

	return err
}

func (g *GoGit) GitUpdateSubmodules() error {
	repo, err := g.OpenRepo()
	if err != nil {
		return err
	}

	w, err := repo.Worktree()

	if err != nil {
		return err
	}

	sub, err := w.Submodule(g.Remote)
	if err != nil {
		return err
	}

	err = sub.Update(&git.SubmoduleUpdateOptions{
		Init:              true,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	return err
}

// GitUpdate performs a Git update operation on the repository.
func (g *GoGit) GitUpdate() error {
	repo, err := g.OpenRepo()
	if err != nil {
		return err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	// Perform the update operation.
	err = worktree.Pull(&git.PullOptions{
		RemoteName: "origin",
	})

	return err
}

// UpdateServerInfo updates the Git server info.
//func (g *GoGit) UpdateServerInfo() error {
//	repo, err := g.OpenRepo()
//	if err != nil {
//		return err
//	}
//
//	err = repo.UpdateServerInfo()
//	return err
//}

// UnpackObjects unpacks Git objects from pack files.
//func (g *GoGit) UnpackObjects() error {
//	// Implement the UnpackObjects method as per your requirements.
//	// You can use the go-git library to manipulate objects if needed.
//	return nil
//}
