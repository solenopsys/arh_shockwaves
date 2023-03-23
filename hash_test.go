package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"xs/utils"
)

func TestHash(t *testing.T) {

	hash := utils.GenHash("bla2", "bla1")
	assert.Equal(t, "e64938fc6124b4dfa8a2f225cc4998df473cbd6710c364684a1f42f6257d8f8c", hash)
}
