package services

import "xs/utils"

type CompileCache struct {
	cacheIndexDir string
}

func NewCompileCache(cacheIndexDir string) *CompileCache {
	if !utils.DirExists(cacheIndexDir) {
		err := utils.CreateDirs(cacheIndexDir)
		if err != nil {
			panic(err)
		}
	}
	return &CompileCache{cacheIndexDir}
}

func (c *CompileCache) clear() error {
	return utils.ClearDirectory(c.cacheIndexDir)
}

func (c *CompileCache) saveHash(srcHash string, dstHash string) error {
	return utils.WriteFile(c.cacheIndexDir+"/"+srcHash, []byte(dstHash))
}

func (c *CompileCache) checkHash(srcHash string, dstHash string) bool {
	dstHashFromFile, err := utils.ReadFile(c.cacheIndexDir + "/" + srcHash)
	if err != nil {
		return false
	} else {
		return string(dstHashFromFile) == dstHash
	}
}
