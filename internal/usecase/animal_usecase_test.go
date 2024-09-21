package usecase

import (
	"enakazoo-backend/internal/entity"
	"enakazoo-backend/internal/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	mockRepo := new(repository.MockAnimalRepository)
	mockAnimals := []entity.Animal{
		{ID: 1, Name: "Lion", Class: "mammal", Legs: 4},
		{ID: 2, Name: "Elephant", Class: "mammal", Legs: 4},
	}

	mockRepo.On("GetAll").Return(mockAnimals, nil)

	animalUsecase := NewAnimalUsecase(mockRepo)
	animals, err := animalUsecase.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, mockAnimals, animals)
	mockRepo.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	mockRepo := new(repository.MockAnimalRepository)
	mockAnimal := entity.Animal{ID: 1, Name: "Lion", Class: "mammal", Legs: 4}

	mockRepo.On("GetByID", 1).Return(mockAnimal, nil)

	animalUsecase := NewAnimalUsecase(mockRepo)
	animal, err := animalUsecase.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, mockAnimal, animal)
	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	mockRepo := new(repository.MockAnimalRepository)
	newAnimal := entity.CreateAnimal{Name: "Tiger", Class: "mammal", Legs: 4}

	mockRepo.On("Create", newAnimal).Return(nil)

	animalUsecase := NewAnimalUsecase(mockRepo)
	err := animalUsecase.Create(newAnimal)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockRepo := new(repository.MockAnimalRepository)
	updateAnimal := entity.Animal{ID: 1, Name: "Updated Lion", Class: "mammal", Legs: 4}

	mockRepo.On("Update", updateAnimal).Return(nil)

	animalUsecase := NewAnimalUsecase(mockRepo)
	err := animalUsecase.Update(updateAnimal)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockRepo := new(repository.MockAnimalRepository)

	mockRepo.On("Delete", 1).Return(nil)

	animalUsecase := NewAnimalUsecase(mockRepo)
	err := animalUsecase.Delete(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetByID_NotFound(t *testing.T) {
	mockRepo := new(repository.MockAnimalRepository)

	mockRepo.On("GetByID", 99).Return(entity.Animal{}, errors.New("not found"))

	animalUsecase := NewAnimalUsecase(mockRepo)
	_, err := animalUsecase.GetByID(99)

	assert.Error(t, err)
	assert.EqualError(t, err, "not found")
	mockRepo.AssertExpectations(t)
}
