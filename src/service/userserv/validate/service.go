package validate

import (
	"encoding/json"
	"mikaellemos.com.br/dload/src/model/user"

	"github.com/go-playground/validator/v10"
	"mikaellemos.com.br/dload/src/model"
)

func Validate(body []byte) (*user.User, *[]model.ApiError) {

	var _user user.User
	_ = json.Unmarshal(body, &_user)

	validate := validator.New()
	err := validate.Struct(_user)

	if err != nil {
		var errors []model.ApiError

		for _, er := range err.(validator.ValidationErrors) {
			var apiError = model.ApiError{Field: er.Field(), Error: er.Error()}
			errors = append(errors, apiError)
		}

		return &user.User{}, &errors
	}

	return &_user, nil
}
