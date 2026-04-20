package main

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUUID(t *testing.T) {
	id := GenerateUUID()
	uuidPattern := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	assert.Regexp(t, uuidPattern, id, "GenerateUUID() should return a valid UUID")
}

func TestGenerateUUIDUniqueness(t *testing.T) {
	id1 := GenerateUUID()
	id2 := GenerateUUID()
	assert.NotEqual(t, id1, id2, "GenerateUUID() should return unique UUIDs")
}
