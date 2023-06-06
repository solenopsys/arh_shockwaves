package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"xs/internal/funcs"
)

func TestHash(t *testing.T) {
	hash := funcs.GenHash("bla2", "bla1")
	assert.Equal(t, "e64938fc6124b4dfa8a2f225cc4998df473cbd6710c364684a1f42f6257d8f8c", hash)
}
