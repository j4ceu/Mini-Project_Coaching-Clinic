package configs

import (
	"Mini-Project_Coaching-Clinic/controllers"
	"Mini-Project_Coaching-Clinic/helper"
	"Mini-Project_Coaching-Clinic/repositories"
	"Mini-Project_Coaching-Clinic/services"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_URI     string
	SECRET_JWT string
}

var Cfg *Config

//User
var userRepository repositories.UserRepositories
var userService services.UserService
var UserController controllers.UserController

//Game
var gameRepository repositories.GameRepositories
var gameService services.GameService
var GameController controllers.GameController

//Coach
var coachRepository repositories.CoachRepositories
var coachService services.CoachService
var CoachController controllers.CoachController

//Coach Availability
var coachAvailabilityRepository repositories.CoachAvailabilityRepositories
var coachAvailabilityService services.CoachAvailabilityService
var CoachAvailabilityController controllers.CoachAvailabilityController

//Coach Experience
var coachExperienceRepository repositories.CoachExperienceRepositories
var coachExperienceService services.CoachExperienceService
var CoachExperienceController controllers.CoachExperienceController

// User Book
var userBookRepository repositories.UserBookRepository
var userBookService services.UserBookService
var UserBookController controllers.UserBookController

// User Payment
var userPaymentRepository repositories.UserPaymentRepository
var userPaymentService services.UserPaymentService
var UserPaymentController controllers.UserPaymentController

func Init() {
	initConfig()
	helper.InitAppFirebase()
	initDatabase()
	initRepository()
	initService()
	initController()
}

func initConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

	cfg := &Config{
		DB_URI:     os.Getenv("DB_URI"),
		SECRET_JWT: os.Getenv("SECRET_JWT"),
	}

	Cfg = cfg
}

func initRepository() {
	userRepository = repositories.NewUserRepositories(DB)
	gameRepository = repositories.NewGameRepositories(DB)
	coachRepository = repositories.NewCoachRepositories(DB)
	coachAvailabilityRepository = repositories.NewCoachAvailabilityRepositories(DB)
	coachExperienceRepository = repositories.NewCoachExperienceRepositories(DB)
	userBookRepository = repositories.NewUserBookRepository(DB)
	userPaymentRepository = repositories.NewUserPaymentRepository(DB)
}

func initService() {
	userService = services.NewUserService(userRepository)
	gameService = services.NewGameService(gameRepository)
	coachService = services.NewCoachService(coachRepository, gameRepository, userRepository)
	coachAvailabilityService = services.NewCoachAvailabilityService(coachAvailabilityRepository)
	coachExperienceService = services.NewCoachExperienceService(coachExperienceRepository)
	userBookService = services.NewUserBookServices(userBookRepository, userPaymentRepository, coachAvailabilityRepository, coachRepository)
	userPaymentService = services.NewUserPaymentServices(userPaymentRepository, userRepository)

}

func initController() {
	UserController = controllers.NewUserController(userService)
	GameController = controllers.NewGameController(gameService)
	CoachController = controllers.NewCoachController(coachService)
	CoachAvailabilityController = controllers.NewCoachAvailabilityController(coachAvailabilityService)
	CoachExperienceController = controllers.NewCoachExperienceController(coachExperienceService)
	UserBookController = controllers.NewUserBookController(userBookService, userPaymentService)
	UserPaymentController = controllers.NewUserPaymentController(userPaymentService)
}
