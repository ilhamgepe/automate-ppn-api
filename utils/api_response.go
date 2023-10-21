package utils

type ApiResponse struct {
	Status  int
	Success bool
	Data    interface{}
	Error   struct {
		Code    string
		Message string
	}
}
