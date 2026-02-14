package json

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/Amirhos-esm/portfolio/models"
)

/* =======================
   Internal Data Structure
======================= */

/* =======================
   JSON Repository
======================= */

type JSONRepository struct {
	path string
	data models.DataStore
	mu   sync.RWMutex
}

func NewJSONRepository(path string, fallbackData *models.DataStore) (*JSONRepository, error) {
	r := &JSONRepository{path: path}

	file, err := os.ReadFile(path)
	if err == nil {
		if err := json.Unmarshal(file, &r.data); err != nil {
			return nil, err
		}
		return r, nil
	}

	r.data = *fallbackData
	if err := r.save(); err != nil {
		return nil, err
	}
	return r, nil
}

/* =======================
   Helpers
======================= */

func (r *JSONRepository) save() error {
	bytes, err := json.MarshalIndent(r.data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.path, bytes, 0644)
}
func (r *JSONRepository) Save() error {
	return r.save()
}

/* =======================
   Personal Information
======================= */

func (r *JSONRepository) GetPersonalInformation() (*models.PersonalInformation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return &r.data.PersonalInformation, nil
}

func (r *JSONRepository) UpdatePersonalInformation(info *models.PersonalInformation) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data.PersonalInformation = *info
	return r.save()
}

/* =======================
   Experience
======================= */

func (r *JSONRepository) GetAllExperiences() ([]*models.Experience, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*models.Experience, len(r.data.Experiences))
	for i := range r.data.Experiences {
		out[i] = &r.data.Experiences[i]
	}
	return out, nil
}

func (r *JSONRepository) AddExperience(exp *models.Experience) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	var maxID uint
	for _, e := range r.data.Experiences {
		if e.ID > maxID {
			maxID = e.ID
		}
	}
	exp.ID = maxID + 1

	r.data.Experiences = append(r.data.Experiences, *exp)
	return r.save()
}

func (r *JSONRepository) UpdateExperience(id uint, exp *models.Experience) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, e := range r.data.Experiences {
		if e.ID == id {
			exp.ID = id
			r.data.Experiences[i] = *exp
			return r.save()
		}
	}
	return errors.New("experience not found")
}

func (r *JSONRepository) DeleteExperience(id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, e := range r.data.Experiences {
		if e.ID == id {
			r.data.Experiences = append(
				r.data.Experiences[:i],
				r.data.Experiences[i+1:]...,
			)
			return r.save()
		}
	}
	return errors.New("experience not found")
}

/* =======================
   Education
======================= */

func (r *JSONRepository) GetAllEducation() ([]*models.Education, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*models.Education, len(r.data.Educations))
	for i := range r.data.Educations {
		out[i] = &r.data.Educations[i]
	}
	return out, nil

}

func (r *JSONRepository) AddEducation(edu *models.Education) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	var maxID uint
	for _, e := range r.data.Educations {
		if e.ID > maxID {
			maxID = e.ID
		}
	}
	edu.ID = maxID + 1

	r.data.Educations = append(r.data.Educations, *edu)
	return r.save()
}

func (r *JSONRepository) UpdateEducation(id uint, edu *models.Education) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, e := range r.data.Educations {
		if e.ID == id {
			edu.ID = id
			r.data.Educations[i] = *edu
			return r.save()
		}
	}
	return errors.New("education not found")
}

func (r *JSONRepository) DeleteEducation(id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, e := range r.data.Educations {
		if e.ID == id {
			r.data.Educations = append(
				r.data.Educations[:i],
				r.data.Educations[i+1:]...,
			)
			return r.save()
		}
	}
	return errors.New("education not found")
}

/* =======================
   Projects
======================= */

func (r *JSONRepository) GetAllProjects() ([]*models.Project, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*models.Project, len(r.data.Projects))
	for i := range r.data.Projects {
		out[i] = &r.data.Projects[i]
	}
	return out, nil
}

func (r *JSONRepository) AddProject(proj *models.Project) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	var maxID uint
	for _, p := range r.data.Projects {
		if p.ID > maxID {
			maxID = p.ID
		}
	}
	proj.ID = maxID + 1

	r.data.Projects = append(r.data.Projects, *proj)
	return r.save()
}

func (r *JSONRepository) UpdateProject(id uint, proj *models.Project) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, p := range r.data.Projects {
		if p.ID == id {
			proj.ID = id
			r.data.Projects[i] = *proj
			return r.save()
		}
	}
	return errors.New("project not found")
}

func (r *JSONRepository) DeleteProject(id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, p := range r.data.Projects {
		if p.ID == id {
			r.data.Projects = append(
				r.data.Projects[:i],
				r.data.Projects[i+1:]...,
			)
			return r.save()
		}
	}
	return errors.New("project not found")
}
func (r *JSONRepository) GetProject(id uint) (*models.Project, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, p := range r.data.Projects {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, nil
}

func (r *JSONRepository) GetEducation(id uint) (*models.Education, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, e := range r.data.Educations {
		if e.ID == id {
			return &e, nil
		}
	}
	return nil, nil
}

func (r *JSONRepository) GetExperience(id uint) (*models.Experience, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, e := range r.data.Experiences {
		if e.ID == id {
			return &e, nil
		}
	}
	return nil, nil
}

func (r *JSONRepository) GetSkills() (*models.Skills, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return &r.data.Skills, nil
}
func (r *JSONRepository) UpdateSkills(skill *models.Skills) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data.Skills = *skill
	return r.save()
}

func (r *JSONRepository) GetAllData() (*models.DataStore, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return &r.data, nil
}
