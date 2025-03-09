package services

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

const (
	letters      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers      = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
)

type generatorService struct {
}

func NewGeneatorService() GeneratorService {
	return &generatorService{}
}

func (s *generatorService) CreateNewPassword(length int, useNumbers bool, useSpecial bool) (string, error) {
	if length <= 0 {
		return "", errors.New("password length needs to be bigger than 0")
	}

	charPool := letters
	if useNumbers {
		charPool += numbers
	}
	if useSpecial {
		charPool += specialChars
	}

	var password strings.Builder
	for i := 0; i < length; i++ {
		randIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charPool))))
		if err != nil {
			return "", err
		}
		password.WriteByte(charPool[randIndex.Int64()])
	}

	return password.String(), nil
}
