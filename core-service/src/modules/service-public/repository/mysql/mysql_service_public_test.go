package mysql_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mysqlRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/service-public/repository/mysql"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	currentTime := time.Now()
	mockGenInfo := domain.DetailGeneralInformation{
		ID:          50,
		Name:        "UPTD Laboratorium Kesehatan Provinsi Jawa Barat",
		Alias:       "LABKESDA JABAR",
		Description: "Laboratorium Kesehatan provinsi Jawa Barat memberikan pelayanan pemeriksaan kesehatan dan juga pemeriksaan lingkungan, diantaranya Imunoserologi, Mikrobiologi, Patologi Klinik, dan Kesehatan Lingkungan. Informasi lebih lengkap seputar Layanan Labkesda Jabar, silakan kunjungi tautan berikut: https://www.labkesprovjabar.id/tentang-labkes/",
		Slug:        "uptd-laboratorium-kesehatan-provinsi-jawa-barat-50",
		Category:    "kesehatan",
		Addresses: []string{
			"Jl. Sederhana No. 3—5 Pasteur, Kecamatan Sukajadi, Kota Bandung, Jawa Barat, 40161",
			"Jl. Sederhana No. 3—5 Pasteur, Kecamatan Sukajadi, Kota Bandung, Jawa Barat, 40161",
		},
		Unit: "Dinas Kesehatan",
		Logo: "https://dvgddkosknh6r.cloudfront.net/staging/media/img/public-service/1667377353-download-(3).png",
		Type: "offline",
		Phone: []string{
			"(022) 2033918",
		},
		Email: "",
		Link: domain.Link{
			Website:    "http://labkesprovjabar.id/",
			GooglePlay: "",
			GoogleForm: "",
			AppStore:   "",
		},
		OperationalHours: []domain.OperationalHours{
			{
				Start: "07:30",
				End:   "16:00",
			},
			{
				Start: "07:30",
				End:   "16:00",
			},
		},
		Media: domain.Media{
			Video: "https://www.youtube.com/watch?v=Ud1v8kqBhlw&ab_channel=LABKESJABAR",
			Images: []string{
				"https://dvgddkosknh6r.cloudfront.net/staging/media/img/public-service/1667377353-download-(3).png",
			},
		},
		SocialMedia: domain.SocialMedia{
			Facebook: domain.NullString{
				String: "https://www.facebook.com/profile.php?id=100064779049587",
				Valid:  true,
			},
			Instagram: domain.NullString{
				String: "https://www.instagram.com/labkes.jabar_/",
				Valid:  true,
			},
			Twitter: domain.NullString{
				String: "https://www.instagram.com/labkes.jabar_/",
				Valid:  true,
			},
			Tiktok: domain.NullString{
				String: "https://www.tiktok.com/@labkesprovjabar",
				Valid:  true,
			},
			Youtube: domain.NullString{
				String: "https://www.youtube.com/channel/UCbKAMocZ1JRg4krr-MQvLrQ",
				Valid:  true,
			},
		},
	}

	mockDetailSpublic := []domain.DetailServicePublicResponse{
		{
			ID:                 1,
			GeneralInformation: mockGenInfo,
			Purpose: domain.Purpose{
				Title: "Gass",
				Items: []string{
					"Gas",
				},
			},
			Facility: domain.FacilityService{
				Title: "Gass",
				Items: []domain.ItemsFacility{
					{
						Title: "Gas",
						Image: "http://localhost:8080/images/1.png",
					},
					{
						Title: "Gas",
						Image: "http://localhost:8080/images/1.png",
					},
				},
			},
			Requirement: domain.Requirement{
				Title: "Gass",
				Items: []domain.Items{
					{
						Description: "Gas",
						Image:       "http://localhost:8080/images/1.png",
					},
				},
			},
			ToS: domain.TermsOfService{
				Title: "Gass",
				Items: []domain.Items{
					{
						Description: "Gas",
						Image:       "http://localhost:8080/images/1.png",
					},
					{
						Description: "Gas",
						Image:       "http://localhost:8080/images/1.png",
					},
				},
				Image: "dummy",
			},
			InfoGraphic: domain.InfoGraphic{
				Images: []string{
					"dummy",
				},
			},
			FAQ: domain.FAQ{
				Items: []domain.QuestionAnswer{
					{
						Question: "Gass",
						Answer:   "Gas",
					},
				},
			},
			UpdatedAt: currentTime,
			CreatedAt: currentTime,
		},
	}
	fmt.Println(mockDetailSpublic)

	mockSpublic := []domain.ListServicePublicResponse{
		{
			ID:          1,
			Slug:        "uptd-laboratorium-kesehatan-provinsi-jawa-barat-50",
			Name:        "UPTD Laboratorium Kesehatan Provinsi Jawa Barat",
			Logo:        "https://dvgddkosknh6r.cloudfront.net/staging/media/img/public-service/1667377353-download-(3).png",
			Description: "UPTD Laboratorium Kesehatan Provinsi Jawa Barat",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "slug", "name", "logo", "description"}).
		AddRow(mockSpublic[0].ID, mockSpublic[0].Slug, mockSpublic[0].Name, mockSpublic[0].Logo, mockSpublic[0].Description)

	query := `SELECT id, slug, name, logo FROM service_public`
	mock.ExpectQuery(query).WillReturnRows(rows)
	sp := mysqlRepo.NewMysqlServicePublicRepository(db)

	params := &domain.Request{
		Keyword:   "",
		PerPage:   10,
		Offset:    0,
		SortBy:    "",
		SortOrder: "",
	}

	list, err := sp.Fetch(context.TODO(), params)
	assert.NoError(t, err)
	assert.Len(t, list, 1)
}
