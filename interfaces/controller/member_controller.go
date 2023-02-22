package controller

import (
	"emoshu_practice/interfaces/database"
	"emoshu_practice/usecase"
	"errors"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MemberController struct {
	MemberInteractor usecase.MemberInteractor
}

func NewMemberController(db database.DBHandler) *MemberController {
	return &MemberController{
		MemberInteractor: usecase.MemberInteractor{
			MemberRepository: &database.MemberRepository{
				DBHandler: db,
			},
		},
	}
}

// Show godoc
// @Summary  Show
// @Tags     Member
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} domain.Member
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router   /api/members/{id} [get]
func (controller *MemberController) Show(c echo.Context) error {
	member, err := controller.MemberInteractor.GetMemberById(c.Param("id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, NewError(404, err))
			return nil
		}
		c.JSON(500, NewError(500, err))
		return nil
	}
	c.JSON(200, member)
	return nil
}
