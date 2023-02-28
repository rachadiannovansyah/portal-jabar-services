package mysql_test

import (
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	ucRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/uptd-cabdin/repository/mysql"
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
	mockStruct := []domain.UptdCabdin{}
	err = faker.FakeData(&mockStruct)
	if err != nil {
		t.Fatalf(`an error ‘%s’ was not expected when faking detail struct`, err)
	}

	// create rows to be created from a sql driver.Value slice
	rows := sqlmock.NewRows([]string{
		"id",
		"prk_name",
		"cbg_name",
		"cbg_kotakab",
		"cbg_alamat",
		"cbg_notlp",
		"cbg_jenis",
	}).
		AddRow(
			mockStruct[0].ID,
			mockStruct[0].PrkName,
			mockStruct[0].CbgName,
			mockStruct[0].CbgKotaKab,
			mockStruct[0].CbgAlamat,
			mockStruct[0].CbgNoTlp,
			mockStruct[0].CbgJenis,
		)

	query := `SELECT id, prk_name, cbg_name, cbg_kotakab, cbg_alamat, cbg_notlp, cbg_jenis FROM uptd_cabdins WHERE 1=1`

	mock.ExpectQuery(query).WillReturnRows(rows)
	pb := ucRepo.NewMysqlUptdCabdinRepository(db)

	list, err := pb.Fetch(context.TODO())

	// make an assertions using testify
	assert.NotNil(t, list)
	assert.NoError(t, err)
}
