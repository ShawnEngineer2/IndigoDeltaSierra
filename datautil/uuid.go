package datautil

import (
	"github.com/google/uuid"
)

func GetUUID() string {
	eventUUID := uuid.New()
	return eventUUID.String()
}
