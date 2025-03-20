package utils

import "github.com/google/uuid"

func GenerateID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}