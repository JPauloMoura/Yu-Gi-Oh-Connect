package duelist

import (
	"errors"
	"testing"

	"github.com/JPauloMoura/Yu-Gi-Oh-Connect/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDuelistRepository struct {
	mock.Mock
}

func (o *mockDuelistRepository) CreateDuelist(duelist entities.Duelist) (*entities.Duelist, error) {
	args := o.Called(duelist)
	return args.Get(0).(*entities.Duelist), args.Error(1)
}

func Test_duelistService_CreateDuelist(t *testing.T) {
	invalidDuelist := entities.Duelist{}
	validDuelist := entities.NewDuelist()

	t.Run("should return error when duelist creation in the repository fails", func(t *testing.T) {
		mRepository := new(mockDuelistRepository)
		mRepository.On("CreateDuelist", invalidDuelist).Return(&entities.Duelist{}, errors.New("internal error"))

		svc := NewDuelistService(mRepository)
		duelist, err := svc.CreateDuelist(invalidDuelist)
		assert.Nil(t, duelist)
		assert.NotNil(t, err)
	})

	t.Run("should return the duelist entity when the creation is done successfully", func(t *testing.T) {
		mRepository := new(mockDuelistRepository)
		mRepository.On("CreateDuelist", validDuelist).Return(&validDuelist, nil)

		svc := NewDuelistService(mRepository)
		duelist, err := svc.CreateDuelist(validDuelist)
		assert.Equal(t, &validDuelist, duelist)
		assert.Nil(t, err)
	})
}
