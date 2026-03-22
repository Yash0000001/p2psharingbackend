package utils

import "github.com/google/uuid"

func UUID() string {
	newUUID := uuid.New()
	return newUUID.String()
}
