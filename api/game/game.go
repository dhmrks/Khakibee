package game

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/khakibee/khakibee/api/logger"
	"gitlab.com/khakibee/khakibee/api/model"
	"gitlab.com/khakibee/khakibee/api/store"
	"gitlab.com/khakibee/khakibee/api/user"
)

func (g *GameHandler) CreateGame(c echo.Context) error {
	var gm model.Game
	req := new(createGameReq)

	if err := req.bind(c, &gm); err != nil {
		return logger.EchoHTTPError(http.StatusUnprocessableEntity, err)
	}

	if err := g.store.InsertGame(c.Request().Context(), gm); err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}

func (g *GameHandler) GetGames(c echo.Context) error {
	userID := c.Get(user.UserIDCtxKey)

	status := store.GameStatusAll
	if userID == user.UserIDLimitedAccess {
		status = store.GameStatusActive
	}

	games, err := g.store.SelectGames(c.Request().Context(), status)
	if err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, games)
}

func (g *GameHandler) GetGame(c echo.Context) error {
	var gm model.Game
	req := new(getGameReq)

	if err := req.bind(c, &gm); err != nil {
		return logger.EchoHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	game, err := g.store.SelectGameByID(c.Request().Context(), gm.GameID)
	if err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err.Error())
	}

	if game == nil {
		return logger.EchoHTTPError(http.StatusNotFound, "game not found")
	}

	return c.JSON(http.StatusOK, game)
}

func (g *GameHandler) EditGame(c echo.Context) error {
	var gm model.Game
	req := new(editGameReq)

	if err := req.bind(c, &gm); err != nil {
		return logger.EchoHTTPError(http.StatusUnprocessableEntity, err)
	}

	if ok, err := g.store.UpdateGame(c.Request().Context(), gm); err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err.Error())
	} else if !ok {
		return logger.EchoHTTPError(http.StatusNotFound, "game not found")
	}

	return c.NoContent(http.StatusNoContent)
}
