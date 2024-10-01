package rest

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const REGEXP_PHONE string = `^(\+62|62|0)8[1-9][0-9]{6,9}$`

func (data DivisionDataPayload) Validate() error {
	return validation.ValidateStruct(&data,
		// Division name can't be empty, length must between 2 and 50
		validation.Field(
			&data.Name,
			validation.Required,
			validation.NotNil,
			validation.Length(2, 50)),
	)
}

func (data UserDataPayload) Validate() error {
	return validation.ValidateStruct(&data,
		// Username can't be empty, length must between 8 and 50
		validation.Field(
			&data.Username,
			validation.Required,
			validation.NotNil,
			validation.Length(8, 50)),

		// Email can't ne empty and must be valid
		validation.Field(
			&data.Email,
			validation.Required,
			validation.NotNil,
			is.Email),

		// Phone can't be empty and must be valid
		validation.Field(
			&data.Phone,
			validation.Required,
			validation.NotNil,
			validation.Match(regexp.MustCompile(REGEXP_PHONE))),

		// Division ID can't be empty
		validation.Field(
			&data.DivisionID,
			validation.Required,
			validation.NotNil),

		// Password can't be empty, with minimum length is 8
		validation.Field(
			&data.Password,
			validation.Required,
			validation.NotNil,
			validation.Length(8, 13)),
	)
}

func (data UserLoginDataPayload) Validate() error {
	return validation.ValidateStruct(&data,
		// Username can't be empty, length must between 8 and 50
		validation.Field(
			&data.Username,
			validation.Required,
			validation.NotNil,
			validation.Length(4, 25)),

		// Password can't be empty, with minimum length is 8
		validation.Field(
			&data.Password,
			validation.Required,
			validation.NotNil,
			validation.Length(8, 13)),
	)
}

func (data UserLogoutDataPayload) Validate() error {
	return validation.ValidateStruct(&data,
		// Username can't be empty, length must between 8 and 50
		validation.Field(
			&data.Username,
			validation.Required,
			validation.NotNil,
			validation.Length(8, 50)),
	)
}

func (data UserRelogDataPayload) Validate() error {
	return validation.ValidateStruct(&data,
		// RefreshToken can't be empty
		validation.Field(
			&data.RefreshToken,
			validation.Required,
			validation.NotNil),
	)
}
