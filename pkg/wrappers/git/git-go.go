package git

import (
	"archive/tar"
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"io"
	"net/http"
	"os"
	"strings"
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
		Hash:  ref.Hash(),
		Force: true,
		//	Branch: plumbing.ReferenceName(targetRefName),

		Create: false, // Set to true if you want to create a new branch if it doesn't exist.
	})
	if err != nil {
		return err
	}

	return nil
}

func replaceRootDir(originalStr string) string {
	parts := strings.Split(originalStr, "/")

	for i, _ := range parts {
		if i == 0 {
			parts[i] = ".git"
			break // Replace only the first occurrence
		}
	}

	// Join the parts back into a string
	return strings.Join(parts, "/")
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
		targetPath := toDir + "/" + replaceRootDir(header.Name)

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

func (g *GoGit) SetHead(branch string) error {
	repo, err := g.OpenRepo()
	if err != nil {
		return err
	}

	// Get the hash of the latest commit on the specified branch
	branchRefName := plumbing.ReferenceName("refs/heads/" + branch)
	branchRef, err := repo.Reference(branchRefName, true)
	if err != nil {
		return err
	}
	if branchRef == nil {
		return errors.New("branch does not exist")
	}
	//	latestCommitHash := branchRef.Hash()

	//	io.Debug("latestCommitHash: " + latestCommitHash.String())

	// Set HEAD to point to the latest commit
	headRefName := plumbing.HEAD
	headRef := plumbing.NewSymbolicReference(headRefName, branchRefName)
	if err := repo.Storer.SetReference(headRef); err != nil {
		return err
	}
	return nil
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
//	// Replace with the path to your Git repository
//	repoPath := "/path/to/your/git/repo"
//
//	// Open the Git repository
//	repo, err := git.PlainOpen(repoPath)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get the object storage to unpack objects
//	objStorage, err := repo.Storer.ObjectStorage()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get a list of packfiles in the .git/objects/pack directory
//	packfileDir := filepath.Join(repoPath, ".git", "objects", "pack")
//	packfiles, err := filepath.Glob(filepath.Join(packfileDir, "*.pack"))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, packfilePath := range packfiles {
//		fmt.Printf("Unpacking objects from: %s\n", packfilePath)
//
//		// Open the packfile
//		packfile, err := os.Open(packfilePath)
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer packfile.Close()
//
//		// Create a packfile reader
//		reader, err := io.Reader()
//
//		// Iterate through the packfile and write objects to the object storage
//		err = reader.ForEach(func(obj plumbing.EncodedObject) error {
//			// Calculate the object hash
//			objHash := obj.Hash()
//
//			// Create the object directory structure based on the hash
//			objDir := filepath.Join(repoPath, ".git", "objects", objHash.String()[:2])
//			err := os.MkdirAll(objDir, os.ModePerm)
//			if err != nil {
//				return err
//			}
//
//			// Create the object file
//			objPath := filepath.Join(objDir, objHash.String()[2:])
//			objFile, err := os.Create(objPath)
//			if err != nil {
//				return err
//			}
//			defer objFile.Close()
//
//			// Encode and write the object to the object file
//			_, err = obj.Encode(objFile)
//			if err != nil {
//				return err
//			}
//
//			fmt.Printf("Unpacked object: %s\n", objHash)
//			return nil
//		})
//
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//	return nil
//}
