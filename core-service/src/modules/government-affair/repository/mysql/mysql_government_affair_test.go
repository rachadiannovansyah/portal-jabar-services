package mysql_test

import (
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	gaRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/government-affair/repository/mysql"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFetch(t *testing.T) {
	// create sql mock db
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf(`an error ‘%s’ was not expected when opening a stub  
        database connection`, err)
	}
	defer db.Close()

	// create mock data from struct
	mockStruct := []domain.GovernmentAffair{}
	err = faker.FakeData(&mockStruct)
	if err != nil {
		t.Fatalf(`an error ‘%s’ was not expected when faking detail service public struct`, err)
	}

	// create rows to be created from a sql driver.Value slice
	rows := sqlmock.NewRows([]string{
		"id",
		"main_affair",
		"sub_main_affair",
	}).
		AddRow(
			mockStruct[0].ID,
			mockStruct[0].MainAffair,
			mockStruct[0].SubMainAffair,
		)

	query := `SELECT id, main_affair, sub_main_affair FROM government_affairs WHERE 1=1`

	mock.ExpectQuery(query).WillReturnRows(rows)
	pb := gaRepo.NewMysqlGovernmentAffairRepository(db)

	list, err := pb.Fetch(context.TODO())

	// make an assertions using testify
	assert.NotNil(t, list)
	assert.NoError(t, err)
}
