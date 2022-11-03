package configs

import (
	"Mini-Project_Coaching-Clinic/controllers"
	"Mini-Project_Coaching-Clinic/repositories"
	"Mini-Project_Coaching-Clinic/services"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB_URI string
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

func Init() {
	initConfig()
	initDatabase()
	initRepository()
	initService()
	initController()
}

func initConfig() {
	cfg := &Config{}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatal(err)
	}

	Cfg = cfg
}

func initRepository() {
	userRepository = repositories.NewUserRepositories(DB)
	gameRepository = repositories.NewGameRepositories(DB)
	coachRepository = repositories.NewCoachRepositories(DB)
	coachAvailabilityRepository = repositories.NewCoachAvailabilityRepositories(DB)
	coachExperienceRepository = repositories.NewCoachExperienceRepositories(DB)
}

func initService() {
	userService = services.NewUserService(userRepository)
	gameService = services.NewGameService(gameRepository)
	coachService = services.NewCoachService(coachRepository, gameRepository, userRepository)
	coachAvailabilityService = services.NewCoachAvailabilityService(coachAvailabilityRepository)
	coachExperienceService = services.NewCoachExperienceService(coachExperienceRepository)
}

func initController() {
	UserController = controllers.NewUserController(userService)
	GameController = controllers.NewGameController(gameService)
	CoachController = controllers.NewCoachController(coachService)
	CoachAvailabilityController = controllers.NewCoachAvailabilityController(coachAvailabilityService)
	CoachExperienceController = controllers.NewCoachExperienceController(coachExperienceService)
}
