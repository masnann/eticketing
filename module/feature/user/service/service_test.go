package service

import (
	"errors"
	"eticketing/module/entities"
	userMocks "eticketing/module/feature/user/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTestService(t *testing.T) (
	*UserService,
	*userMocks.UserRepositoryInterface) {

	repo := userMocks.NewUserRepositoryInterface(t)
	service := NewUserService(repo)

	return service.(*UserService), repo
}

func TestUserService_GetUsersById(t *testing.T) {
	userID := uint64(1)
	user := &entities.UserModels{
		ID:   userID,
		Name: "User 1",
		Role: "customer",
	}
	service, repo := setupTestService(t)
	t.Run("Failed Case - User Not Found", func(t *testing.T) {
		expectedErr := errors.New("pengguna tidak ditemukan")
		repo.On("GetUserByID", userID).Return(nil, expectedErr).Once()

		result, err := service.GetUserByID(userID)

		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)
		repo.AssertExpectations(t)
	})

	t.Run("Success Case", func(t *testing.T) {
		repo.On("GetUserByID", userID).Return(user, nil)

		result, err := service.GetUserByID(userID)

		assert.NoError(t, err)
		assert.NotNil(t, result)

		repo.AssertExpectations(t)
	})

}
