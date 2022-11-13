package configs

import (
	"Mini-Project_Coaching-Clinic/controllers/CoachAvailability_Controller"
	"Mini-Project_Coaching-Clinic/controllers/CoachExperience_Controller"
	"Mini-Project_Coaching-Clinic/controllers/Coach_Controller"
	"Mini-Project_Coaching-Clinic/controllers/Game_Controller"
	"Mini-Project_Coaching-Clinic/controllers/UserBook_Controller"
	"Mini-Project_Coaching-Clinic/controllers/UserPayment_Controller"
	"Mini-Project_Coaching-Clinic/controllers/User_Controller"
	"Mini-Project_Coaching-Clinic/helper"
	"Mini-Project_Coaching-Clinic/repositories/CoachAvailability_Repository"
	"Mini-Project_Coaching-Clinic/repositories/CoachExperience_Repository"
	"Mini-Project_Coaching-Clinic/repositories/Coach_Repository"
	"Mini-Project_Coaching-Clinic/repositories/Game_Repository"
	"Mini-Project_Coaching-Clinic/repositories/UserBook_Repository"
	"Mini-Project_Coaching-Clinic/repositories/UserPayment_Repository"
	"Mini-Project_Coaching-Clinic/repositories/User_Repository"
	"Mini-Project_Coaching-Clinic/services/CoachAvailability_Service"
	"Mini-Project_Coaching-Clinic/services/CoachExperience_Service"
	"Mini-Project_Coaching-Clinic/services/Coach_Service"
	"Mini-Project_Coaching-Clinic/services/Game_Service"
	"Mini-Project_Coaching-Clinic/services/UserBook_Service"
	"Mini-Project_Coaching-Clinic/services/UserPayment_Service"
	"Mini-Project_Coaching-Clinic/services/User_Service"

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
var userRepository User_Repository.UserRepositories
var userService User_Service.UserService
var UserController User_Controller.UserController

//Game
var gameRepository Game_Repository.GameRepositories
var gameService Game_Service.GameService
var GameController Game_Controller.GameController

//Coach
var coachRepository Coach_Repository.CoachRepositories
var coachService Coach_Service.CoachService
var CoachController Coach_Controller.CoachController

//Coach Availability
var coachAvailabilityRepository CoachAvailability_Repository.CoachAvailabilityRepositories
var coachAvailabilityService CoachAvailability_Service.CoachAvailabilityService
var CoachAvailabilityController CoachAvailability_Controller.CoachAvailabilityController

//Coach Experience
var coachExperienceRepository CoachExperience_Repository.CoachExperienceRepositories
var coachExperienceService CoachExperience_Service.CoachExperienceService
var CoachExperienceController CoachExperience_Controller.CoachExperienceController

// User Book
var userBookRepository UserBook_Repository.UserBookRepository
var userBookService UserBook_Service.UserBookService
var UserBookController UserBook_Controller.UserBookController

// User Payment
var userPaymentRepository UserPayment_Repository.UserPaymentRepository
var userPaymentService UserPayment_Service.UserPaymentService
var UserPaymentController UserPayment_Controller.UserPaymentController

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
	userRepository = User_Repository.NewUserRepositories(DB)
	gameRepository = Game_Repository.NewGameRepositories(DB)
	coachRepository = Coach_Repository.NewCoachRepositories(DB)
	coachAvailabilityRepository = CoachAvailability_Repository.NewCoachAvailabilityRepositories(DB)
	coachExperienceRepository = CoachExperience_Repository.NewCoachExperienceRepositories(DB)
	userBookRepository = UserBook_Repository.NewUserBookRepository(DB)
	userPaymentRepository = UserPayment_Repository.NewUserPaymentRepository(DB)
}

func initService() {
	userService = User_Service.NewUserService(userRepository)
	gameService = Game_Service.NewGameService(gameRepository)
	coachService = Coach_Service.NewCoachService(coachRepository, gameRepository, userRepository)
	coachAvailabilityService = CoachAvailability_Service.NewCoachAvailabilityService(coachAvailabilityRepository)
	coachExperienceService = CoachExperience_Service.NewCoachExperienceService(coachExperienceRepository)
	userBookService = UserBook_Service.NewUserBookServices(userBookRepository, userPaymentRepository, coachAvailabilityRepository, coachRepository)
	userPaymentService = UserPayment_Service.NewUserPaymentServices(userPaymentRepository, userRepository)

}

func initController() {
	UserController = User_Controller.NewUserController(userService)
	GameController = Game_Controller.NewGameController(gameService)
	CoachController = Coach_Controller.NewCoachController(coachService)
	CoachAvailabilityController = CoachAvailability_Controller.NewCoachAvailabilityController(coachAvailabilityService)
	CoachExperienceController = CoachExperience_Controller.NewCoachExperienceController(coachExperienceService)
	UserBookController = UserBook_Controller.NewUserBookController(userBookService, userPaymentService)
	UserPaymentController = UserPayment_Controller.NewUserPaymentController(userPaymentService)
}
