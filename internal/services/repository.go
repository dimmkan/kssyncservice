package services

import "kssyncservice_go/pkg/db"

type ServicesRepository struct {
	Database *db.Db
}

func NewServicesRepository(database *db.Db) *ServicesRepository {
	return &ServicesRepository{
		Database: database,
	}
}

func (repo *ServicesRepository) GetAllServices() (*[]Service, error) {
	var services []Service
	result := repo.Database.DB.Limit(100).Find(&services)

	if result.Error != nil {
		return nil, result.Error
	}

	return &services, nil
}