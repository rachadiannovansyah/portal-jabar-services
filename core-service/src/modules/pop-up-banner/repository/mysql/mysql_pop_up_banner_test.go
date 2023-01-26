package mysql_test

import (
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	pRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/pop-up-banner/repository/mysql"
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
	mockStruct := []domain.PopUpBanner{}
	err = faker.FakeData(&mockStruct)
	if err != nil {
		t.Fatalf(`an error ‘%s’ was not expected when faking detail service public struct`, err)
	}

	// create rows to be created from a sql driver.Value slice
	rows := sqlmock.NewRows([]string{
		"id",
		"title",
		"button_label",
		"image",
		"link",
		"status",
		"duration",
		"start_date",
		"end_date",
		"created_at",
		"updated_at",
	}).
		AddRow(
			mockStruct[0].ID,
			mockStruct[0].Title,
			mockStruct[0].ButtonLabel,
			helpers.GetStringFromObject(mockStruct[0].Image),
			mockStruct[0].Link,
			mockStruct[0].Status,
			mockStruct[0].Duration,
			mockStruct[0].StartDate,
			mockStruct[0].EndDate,
			mockStruct[0].CreatedAt,
			mockStruct[0].UpdatedAt,
		)

	query := `SELECT id, title, button_label, image, link, status, duration, start_date, end_date, created_at, updated_at FROM pop_up_banners WHERE 1=1`

	mock.ExpectQuery(query).WillReturnRows(rows)
	pb := pRepo.NewMysqlPopUpBannerRepository(db)

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

func TestGetByID(t *testing.T) {
	// create sql mock db
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf(`an error ‘%s’ was not expected when opening a stub  
        database connection`, err)
	}
	defer db.Close()

	// create mock data from struct
	mockStruct := domain.PopUpBanner{}
	err = faker.FakeData(&mockStruct)
	if err != nil {
		t.Fatalf(`an error ‘%s’ was not expected when faking detail service public struct`, err)
	}

	// create rows to be created from a sql driver.Value slice
	rows := sqlmock.NewRows([]string{
		"id",
		"title",
		"button_label",
		"image",
		"link",
		"status",
		"duration",
		"start_date",
		"end_date",
		"created_at",
		"updated_at",
	}).
		AddRow(
			mockStruct.ID,
			mockStruct.Title,
			mockStruct.ButtonLabel,
			helpers.GetStringFromObject(mockStruct.Image),
			mockStruct.Link,
			mockStruct.Status,
			mockStruct.Duration,
			mockStruct.StartDate,
			mockStruct.EndDate,
			mockStruct.CreatedAt,
			mockStruct.UpdatedAt,
		)

	query := `SELECT id, title, button_label, image, link, status, duration, start_date, end_date, created_at, updated_at FROM pop_up_banners WHERE 1=1`

	mock.ExpectQuery(query).WillReturnRows(rows)
	pb := pRepo.NewMysqlPopUpBannerRepository(db)

	obj, err := pb.GetByID(context.TODO(), int64(1))

	// make an assertions using testify
	assert.NotNil(t, obj)
	assert.NoError(t, err)
}
