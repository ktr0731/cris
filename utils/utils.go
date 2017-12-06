package utils

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

func NewUUID() string {
	uuid := uuid.NewV4().String()
	return strings.Replace(uuid, "-", "", 4)
}
