package util

import uuid "github.com/satori/go.uuid"

// GenerateUUID will return a new valid UUID
func GenerateUUID() string {

	// Abstract away the actual implementation making it easier to change
	return uuid.NewV4().String()
}
