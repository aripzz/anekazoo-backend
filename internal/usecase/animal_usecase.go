package usecase

import (
	"enakazoo-backend/internal/entity"
	"enakazoo-backend/internal/repository"
)

type AnimalUsecase interface {
	GetAll() ([]entity.Animal, error)
	GetByID(id int) (entity.Animal, error)
	Create(animal entity.CreateAnimal) error
	Update(animal entity.Animal) error
	Delete(id int) error
}

type animalUsecase struct {
	animalRepo repository.AnimalRepository
}

func NewAnimalUsecase(r repository.AnimalRepository) AnimalUsecase {
	return &animalUsecase{animalRepo: r}
}

func (u *animalUsecase) GetAll() ([]entity.Animal, error) {
	return u.animalRepo.GetAll()
}

func (u *animalUsecase) GetByID(id int) (entity.Animal, error) {
	return u.animalRepo.GetByID(id)
}

func (u *animalUsecase) Create(animal entity.CreateAnimal) error {
	return u.animalRepo.Create(animal)
}

func (u *animalUsecase) Update(animal entity.Animal) error {
	return u.animalRepo.Update(animal)
}

func (u *animalUsecase) Delete(id int) error {
	return u.animalRepo.Delete(id)
}
