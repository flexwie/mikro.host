package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDb(t *testing.T) {
	path := ":memory:"
	db := GetDb(&path)
	assert.NotNil(t, db)
}
