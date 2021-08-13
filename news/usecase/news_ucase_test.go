package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/jabardigitalservice/portal-jabar-api/domain"
	"github.com/jabardigitalservice/portal-jabar-api/domain/mocks"
	ucase "github.com/jabardigitalservice/portal-jabar-api/news/usecase"
)

func TestFetch(t *testing.T) {
	mockNewsRepo := new(mocks.NewsRepository)
	mockNews := domain.News{
		Title:   "Hello",
		Content: "News",
	}

	mockListNews := make([]domain.News, 0)
	mockListNews = append(mockListNews, mockNews)

	t.Run("success", func(t *testing.T) {
		mockNewsRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int64")).Return(mockListNews, "next-cursor", nil).Once()

		u := ucase.NewNewsUsecase(mockNewsRepo, time.Second*2)
		num := int64(1)
		cursor := "12"
		list, nextCursor, err := u.Fetch(context.TODO(), cursor, num)
		cursorExpected := "next-cursor"
		assert.Equal(t, cursorExpected, nextCursor)
		assert.NotEmpty(t, nextCursor)
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListNews))

		mockNewsRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockNewsRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int64")).Return(nil, "", errors.New("Unexpexted Error")).Once()

		u := ucase.NewNewsUsecase(mockNewsRepo, time.Second*2)
		num := int64(1)
		cursor := "12"
		list, nextCursor, err := u.Fetch(context.TODO(), cursor, num)

		assert.Empty(t, nextCursor)
		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockNewsRepo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	mockNewsRepo := new(mocks.NewsRepository)
	mockNews := domain.News{
		Title:   "Hello",
		Content: "News",
	}

	t.Run("success", func(t *testing.T) {
		mockNewsRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockNews, nil).Once()

		mockNewsRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()

		u := ucase.NewNewsUsecase(mockNewsRepo, time.Second*2)

		err := u.Delete(context.TODO(), mockNews.ID)

		assert.NoError(t, err)
		mockNewsRepo.AssertExpectations(t)
	})
	t.Run("news-is-not-exist", func(t *testing.T) {
		mockNewsRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.News{}, nil).Once()

		u := ucase.NewNewsUsecase(mockNewsRepo, time.Second*2)

		err := u.Delete(context.TODO(), mockNews.ID)

		assert.Error(t, err)
		mockNewsRepo.AssertExpectations(t)
	})
	t.Run("error-happens-in-db", func(t *testing.T) {
		mockNewsRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.News{}, errors.New("Unexpected Error")).Once()

		u := ucase.NewNewsUsecase(mockNewsRepo, time.Second*2)

		err := u.Delete(context.TODO(), mockNews.ID)

		assert.Error(t, err)
		mockNewsRepo.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	mockNewsRepo := new(mocks.NewsRepository)
	mockNews := domain.News{
		Title:   "Hello",
		Content: "News",
		ID:      23,
	}

	t.Run("success", func(t *testing.T) {
		mockNewsRepo.On("Update", mock.Anything, &mockNews).Once().Return(nil)

		u := ucase.NewNewsUsecase(mockNewsRepo, time.Second*2)

		err := u.Update(context.TODO(), &mockNews)
		assert.NoError(t, err)
		mockNewsRepo.AssertExpectations(t)
	})
}
