package migration

import (
	"github.com/Witor-Silva/primeiro-crud-go/entity"

	"gorm.io/gorm"
)

type DatabaseMakeMigration struct {
	conn *gorm.DB
}

func NewDatabaseMakeMigration(conn *gorm.DB) *DatabaseMakeMigration {
	return &DatabaseMakeMigration{conn: conn}
}

func (m *DatabaseMakeMigration) MakeMigrations() {
	err := m.conn.Migrator().AutoMigrate(&entity.UserEntity{})
	if err != nil {
		panic(err)
	}
}
