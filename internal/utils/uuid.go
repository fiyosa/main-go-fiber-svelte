package utils

import (
	"github.com/google/uuid"
)

func CreateUUID() string {
	return uuid.New().String()
}

func VerifyUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
