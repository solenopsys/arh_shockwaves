package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"xs/internal/compilers/sorters/fl"
	"xs/internal/configs"
)

func PackCreateHelper(name string, depsNames []string) *configs.NpmLibPackage {
	deps := map[string]string{}
	for _, depName := range depsNames {
		deps[depName] = "file:" + depName
	}

	return &configs.NpmLibPackage{
		Name:                       name,
		AllowedNonPeerDependencies: deps,
	}
}

func TestNextList(t *testing.T) {
	ord := fl.NewNpmLibPackagesOrder(false)
	ord.AddPackage(PackCreateHelper("a", []string{"b", "c"}))
	ord.AddPackage(PackCreateHelper("b", []string{"c"}))
	ord.AddPackage(PackCreateHelper("c", []string{}))

	list := ord.NextList()

	assert.Equal(t, 1, len(list))
	assert.Equal(t, "c", list[0].Name)

	ord.CompileList(list)

	list2 := ord.NextList()
	assert.Equal(t, 1, len(list2))
	assert.Equal(t, "b", list2[0].Name)

	ord.CompileList(list2)

	list3 := ord.NextList()
	assert.Equal(t, 1, len(list3))
	assert.Equal(t, "a", list3[0].Name)

	ord.CompileList(list3)

	list4 := ord.NextList()
	assert.Equal(t, 0, len(list4))
}
