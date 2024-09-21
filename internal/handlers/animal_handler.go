package handlers

import (
	"database/sql"
	"enakazoo-backend/internal/entity"
	"enakazoo-backend/internal/helpers"
	"enakazoo-backend/internal/repository"
	"enakazoo-backend/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AnimalHandler struct {
	animalUsecase usecase.AnimalUsecase
}

func NewAnimalHandler(app *fiber.App, db *sql.DB) {
	repo := repository.NewAnimalRepository(db)
	usecase := usecase.NewAnimalUsecase(repo)
	handler := &AnimalHandler{animalUsecase: usecase}

	apiv1 := app.Group("/api/v1")

	apiv1.Post("/animals", handler.Create)
	apiv1.Get("/animals", handler.GetAll)
	apiv1.Get("/animals/:id", handler.GetByID)
	apiv1.Put("/animals/:id", handler.Update)
	apiv1.Delete("/animals/:id", handler.Delete)
}

// @Summary Get all animals
// @Description Get a list of all animals
// @Tags animals
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Animal
// @Router /api/v1/animals [get]
func (h *AnimalHandler) GetAll(c *fiber.Ctx) error {
	animals, err := h.animalUsecase.GetAll()
	if err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SendResponse(c, fiber.StatusOK, "successfully", animals)
}

// @Summary Get an animal by ID
// @Description Get a single animal by ID
// @Tags animals
// @Accept  json
// @Produce  json
// @Param id path int true "Animal ID"
// @Success 200 {object} entity.Animal
// @Router /api/v1/animals/{id} [get]
func (h *AnimalHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid ID")
	}
	animal, err := h.animalUsecase.GetByID(id)
	if err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SendResponse(c, fiber.StatusOK, "successfully", animal)
}

// @Summary Create a new animal
// @Description Add a new animal to the database
// @Tags animals
// @Accept  json
// @Produce  json
// @Param animal body entity.CreateAnimal true "Animal data"
// @Success 201 {string} string "Animal created"
// @Router /api/v1/animals [post]
func (h *AnimalHandler) Create(c *fiber.Ctx) error {
	var animal entity.CreateAnimal
	if err := c.BodyParser(&animal); err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if err := h.animalUsecase.Create(animal); err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SendResponse(c, fiber.StatusCreated, "successfully", nil)
}

// @Summary Update an existing animal
// @Description Update an animal's data
// @Tags animals
// @Accept  json
// @Produce  json
// @Param id path int true "Animal ID"
// @Param animal body entity.Animal true "Animal data"
// @Success 200 {string} string "Animal updated"
// @Router /api/v1/animals/{id} [put]
func (h *AnimalHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid ID")
	}

	var animal entity.Animal
	if err := c.BodyParser(&animal); err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	animal.ID = id
	if err := h.animalUsecase.Update(animal); err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SendResponse(c, fiber.StatusOK, "update successfully", nil)
}

// @Summary Delete an animal
// @Description Delete an animal by ID
// @Tags animals
// @Accept  json
// @Produce  json
// @Param id path int true "Animal ID"
// @Success 200 {string} string "Animal deleted"
// @Router /api/v1/animals/{id} [delete]
func (h *AnimalHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid ID")
	}
	if err := h.animalUsecase.Delete(id); err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SendResponse(c, fiber.StatusOK, "deleted successfully", nil)
}
