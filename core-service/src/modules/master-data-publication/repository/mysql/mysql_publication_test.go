package mysql_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mdpRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/master-data-publication/repository/mysql"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var mockStruct domain.MasterDataPublication

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	err = faker.FakeData(&mockStruct)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when generate fake data", err)
	}

	query := "DELETE FROM masterdata_publications WHERE id = ?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(mockStruct.ID).WillReturnResult(sqlmock.NewResult(mockStruct.ID, 1))

	a := mdpRepo.NewMysqlMasterDataPublicationRepository(db)

	num := int64(mockStruct.ID)
	err = a.Delete(context.TODO(), num)
	assert.NoError(t, err)
}
