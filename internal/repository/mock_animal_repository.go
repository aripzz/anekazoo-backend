package repository

import (
	"enakazoo-backend/internal/entity"

	"github.com/stretchr/testify/mock"
)

type MockAnimalRepository struct {
	mock.Mock
}

func (m *MockAnimalRepository) GetAll() ([]entity.Animal, error) {
	args := m.Called()
	return args.Get(0).([]entity.Animal), args.Error(1)
}

func (m *MockAnimalRepository) GetByID(id int) (entity.Animal, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Animal), args.Error(1)
}

func (m *MockAnimalRepository) Create(animal entity.CreateAnimal) error {
	args := m.Called(animal)
	return args.Error(0)
}

func (m *MockAnimalRepository) Update(animal entity.Animal) error {
	args := m.Called(animal)
	return args.Error(0)
}

func (m *MockAnimalRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
