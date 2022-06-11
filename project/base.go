package project

import "kel15/usecase"

type Project struct {
	Usecase *usecase.Usecase
}

func NewProject(usecase *usecase.Usecase) *Project {
	return &Project{Usecase: usecase}
}
