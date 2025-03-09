package services

type GeneratorService interface {
	CreateNewPassword(length int, useNumbers bool, useSpecial bool) (string, error)
}
