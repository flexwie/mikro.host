package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDb(t *testing.T) {
	db := GetDb("file::memory:?cache=shared")
	assert.NotNil(t, db)
}
