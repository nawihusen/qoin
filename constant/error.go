package constant

import (
	"github.com/gofiber/fiber/v2"
)

// InternalError is declare type InternalError
type InternalError int

// Is the variable response status
const (
	StatusBadRequestErrorValidation            = 4001 //StatusBadRequestErrorValidation
	StatusBadRequestErrorParsingJSON           = 4002 //StatusBadRequestErrorParsingJSON
	StatusBadRequestBodyMultipartForm          = 4003 //StatusBadRequestBodyMultipartForm
	StatusBadRequestExceededFileSize           = 4004 //StatusBadRequestExceededFileSize
	StatusBadRequestNotExists                  = 4005 //StatusBadRequestNotExists
	StatusBadRequestExists                     = 4006 //StatusBadRequestExists
	StatusBadRequestExceededLimitSupporter     = 4007 //StatusBadRequestExceededLimitSupporter
	StatusUnauthorizedMemberIsNotRegistered    = 4011 //StatusUnauthorizedMemberIsNotRegistered
	StatusUnauthorizedBacalegIsRegistered      = 4012 //StatusUnauthorizedBacalegIsRegistered
	StatusUnauthorizedTokenExpired             = 4013 //StatusUnauthorizedTokenExpired
	StatusForbiddenInvalidToken                = 4031 //StatusForbiddenInvalidToken
	StatusNotFound                             = 4041 //StatusNotFound
	StatusMethodNotAllowed                     = 4051 //StatusMethodNotAllowed
	StatusInternalServerErrorDatabaseMysql     = 5001 //StatusInternalServerErrorDatabaseMysql
	StatusInternalServerErrorApps              = 5002 //StatusInternalServerErrorApps
	StatusInternalServerErrorDatabaseRedis     = 5003 //StatusInternalServerErrorDatabaseRedis
	StatusInternalServerErrorServiceMemberApps = 5004 //StatusInternalServerErrorServiceMemberApps
	StatusInternalServerErrorServiceAuthApps   = 5005 //StatusInternalServerErrorServiceAuthApps
)

// errorInfo is struct for HTTPCode
type errorInfo struct {
	HTTPCode    int
	Title       string
	UserMessage UserMessage
}

// UserMessage is struct for UserMessage
type UserMessage struct {
	En string `json:"en"`
	ID string `json:"id"`
}

// constantError is logic for constantError
var constantError = map[InternalError]errorInfo{
	StatusUnauthorizedMemberIsNotRegistered: {
		HTTPCode: fiber.StatusUnauthorized,
		Title:    "Member is not registered",
		UserMessage: UserMessage{
			En: "Member is not registered, please register first!",
			ID: "Member belum registrasi, silahkan register terlebih dahulu!",
		},
	},
	StatusUnauthorizedBacalegIsRegistered: {
		HTTPCode: fiber.StatusUnauthorized,
		Title:    "Bacaleg is registered",
		UserMessage: UserMessage{
			En: "Bacaleg is registered, you cannot register again!",
			ID: "Bacaleg sudah teregistrasi, kamu tidak bisa registrasi lagi!",
		},
	},
	StatusUnauthorizedTokenExpired: {
		HTTPCode: fiber.StatusUnauthorized,
		Title:    "Token is expired",
		UserMessage: UserMessage{
			En: "Token is expired, please do login",
			ID: "Token sudah kadaluarsa, silahkan lakukan login",
		},
	},
	StatusForbiddenInvalidToken: {
		HTTPCode: fiber.StatusForbidden,
		Title:    "Invalid token",
		UserMessage: UserMessage{
			En: "Invalid token, please try again",
			ID: "Token tidak sah, silahkan dicoba lagi",
		},
	},
	StatusBadRequestErrorValidation: {
		HTTPCode: fiber.StatusBadRequest,
		Title:    "The request is invalid",
		UserMessage: UserMessage{
			En: "The user request is invalid, please try again",
			ID: "Request user tidak sah, silahkan dicoba lagi",
		},
	},
	StatusBadRequestErrorParsingJSON: {
		HTTPCode: fiber.StatusBadRequest,
		Title:    "The json request is invalid",
		UserMessage: UserMessage{
			En: "The json request is invalid, please try again",
			ID: "Request json tidak sah, silahkan dicoba lagi",
		},
	},
	StatusBadRequestBodyMultipartForm: {
		HTTPCode: fiber.StatusBadRequest,
		Title:    "The multipart form is invalid",
		UserMessage: UserMessage{
			En: "The multipart form is invalid, please try again",
			ID: "Form multipart tidak sah, silahkan dicoba lagi",
		},
	},
	StatusBadRequestExceededFileSize: {
		HTTPCode: fiber.StatusBadRequest,
		Title:    "The file size is exceeded limit",
		UserMessage: UserMessage{
			En: "The file size is exceeded limit, please try again",
			ID: "Ukuran berkas melebihi batas maksimal, silahkan dicoba lagi",
		},
	},
	StatusBadRequestNotExists: {
		HTTPCode: fiber.StatusBadRequest,
		Title:    "The data is not exists",
		UserMessage: UserMessage{
			En: "The data is not exists",
			ID: "Data tidak ada",
		},
	},
	StatusBadRequestExists: {
		HTTPCode: fiber.StatusBadRequest,
		Title:    "The data is exists",
		UserMessage: UserMessage{
			En: "The data is exists",
			ID: "Data sudah ada",
		},
	},
	StatusBadRequestExceededLimitSupporter: {
		HTTPCode: fiber.StatusBadRequest,
		Title:    "The total supporter is exceeded limit",
		UserMessage: UserMessage{
			En: "The total supporter bacaleg is exceeded limit",
			ID: "Jumlah supporter bacaleg melebihi limitasi",
		},
	},
	StatusNotFound: {
		HTTPCode: fiber.StatusNotFound,
		Title:    "The endpoint is not found",
		UserMessage: UserMessage{
			En: "The endpoint is not found, please try again",
			ID: "Endpoint tidak ditemukan, silahkan dicoba lagi",
		},
	},
	StatusMethodNotAllowed: {
		HTTPCode: fiber.StatusMethodNotAllowed,
		Title:    "The method is not allowed",
		UserMessage: UserMessage{
			En: "The method is not allowed on this endpoint",
			ID: "Method tidak diizinkan untuk endpoint ini",
		},
	},
	StatusInternalServerErrorDatabaseMysql: {
		HTTPCode: fiber.StatusInternalServerError,
		Title:    "Internal Server Error ",
		UserMessage: UserMessage{
			En: "There is some problem with server in Database Mysql",
			ID: "Ada masalah dengan server database Mysql",
		},
	},
	StatusInternalServerErrorApps: {
		HTTPCode: fiber.StatusInternalServerError,
		Title:    "Internal Server Error Apps",
		UserMessage: UserMessage{
			En: "There is some problem with Apps, Please try again",
			ID: "Ada masalah dengan apps, silahkan dicoba lagi",
		},
	},
	StatusInternalServerErrorDatabaseRedis: {
		HTTPCode: fiber.StatusInternalServerError,
		Title:    "Internal Server Error",
		UserMessage: UserMessage{
			En: "There is some problem with server",
			ID: "Ada masalah dengan server",
		},
	},
	StatusInternalServerErrorServiceMemberApps: {
		HTTPCode: fiber.StatusInternalServerError,
		Title:    "Internal Server Error Service",
		UserMessage: UserMessage{
			En: "There is some problem with other service, Please try again",
			ID: "Ada masalah dengan service lain, silahkan dicoba lagi",
		},
	},
	StatusInternalServerErrorServiceAuthApps: {
		HTTPCode: fiber.StatusInternalServerError,
		Title:    "Internal Server Error Service",
		UserMessage: UserMessage{
			En: "There is some problem with other service, please try again",
			ID: "Ada masalah dengan service lain, silahkan dicoba lagi",
		},
	},
}

// Info for InternalError
func (i InternalError) Info() errorInfo {
	return constantError[i]
}
