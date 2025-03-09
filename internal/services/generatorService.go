package services

type generatorService struct {
}

func NewGeneatorService() GeneratorService {
	return &generatorService{}
}

func (s *generatorService) CreateNewPassword() string {
	return "password"
}
