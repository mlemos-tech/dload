package user

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"mikaellemos.com.br/dload/src/model"
)

func ValidateUser(body []byte) (model.User, []model.ApiError) {

	var user model.User
	_ = json.Unmarshal(body, &user)

	validate := validator.New()
	err := validate.Struct(user)

	if err != nil {
		var errors []model.ApiError

		for _, er := range err.(validator.ValidationErrors) {
			var apiError = model.ApiError{Field: er.Field(), Error: er.Error()}
			errors = append(errors, apiError)
		}

		return model.User{}, errors
	}

	return user, nil
}
