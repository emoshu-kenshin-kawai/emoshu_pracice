package controller

import (
	"emoshu_practice/domain"
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

// Index godoc
// @Summary  Index
// @Tags     Members
// @Produce  json
// @Success 200 {array} domain.Member
// @Failure 500 {object} Error
// @Router   /api/members [get]
func (controller *MemberController) Index(c echo.Context) error {
	members, err := controller.MemberInteractor.GetMembers()
	if err != nil {
		c.JSON(500, NewError(500, err))
		return nil
	}
	c.JSON(200, members)
	return nil
}

// Show godoc
// @Summary  Show
// @Tags     Member
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} domain.Member
// @Failure 404 {object} Error
// @Failure 500 {object} Error
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

// Create godoc
// @Summary  Create
// @Tags     Member
// @Produce  json
// @Param no query string true "Employee Number"
// @Param full_name query string false "Full Name"
// @Param profile_img query string true "Profile Image"
// @Param kana_name query string false "Kana Name"
// @Param start_date query string true "Employment Start Date"
// @Param end_date query string true "Emoployment End Date"
// @Param employment_status_id query uint false "Employment Status"
// @Param status_id query uint false "Status"
// @Success 201 {object} domain.Member
// @Failure 500 {object} Error
// @Router   /api/member [post]
func (controller *MemberController) Create(c echo.Context) error {
	member := domain.Member{}
	if err := c.Bind(&member); err != nil {
		c.JSON(500, NewError(500, err))
	}
	err := controller.MemberInteractor.CreateMember(member)
	if err != nil {
		c.JSON(500, NewError(500, err))
		return nil
	}
	c.JSON(201, member)
	return nil
}
