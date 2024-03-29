package helper

import (
	"log"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func JsonResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func ValidationError(err error) []string {
	var msgErrors []string
	log.Print(err)

	for _, e := range err.(validator.ValidationErrors) {
		msgErrors = append(msgErrors, e.Error())
	}

	return msgErrors
}
