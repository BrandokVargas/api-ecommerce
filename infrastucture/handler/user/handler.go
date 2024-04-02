package user

import (
	"net/http"

	"github.com/BrandokVargas/api-ecommerce/domain/user"
	"github.com/BrandokVargas/api-ecommerce/model"
	"github.com/labstack/echo/v4"
)

// EL HANDLER CONOCE EL USECASE QUE SE CONECTA CON EL DOMINIO
// PUNTO DE ENTRA QUE SOLO LOS LA RUTA VA PODER ACCEDER
type handler struct {
	useCase user.UseCase
}

func newHandler(useCase user.UseCase) *handler {
	return &handler{useCase: useCase}
}

func (h *handler) Create(c echo.Context) error {
	m := &model.User{}

	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.useCase.Create(m); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, m)
}

func (h *handler) GetAll(c echo.Context) error {
	users, err := h.useCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, users)
}
