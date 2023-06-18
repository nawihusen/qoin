package constant

// Variabel for TimeLocation & MethodNotAllowed
const (
	TimeLocation     = "Asia/Jakarta"
	MethodNotAllowed = "Method Not Allowed"
)

// ResultError is struct
type ResultError struct {
	Code InternalError
	Err  error
}
