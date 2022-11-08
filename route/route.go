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
		SigningKey: []byte(configs.Cfg.SECRET_JWT),
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
	auth.GET(constant.ENDPOINT_COACH_GETBYGAME, configs.CoachController.FindCoachByGameID)                                           // Find Coach By Game ID
	auth.GET(constant.ENDPOINT_COACH_GETBYCODE, configs.CoachController.FindCoachByCode)                                             // Find Coach By Code

	//Coach Availability Router With Auth
	auth.POST(constant.ENDPOINT_COACH_AVAILABILITY_CREATE, configs.CoachAvailabilityController.CreateCoachAvailability, middlewares.UnauthorizedRole([]string{"User"}))   // Create Coach Availability
	auth.PUT(constant.ENDPOINT_COACH_AVAILABILITY_UPDATE, configs.CoachAvailabilityController.UpdateCoachAvailability, middlewares.UnauthorizedRole([]string{"User"}))    // Update Coach Availability
	auth.DELETE(constant.ENDPOINT_COACH_AVAILABILITY_DELETE, configs.CoachAvailabilityController.DeleteCoachAvailability, middlewares.UnauthorizedRole([]string{"User"})) // Delete Coach Availability

	//Coach Experience Router With Auth
	auth.POST(constant.ENDPOINT_COACH_EXPERIENCE_CREATE, configs.CoachExperienceController.CreateCoachExperience, middlewares.UnauthorizedRole([]string{"User"}))   // Create Coach Experience
	auth.PUT(constant.ENDPOINT_COACH_EXPERIENCE_UPDATE, configs.CoachExperienceController.UpdateCoachExperience, middlewares.UnauthorizedRole([]string{"User"}))    // Update Coach Experience
	auth.DELETE(constant.ENDPOINT_COACH_EXPERIENCE_DELETE, configs.CoachExperienceController.DeleteCoachExperience, middlewares.UnauthorizedRole([]string{"User"})) // Delete Coach Experience

	//User Book Router With Auth
	auth.POST(constant.ENDPOINT_USER_BOOK_CREATE, configs.UserBookController.CreateUserBook)                                                   // Create User Book
	auth.PUT(constant.ENDPOINT_USER_BOOK_UPDATE, configs.UserBookController.UpdateUserBook, middlewares.UnauthorizedRole([]string{"User"}))    // Update User Book
	auth.DELETE(constant.ENDPOINT_USER_BOOK_DELETE, configs.UserBookController.DeleteUserBook, middlewares.UnauthorizedRole([]string{"User"})) // Delete User Book
	auth.GET(constant.ENDPOINT_USER_BOOKS, configs.UserBookController.FindAllUserBook, middlewares.UnauthorizedRole([]string{"User"}))         // Find All User Book
	auth.GET(constant.ENDPOINT_USER_BOOK_GETBYUSERID, configs.UserBookController.FindUserBookByUserID)                                         // Find User Book By User ID
	auth.GET(constant.ENDPOINT_USER_BOOK_DETAIL, configs.UserBookController.FindUserBookById)                                                  // Find User Book By ID

	//User Payment Router With Auth
	auth.POST(constant.ENDPOINT_USER_PAYMENT_CREATE, configs.UserPaymentController.CreateUserPayment)                                                   // Create User Payment
	auth.PUT(constant.ENDPOINT_USER_PAYMENT_UPDATE, configs.UserPaymentController.UpdateUserPayment)                                                    // Update User Payment
	auth.DELETE(constant.ENDPOINT_USER_PAYMENT_DELETE, configs.UserPaymentController.DeleteUserPayment, middlewares.UnauthorizedRole([]string{"User"})) // Delete User Payment
	auth.GET(constant.ENDPOINT_USER_PAYMENTS, configs.UserPaymentController.FindAllUserPayment, middlewares.UnauthorizedRole([]string{"User"}))         // Find All User Payment
	auth.GET(constant.ENDPOINT_USER_PAYMENT_DETAIL, configs.UserPaymentController.FindUserPaymentById)                                                  // Find User Payment By User ID
	auth.GET(constant.ENDPOINT_USER_PAYMENT_GETBYINVOICE, configs.UserPaymentController.FindUserPaymentByInvoice)                                       // Find User Payment By Invoice

	return e
}
