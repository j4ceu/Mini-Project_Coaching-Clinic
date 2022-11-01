package route

import (
	"Mini-Project_Coaching-Clinic/configs"
	"Mini-Project_Coaching-Clinic/constant"
	"Mini-Project_Coaching-Clinic/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	//Init Router
	e.GET(constant.ENDPOINT_USERS, configs.UserController.GetAllUser)          // Get All User
	e.POST(constant.ENDPOINT_USER_LOGIN, configs.UserController.LoginUser)     // Login User
	e.POST(constant.ENDPOINT_USER_REGISTER, configs.UserController.CreateUser) // Register User

	auth := e.Group("/auth", middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("jace"),
		ContextKey: "token",
	}))

	//User Router With Auth
	auth.PUT(constant.ENDPOINT_USER_UPDATE, configs.UserController.UpdateUser)    // Update User
	auth.DELETE(constant.ENDPOINT_USER_DELETE, configs.UserController.DeleteUser) // Delete User

	//Game Router With Auth
	auth.POST(constant.ENDPOINT_GAME_CREATE, configs.GameController.CreateGame, middlewares.UnauthorizedRole([]string{"User"}))   // Create Game
	auth.PUT(constant.ENDPOINT_GAME_UPDATE, configs.GameController.UpdateGame, middlewares.UnauthorizedRole([]string{"User"}))    // Update Game
	auth.DELETE(constant.ENDPOINT_GAME_DELETE, configs.GameController.DeleteGame, middlewares.UnauthorizedRole([]string{"User"})) // Delete Game
	auth.GET(constant.ENDPOINT_GAME_DETAIL, configs.GameController.FindGameByID, middlewares.UnauthorizedRole([]string{"User"}))  // Find Game By ID
	auth.GET(constant.ENDPOINT_GAMES, configs.GameController.FindAllGame, middlewares.UnauthorizedRole([]string{"User"}))         // Find All Game

	//Coach Router With Auth
	auth.POST(constant.ENDPOINT_COACH_CREATE, configs.CoachController.CreateCoach, middlewares.UnauthorizedRole([]string{"User"}))   // Create Coach
	auth.PUT(constant.ENDPOINT_COACH_UPDATE, configs.CoachController.UpdateCoach, middlewares.UnauthorizedRole([]string{"User"}))    // Update Coach
	auth.DELETE(constant.ENDPOINT_COACH_DELETE, configs.CoachController.DeleteCoach, middlewares.UnauthorizedRole([]string{"User"})) // Delete Coach
	auth.GET(constant.ENDPOINT_COACH_DETAIL, configs.CoachController.FindCoachByID)                                                  // Find Coach By ID
	auth.GET(constant.ENDPOINT_COACH_BY_GAME, configs.CoachController.FindCoachByGameID)                                             // Find Coach By Game ID

	return e
}
