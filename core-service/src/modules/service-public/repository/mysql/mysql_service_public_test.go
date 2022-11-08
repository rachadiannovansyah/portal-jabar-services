package mysql_test

import (
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	spRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/service-public/repository/mysql"
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
	mockGenInfoStruct := domain.DetailGeneralInformation{}
	err = faker.FakeData(&mockGenInfoStruct)
	if err != nil {
		t.Fatalf(`an error ‘%s’ was not expected when faking general informations struct`, err)
	}

	mockDetailSpublicStruct := []domain.DetailServicePublicResponse{}
	err = faker.FakeData(&mockDetailSpublicStruct)
	if err != nil {
		t.Fatalf(`an error ‘%s’ was not expected when faking detail service public struct`, err)
	}

	mockSpublicStruct := []domain.ListServicePublicResponse{}
	err = faker.FakeData(&mockSpublicStruct)

	if err != nil {
		t.Fatalf(`an error ‘%s’ was not expected when list service publiv struct`, err)
	}

	// create rows to be created from a sql driver.Value slice
	rows := sqlmock.NewRows([]string{
		"id",
		"purpose",
		"facility",
		"requirement",
		"tos",
		"info_graphic",
		"faq",
		"created_at",
		"updated_at",
		"general_information_id",
		"general_information_name",
		"general_information_alias",
		"general_information_description",
		"general_information_slug",
		"general_information_category",
		"general_information_addresses",
		"general_information_unit",
		"general_information_phone",
		"general_information_email",
		"general_information_logo",
		"general_information_operational_hours",
		"general_information_link",
		"general_information_media",
		"general_information_social_media",
		"general_information_type",
	}).
		AddRow(
			mockDetailSpublicStruct[0].ID,
			helpers.GetStringFromObject(mockDetailSpublicStruct[0].Purpose),
			helpers.GetStringFromObject(mockDetailSpublicStruct[0].Facility),
			helpers.GetStringFromObject(mockDetailSpublicStruct[0].Requirement),
			helpers.GetStringFromObject(mockDetailSpublicStruct[0].ToS),
			helpers.GetStringFromObject(mockDetailSpublicStruct[0].InfoGraphic),
			helpers.GetStringFromObject(mockDetailSpublicStruct[0].FAQ),
			mockDetailSpublicStruct[0].CreatedAt,
			mockDetailSpublicStruct[0].UpdatedAt,
			mockDetailSpublicStruct[0].GeneralInformation.ID,
			mockDetailSpublicStruct[0].GeneralInformation.Name,
			mockDetailSpublicStruct[0].GeneralInformation.Alias,
			mockDetailSpublicStruct[0].GeneralInformation.Description,
			mockDetailSpublicStruct[0].GeneralInformation.Slug,
			mockDetailSpublicStruct[0].GeneralInformation.Category,
			helpers.GetStringFromObject(mockDetailSpublicStruct[0].GeneralInformation.Addresses),
			mockDetailSpublicStruct[0].GeneralInformation.Unit,
			helpers.GetStringFromObject(mockDetailSpublicStruct[0].GeneralInformation.Phone),
			mockDetailSpublicStruct[0].GeneralInformation.Email,
			mockDetailSpublicStruct[0].GeneralInformation.Logo,
			helpers.GetStringFromObject(mockDetailSpublicStruct[0].GeneralInformation.OperationalHours),
			helpers.GetStringFromObject(mockDetailSpublicStruct[0].GeneralInformation.Link),
			helpers.GetStringFromObject(mockDetailSpublicStruct[0].GeneralInformation.Media),
			helpers.GetStringFromObject(mockDetailSpublicStruct[0].GeneralInformation.SocialMedia),
			mockDetailSpublicStruct[0].GeneralInformation.Type,
		)

	query := `SELECT s.id, s.purpose, s.facility, s.requirement, s.tos, s.info_graphic, s.faq, s.created_at, s.updated_at,
	g.ID, g.name, g.alias, g.Description, g.slug, g.category, g.addresses, g.unit, g.phone, g.email, g.logo, g.operational_hours, g.link, g.media, g.social_media, g.type
	FROM service_public s
	JOIN general_informations g
	ON s.general_information_id = g.id
	WHERE 1=1`

	mock.ExpectQuery(query).WillReturnRows(rows)
	sp := spRepo.NewMysqlServicePublicRepository(db)

	params := &domain.Request{
		Keyword:   "",
		PerPage:   10,
		Offset:    0,
		SortBy:    "",
		SortOrder: "",
	}
	list, err := sp.Fetch(context.TODO(), params)

	// make an assertions using testify
	assert.NoError(t, err)
	assert.NotNil(t, list)
}
