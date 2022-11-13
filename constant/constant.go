package constant

const (
	ENDPOINT_WITH_ID = "/:id"

	// Endpoint User
	ENDPOINT_USERS    = "/user"
	ENDPOINT_REGISTER = "/register"
	ENDPOINT_LOGIN    = "/login"

	// Endpoint Coach
	ENDPOINT_COACH            = "/coach"
	ENDPOINT_COACH_GETBYCODE = "/code/:code" // :code adalah parameter yang akan diisi oleh user
	ENDPOINT_COACH_GETBYGAME = "/game/:id"   // :id adalah parameter yang akan diisi oleh user

	// Endpoint Game
	ENDPOINT_GAME = "/game"

	// Endpoint Coach Availability
	ENDPOINT_COACH_AVAILABILITY = "/coach/availability"

	// Endpoint Coach Experience
	ENDPOINT_COACH_EXPERIENCE = "/coach/experience"

	// Endpoint User Book
	ENDPOINT_USER_BOOK             = "/user/book"
	ENDPOINT_USER_BOOK_GETBYUSERID = "/userid/:id" // :id adalah parameter yang akan diisi oleh user

	// Endpoint User Payment
	ENDPOINT_USER_PAYMENT              = "/user/payment"
	ENDPOINT_USER_PAYMENT_GETBYINVOICE = "/invoice/:invoiceNumber" // :id adalah parameter yang akan diisi oleh user

	// FIREBASE CONSTANT
	AuthURI = "https://accounts.google.com/o/oauth2/auth"
	TokenURI = "https://oauth2.googleapis.com/token"
	AuthProviderX509CertURL = "https://www.googleapis.com/oauth2/v1/certs"
	ClientX509CertURL = "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-b6rek%40coaching-clinic.iam.gserviceaccount.com"

)
