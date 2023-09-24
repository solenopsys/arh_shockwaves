package git

import (
	"archive/tar"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

func checkoutFromBareRepository(repoPath, targetRefName string) error {
	// Open the bare repository.
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}

	// Get the reference (branch or commit) you want to checkout.
	ref, err := repo.Reference(plumbing.ReferenceName(targetRefName), true)
	if err != nil {
		return err
	}

	// Create a worktree for the repository.
	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	// Checkout the desired branch or commit.
	err = worktree.Checkout(&git.CheckoutOptions{
		Hash: ref.Hash(),
		//	Branch: plumbing.ReferenceName(targetRefName),

		Create: false, // Set to true if you want to create a new branch if it doesn't exist.
	})
	if err != nil {
		return err
	}

	return nil
}

func downloadAndExtractTar(url, toDir string) error {
	// Create the destination directory if it doesn't exist.
	if err := os.MkdirAll(toDir, os.ModePerm); err != nil {
		return err
	}

	// Send an HTTP GET request to the URL.
	response, err := http.Get(url + "?download=true&format=tar")
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Check if the response status code indicates success (200 OK).
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code %d", response.StatusCode)
	}

	body := response.Body

	// Create a tar reader to read the tarball.
	tarReader := tar.NewReader(body)

	// Iterate through the tarball and extract its contents.
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break // Reached the end of the tarball.
		}
		if err != nil {
			return err
		}

		// Ensure the extracted file or directory is within the destination directory.
		targetPath := filepath.Join(toDir, header.Name)

		// Create parent directories as needed.
		if header.FileInfo().IsDir() {
			if err := os.MkdirAll(targetPath, os.ModePerm); err != nil {
				return err
			}
		} else {
			// Create the target file.
			outputFile, err := os.Create(targetPath)
			if err != nil {
				return err
			}
			defer outputFile.Close()

			// Copy the file content from the tarball to the target file.
			if _, err := io.Copy(outputFile, tarReader); err != nil {
				return err
			}
		}
	}

	return nil
}

func CloneFromIpfs(url string) {
	//	fullUrl := url + "?download=true&format=tar"

}

// GitClone clones a Git repository from the remote URL to the specified path.
func (g *GoGit) GitClone(bare bool) error {

	err := os.MkdirAll(g.BasePath, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = git.PlainClone(g.BasePath, bare, &git.CloneOptions{
		URL: g.Remote,
	})

	return err
}

func (g *GoGit) CloneFromIpfs() error {
	err := downloadAndExtractTar(g.Remote, g.BasePath)
	if err != nil {
		return err
	}

	err = checkoutFromBareRepository(g.BasePath, "refs/heads/master")
	if err != nil {
		return err
	}

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
