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
	e.POST(constant.ENDPOINT_LOGIN, configs.UserController.LoginUser)     // Login User
	e.POST(constant.ENDPOINT_REGISTER, configs.UserController.CreateUser) // Register User

	auth := e.Group("/auth", middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(configs.Cfg.SECRET_JWT),
		ContextKey: "token",
	}))

	//User Router With Auth
	user := auth.Group(constant.ENDPOINT_USERS)
	user.GET("", configs.UserController.GetAllUser, middlewares.UnauthorizedRole([]string{"User"})) // Get All User
	user.PUT(constant.ENDPOINT_WITH_ID, configs.UserController.UpdateUser)                          // Update User
	user.DELETE(constant.ENDPOINT_WITH_ID, configs.UserController.DeleteUser)                       // Delete User

	//Game Router With Auth
	game := auth.Group(constant.ENDPOINT_GAME)
	game.POST("", configs.GameController.CreateGame, middlewares.UnauthorizedRole([]string{"User"}))                          // Create Game
	game.PUT(constant.ENDPOINT_WITH_ID, configs.GameController.UpdateGame, middlewares.UnauthorizedRole([]string{"User"}))    // Update Game
	game.DELETE(constant.ENDPOINT_WITH_ID, configs.GameController.DeleteGame, middlewares.UnauthorizedRole([]string{"User"})) // Delete Game
	game.GET(constant.ENDPOINT_WITH_ID, configs.GameController.FindGameByID)                                                  // Find Game By ID
	game.GET("", configs.GameController.FindAllGame)                                                   // Find All Game

	//Coach Router With Auth
	coach := auth.Group(constant.ENDPOINT_COACH)
	coach.POST("", configs.CoachController.CreateCoach, middlewares.UnauthorizedRole([]string{"User"}))   // Create Coach
	coach.PUT(constant.ENDPOINT_WITH_ID, configs.CoachController.UpdateCoach, middlewares.UnauthorizedRole([]string{"User"}))    // Update Coach
	coach.DELETE(constant.ENDPOINT_WITH_ID, configs.CoachController.DeleteCoach, middlewares.UnauthorizedRole([]string{"User"})) // Delete Coach
	coach.GET(constant.ENDPOINT_WITH_ID, configs.CoachController.FindCoachByID)                                                  // Find Coach By ID
	coach.GET(constant.ENDPOINT_COACH_GETBYGAME, configs.CoachController.FindCoachByGameID)                                           // Find Coach By Game ID
	coach.GET(constant.ENDPOINT_COACH_GETBYCODE, configs.CoachController.FindCoachByCode)                                             // Find Coach By Code

	//Coach Availability Router With Auth
	coachAvailability := auth.Group(constant.ENDPOINT_COACH_AVAILABILITY)
	coachAvailability.POST("", configs.CoachAvailabilityController.CreateCoachAvailability, middlewares.UnauthorizedRole([]string{"User"}))   // Create Coach Availability
	coachAvailability.PUT(constant.ENDPOINT_WITH_ID, configs.CoachAvailabilityController.UpdateCoachAvailability, middlewares.UnauthorizedRole([]string{"User"}))    // Update Coach Availability
	coachAvailability.DELETE(constant.ENDPOINT_WITH_ID, configs.CoachAvailabilityController.DeleteCoachAvailability, middlewares.UnauthorizedRole([]string{"User"})) // Delete Coach Availability

	//Coach Experience Router With Auth
	coachExperience := auth.Group(constant.ENDPOINT_COACH_EXPERIENCE)
	coachExperience.POST("", configs.CoachExperienceController.CreateCoachExperience, middlewares.UnauthorizedRole([]string{"User"}))   // Create Coach Experience
	coachExperience.PUT(constant.ENDPOINT_WITH_ID, configs.CoachExperienceController.UpdateCoachExperience, middlewares.UnauthorizedRole([]string{"User"}))    // Update Coach Experience
	coachExperience.DELETE(constant.ENDPOINT_WITH_ID, configs.CoachExperienceController.DeleteCoachExperience, middlewares.UnauthorizedRole([]string{"User"})) // Delete Coach Experience

	//User Book Router With Auth
	userBook := auth.Group(constant.ENDPOINT_USER_BOOK)
	userBook.POST("", configs.UserBookController.CreateUserBook)                                                   // Create User Book
	userBook.PUT(constant.ENDPOINT_WITH_ID, configs.UserBookController.UpdateUserBook, middlewares.UnauthorizedRole([]string{"User"}))    // Update User Book
	userBook.DELETE(constant.ENDPOINT_WITH_ID, configs.UserBookController.DeleteUserBook, middlewares.UnauthorizedRole([]string{"User"})) // Delete User Book
	userBook.GET("", configs.UserBookController.FindAllUserBook, middlewares.UnauthorizedRole([]string{"User"}))         // Find All User Book
	userBook.GET(constant.ENDPOINT_USER_BOOK_GETBYUSERID, configs.UserBookController.FindUserBookByUserID)                                         // Find User Book By User ID
	userBook.GET(constant.ENDPOINT_WITH_ID, configs.UserBookController.FindUserBookById)                                                  // Find User Book By ID

	//User Payment Router With Auth
	userPayment := auth.Group(constant.ENDPOINT_USER_PAYMENT)
	userPayment.POST("", configs.UserPaymentController.CreateUserPayment)                                                   // Create User Payment
	userPayment.PUT(constant.ENDPOINT_WITH_ID, configs.UserPaymentController.UpdateUserPayment)                                                    // Update User Payment
	userPayment.DELETE(constant.ENDPOINT_WITH_ID, configs.UserPaymentController.DeleteUserPayment, middlewares.UnauthorizedRole([]string{"User"})) // Delete User Payment
	userPayment.GET("", configs.UserPaymentController.FindAllUserPayment, middlewares.UnauthorizedRole([]string{"User"}))         // Find All User Payment
	userPayment.GET(constant.ENDPOINT_WITH_ID, configs.UserPaymentController.FindUserPaymentById)                                                  // Find User Payment By User ID
	userPayment.GET(constant.ENDPOINT_USER_PAYMENT_GETBYINVOICE, configs.UserPaymentController.FindUserPaymentByInvoice)                                       // Find User Payment By Invoice

	return e
}
