package datastore

import (
	"github.com/google/uuid"

	"github.com/dyweb/sundial/pkg/models"
)

func (ds *datastore) GetProjects() ([]models.Project, error) {
	projects := []models.Project{}
	err := ds.Find(projects).Error
	return projects, err
}

func (ds *datastore) GetProject(UUID uuid.UUID) (*models.Project, error) {
	var project = &models.Project{}
	err := ds.First(project, "id = ?", UUID).Error
	return project, err
}

func (ds *datastore) CreateProject(project *models.Project) error {
	return ds.Create(project).Error
}
