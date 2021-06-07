package repositories

import (
	pg "github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	models "example.com/grpc_with_go/models"
)

type Manager interface {
	GetByID(user *models.User) (*models.User, error)
	Create(user *models.User) error
	CreateSchema() error
}

type manager struct {
	db *pg.DB
}

var Mgr Manager

func init() {

	Mgr = &manager{db: pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "fred",
		Database: "postgres",
	})}
	Mgr.CreateSchema()
}

func (mgr *manager) CreateSchema() error {
	models := []interface{}{
		(*models.User)(nil),
	}

	for _, model := range models {
		err := mgr.db.Model(model).CreateTable(&orm.CreateTableOptions{
			// Temp:        true,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (mgr *manager) GetByID(user *models.User) (*models.User, error) {
	err := mgr.db.Model(user).WherePK().Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (mgr *manager) Create(user *models.User) error {
	_, err := mgr.db.Model(user).Insert()
	if err != nil {
		return err
	}
	return nil
}
