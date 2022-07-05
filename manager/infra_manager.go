package manager

import (
	"api_gin_ref/config"
	"api_gin_ref/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Infra di sini bertugas sebagai databse penyimpanan pengganti slice
type Infra interface {
	SqlDb() *gorm.DB
}

type infra struct {
	db *gorm.DB
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func NewInfra(config *config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &infra{db: resource}
}

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	db.AutoMigrate(&model.Product{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
