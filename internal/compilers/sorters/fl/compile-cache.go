package fl

import (
	"xs/pkg/io"
	"xs/pkg/tools"
)

type CompileCache struct {
	cacheIndexDir string
}

func NewCompileCache(cacheIndexDir string) *CompileCache {
	if !tools.Exists(cacheIndexDir) {
		err := tools.CreateDirs(cacheIndexDir)
		if err != nil {
			io.Panic(err)
		}
	}
	return &CompileCache{cacheIndexDir}
}

func (c *CompileCache) clear() error {
	return tools.ClearDirectory(c.cacheIndexDir)
}

func (c *CompileCache) SaveHash(srcHash string, dstHash string) error {
	return tools.WriteFile(c.cacheIndexDir+"/"+srcHash, []byte(dstHash))
}

func (c *CompileCache) CheckHash(srcHash string, dstHash string) bool {
	dstHashFromFile, err := tools.ReadFile(c.cacheIndexDir + "/" + srcHash)
	if err != nil {
		return false
	} else {
		return string(dstHashFromFile) == dstHash
	}
}
