package repository

import "github.com/Amirhos-esm/portfolio/models"

type Repository interface {
	GetPersonalInformation() (*models.PersonalInformation, error)
	UpdatePersonalInformation(info *models.PersonalInformation) error

	GetAllExperiences() ([]*models.Experience, error)
	GetExperience(id uint) (*models.Experience, error)
	AddExperience(exp *models.Experience) error
	UpdateExperience(id uint, exp *models.Experience) error
	DeleteExperience(id uint) error

	GetAllEducation() ([]*models.Education, error)
	GetEducation(id uint) (*models.Education, error)
	AddEducation(edu *models.Education) error
	UpdateEducation(id uint, edu *models.Education) error
	DeleteEducation(id uint) error

	GetAllProjects() ([]*models.Project, error)
	GetProject(id uint) (*models.Project, error)
	AddProject(proj *models.Project) error
	UpdateProject(id uint, proj *models.Project) error
	DeleteProject(id uint) error

	GetSkills() (*models.Skills, error)
	UpdateSkills(*models.Skills) error

	GetAllData() (*models.DataStore, error)
}

type MessageRepository interface {
	Create(*models.Message) error
	GetbyId(id uint) (*models.Message, error)
	Get(offset uint,limit uint)([]*models.Message, error)
	Delete(id uint) error
}
