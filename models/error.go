package models

type APPError struct {
	Error   error
	Message string
	Code    string
	Status  int
}
