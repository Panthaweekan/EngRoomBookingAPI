package api

import (
	"github.com/Panthaweekan/EngRoomBookingAPI/infrastructure"
	"github.com/Panthaweekan/EngRoomBookingAPI/internal/adaptor/handler"
	"github.com/Panthaweekan/EngRoomBookingAPI/internal/adaptor/repo"
	"github.com/Panthaweekan/EngRoomBookingAPI/internal/core/service"
	middlewares "github.com/Panthaweekan/EngRoomBookingAPI/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

const OAUTH_PREFIX = "/oauth"

func bindOauthRouter(router fiber.Router) {
	oauth := router.Group(OAUTH_PREFIX)

	accountRepo := repo.NewAccountRepo(infrastructure.DB)
	accountTypeRepo := repo.NewAccountTypeRepo(infrastructure.DB)
	organizationRepo := repo.NewOrganizationRepo(infrastructure.DB)
	accountService := service.NewAccountService(accountRepo, accountTypeRepo, organizationRepo)
	studentService := service.NewStudentService(repo.NewStudentRepo(infrastructure.DB))
	hdl := handler.NewOauthHandler(accountService, studentService)
	oauth.Get("/me", middlewares.AuthMiddleware(), hdl.GetUser)

	oauth.Post("", hdl.SignIn)
	oauth.Post("/signout", middlewares.AuthMiddleware(), hdl.Logout)
}
