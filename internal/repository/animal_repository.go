package repository

import (
	"database/sql"
	"enakazoo-backend/internal/entity"
)

type AnimalRepository interface {
	GetAll() ([]entity.Animal, error)
	GetByID(id int) (entity.Animal, error)
	Create(animal entity.CreateAnimal) error
	Update(animal entity.Animal) error
	Delete(id int) error
}

type animalRepository struct {
	db *sql.DB
}

func NewAnimalRepository(db *sql.DB) AnimalRepository {
	return &animalRepository{db: db}
}

func (r *animalRepository) GetAll() ([]entity.Animal, error) {
	rows, err := r.db.Query("SELECT id, name, class, legs FROM animals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var animals []entity.Animal
	for rows.Next() {
		var animal entity.Animal
		if err := rows.Scan(&animal.ID, &animal.Name, &animal.Class, &animal.Legs); err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}

	return animals, nil
}

func (r *animalRepository) GetByID(id int) (entity.Animal, error) {
	var animal entity.Animal
	err := r.db.QueryRow("SELECT id, name, class, legs FROM animals WHERE id = ?", id).Scan(&animal.ID, &animal.Name, &animal.Class, &animal.Legs)
	if err != nil {
		return animal, err
	}
	return animal, nil
}

func (r *animalRepository) Create(animal entity.CreateAnimal) error {
	_, err := r.db.Exec("INSERT INTO animals (name, class, legs) VALUES (?, ?, ?)", animal.Name, animal.Class, animal.Legs)
	return err
}

func (r *animalRepository) Update(animal entity.Animal) error {
	_, err := r.db.Exec("UPDATE animals SET name = ?, class = ?, legs = ? WHERE id = ?", animal.Name, animal.Class, animal.Legs, animal.ID)
	return err
}

func (r *animalRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM animals WHERE id = ?", id)
	return err
}
