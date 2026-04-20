package main

import (
	"regexp"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	id := GenerateUUID()
	uuidPattern := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	if !uuidPattern.MatchString(id) {
		t.Errorf("GenerateUUID() = %q, does not match UUID format", id)
	}
}

func TestGenerateUUIDUniqueness(t *testing.T) {
	id1 := GenerateUUID()
	id2 := GenerateUUID()
	if id1 == id2 {
		t.Errorf("GenerateUUID() returned duplicate UUIDs: %s", id1)
	}
}
