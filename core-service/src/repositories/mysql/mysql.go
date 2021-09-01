package mysql

import (
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

// Repositories ...
type Repositories struct {
	CategoryRepo    domain.CategoryRepository
	NewsRepo        domain.NewsRepository
	InformationRepo domain.InformationRepository
	UnitRepo        domain.UnitRepository
	AreaRepo        domain.AreaRepository
	EventRepo       domain.EventRepository
}

// NewMysqlRepositories will create an object that represent all repos interface
func NewMysqlRepositories(Conn *sql.DB) *Repositories {
	return &Repositories{
		CategoryRepo:    NewMysqlCategoryRepository(Conn),
		NewsRepo:        NewMysqlNewsRepository(Conn),
		InformationRepo: NewMysqlInformationRepository(Conn),
		UnitRepo:        NewMysqlUnitRepository(Conn),
		AreaRepo:        NewMysqlAreaRepository(Conn),
		EventRepo:       NewMysqlEventRepository(Conn),
	}
}
