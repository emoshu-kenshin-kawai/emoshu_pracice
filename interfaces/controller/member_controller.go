package controller

import (
	"emoshu_practice/interfaces/database"
	"emoshu_practice/usecase"
	"errors"

	"github.com/labstack/echo"
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
