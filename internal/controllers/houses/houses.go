package houses

import (
	"net/http"

	"github.com/PatrickChagastavares/game-of-thrones/internal/entities"
	"github.com/PatrickChagastavares/game-of-thrones/internal/services"
	"github.com/PatrickChagastavares/game-of-thrones/pkg/httpRouter"
	"github.com/PatrickChagastavares/game-of-thrones/pkg/logger"
	"github.com/PatrickChagastavares/game-of-thrones/pkg/tracer"
)

type (
	IController interface {
		Create(c httpRouter.Context)
		Find(c httpRouter.Context)
		FindByID(c httpRouter.Context)
		Update(c httpRouter.Context)
		Delete(c httpRouter.Context)
	}
	controllers struct {
		srv *services.Container
		log logger.Logger
	}
)

func New(srv *services.Container, log logger.Logger) IController {
	return &controllers{srv: srv, log: log}
}

// house swagger document
// @Description Create one house
// @Tags house
// @Accept json
// @Produce json
// @Param house body entities.HouseRequest true "create new house"
// @Success 201
// @Failure 400 {object} entities.HttpErr
// @Failure 409 {object} entities.HttpErr
// @Failure 500
// @Security ApiKeyAuth
// @Router /houses [post]
func (ctrl *controllers) Create(c httpRouter.Context) {
	ctx, span := tracer.Span(c.Context(), "controllers.houses.create")
	defer span.End()

	var newHouse entities.HouseRequest
	if err := c.Decode(&newHouse); err != nil {
		c.JSON(http.StatusBadRequest, entities.ErrDecode)
		return
	}

	if err := c.Validate(newHouse); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	id, err := ctrl.srv.House.Create(ctx, newHouse)
	if err != nil {
		ctrl.log.Error("Ctrl.Create: ", "Error on create house: ", newHouse)
		responseErr(ctx, err, c.JSON)
		return
	}

	c.JSON(http.StatusCreated, map[string]any{
		"id": id,
	})
}

// house swagger document
// @Description Find houses
// @Tags house
// @Accept json
// @Produce json
// @Param	name	query	string	false	"name house"
// @Success 200 {object} []entities.House
// @Failure 500
// @Security ApiKeyAuth
// @Router /houses [get]
func (ctrl *controllers) Find(c httpRouter.Context) {
	ctx, span := tracer.Span(c.Context(), "controllers.houses.find")
	defer span.End()

	name := c.GetQuery("name")

	houses, err := ctrl.srv.House.Find(ctx, name)
	if err != nil {
		ctrl.log.Error("Ctrl.Find: ", "Error on find houses: ", name)
		responseErr(ctx, err, c.JSON)
		return
	}

	c.JSON(http.StatusOK, houses)
}

// house swagger document
// @Description find house by id
// @Tags house
// @Accept json
// @Produce json
// @Param id path string true "House ID"
// @Success 200 {object} entities.House
// @Failure 500
// @Security ApiKeyAuth
// @Router /houses/:id [get]
func (ctrl *controllers) FindByID(c httpRouter.Context) {
	ctx, span := tracer.Span(c.Context(), "controllers.houses.findbyid")
	defer span.End()

	id := c.GetParam("id")

	houses, err := ctrl.srv.House.FindByID(ctx, id)
	if err != nil {
		ctrl.log.Error("Ctrl.FindByID: ", "Error on find house: ", id)
		responseErr(ctx, err, c.JSON)
		return
	}

	c.JSON(http.StatusOK, houses)
}

// house swagger document
// @Description Update house
// @Tags house
// @Accept json
// @Produce json
// @Param id path string true "House ID"
// @Param house body entities.HouseRequest true "create new house"
// @Success 200 {object} entities.House
// @Failure 400 {object} entities.HttpErr
// @Failure 500
// @Security ApiKeyAuth
// @Router /houses/:id [put]
func (ctrl *controllers) Update(c httpRouter.Context) {
	ctx, span := tracer.Span(c.Context(), "controllers.houses.update")
	defer span.End()

	var updateHouse entities.HouseRequest
	if err := c.Decode(&updateHouse); err != nil {
		c.JSON(http.StatusBadRequest, entities.ErrDecode)
		return
	}

	if err := c.Validate(updateHouse); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	updateHouse.ID = c.GetParam("id")

	houses, err := ctrl.srv.House.Update(ctx, updateHouse)
	if err != nil {
		ctrl.log.Error("Ctrl.Update: ", "Error on update house: ", updateHouse)
		responseErr(ctx, err, c.JSON)
		return
	}

	c.JSON(http.StatusOK, houses)
}

// house swagger document
// @Description Delete house
// @Tags house
// @Accept json
// @Produce json
// @Param id path string true "House ID"
// @Param house body entities.HouseRequest true "create new house"
// @Success 200 {object} entities.House
// @Failure 400 {object} entities.HttpErr
// @Failure 500
// @Security ApiKeyAuth
// @Router /houses/:id [delete]
func (ctrl *controllers) Delete(c httpRouter.Context) {
	ctx, span := tracer.Span(c.Context(), "controllers.houses.delete")
	defer span.End()

	id := c.GetParam("id")

	err := ctrl.srv.House.Delete(ctx, id)
	if err != nil {
		ctrl.log.Error("Ctrl.Delete: ", "Error on delete house: ", id)
		responseErr(ctx, err, c.JSON)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
