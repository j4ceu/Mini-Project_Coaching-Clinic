package constant

const (
	// Endpoint User
	ENDPOINT_USERS         = "/users"
	ENDPOINT_USER_REGISTER = "/register"
	ENDPOINT_USER_LOGIN    = "/login"
	ENDPOINT_USER_UPDATE   = "/users/:id" // :id adalah parameter yang akan diisi oleh user
	ENDPOINT_USER_DELETE   = "/users/:id" // :id adalah parameter yang akan diisi oleh user

	// Endpoint Game
	ENDPOINT_GAMES       = "/games"
	ENDPOINT_GAME_CREATE = "/game"
	ENDPOINT_GAME_UPDATE = "/game/:id" // :id adalah parameter yang akan diisi oleh user
	ENDPOINT_GAME_DELETE = "/game/:id" // :id adalah parameter yang akan diisi oleh user
	ENDPOINT_GAME_DETAIL = "/game/:id" // :id adalah parameter yang akan diisi oleh user

	// Endpoint Coach
	ENDPOINT_COACHES       = "/coaches"
	ENDPOINT_COACH_CREATE  = "/coach"
	ENDPOINT_COACH_UPDATE  = "/coach/:id"      // :id adalah parameter yang akan diisi oleh user
	ENDPOINT_COACH_DELETE  = "/coach/:id"      // :id adalah parameter yang akan diisi oleh user
	ENDPOINT_COACH_DETAIL  = "/coach/:id"      // :id adalah parameter yang akan diisi oleh user
	ENDPOINT_COACH_BY_GAME = "/coach/game/:id" // :id adalah parameter yang akan diisi oleh user

	// Endpoint Coach Availability
	ENDPOINT_COACH_AVAILABILITY_CREATE = "/coach/availability"
	ENDPOINT_COACH_AVAILABILITY_UPDATE = "/coach/availability/:id" // :id adalah parameter yang akan diisi oleh user
	ENDPOINT_COACH_AVAILABILITY_DELETE = "/coach/availability/:id" // :id adalah parameter yang akan diisi oleh user

	// Endpoint Coach Experience
	ENDPOINT_COACH_EXPERIENCE_CREATE = "/coach/experience"
	ENDPOINT_COACH_EXPERIENCE_UPDATE = "/coach/experience/:id" // :id adalah parameter yang akan diisi oleh user
	ENDPOINT_COACH_EXPERIENCE_DELETE = "/coach/experience/:id" // :id adalah parameter yang akan diisi oleh user

)
