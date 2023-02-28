package mysql_test

import (
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	srRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/spbe-rals/repository/mysql"
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
	mockStruct := []domain.SpbeRals{}
	err = faker.FakeData(&mockStruct)
	if err != nil {
		t.Fatalf(`an error ‘%s’ was not expected when faking detail struct`, err)
	}

	// create rows to be created from a sql driver.Value slice
	rows := sqlmock.NewRows([]string{
		"id",
		"kode_ral_2",
		"kode",
		"item",
	}).
		AddRow(
			mockStruct[0].ID,
			mockStruct[0].RalCode2,
			mockStruct[0].Code,
			mockStruct[0].Item,
		)

	query := `SELECT id, kode_ral_2, kode, item FROM spbe_rals WHERE 1=1`

	mock.ExpectQuery(query).WillReturnRows(rows)
	pb := srRepo.NewMysqlSpbeRalsRepository(db)

	params := &domain.Request{
		Keyword:   "",
		PerPage:   10,
		Offset:    0,
		SortBy:    "",
		SortOrder: "",
	}

	list, _, err := pb.Fetch(context.TODO(), params)

	// make an assertions using testify
	assert.NotNil(t, list)
	assert.NoError(t, err)
}
