package rest

type DivisionCreateRequest struct {
	Data CreateDivisionData `json:"data"`
}

type CreateDivisionData struct {
	Division *DivisionDataPayload `json:"division,omitempty"`
}

type DivisionDataPayload struct {
	Name string `json:"name"`
}

type UserCreateRequest struct {
	Data CreateUserData `json:"data"`
}

type CreateUserData struct {
	User *UserDataPayload `json:"user,omitempty"`
}

type UserDataPayload struct {
	Username   string `json:"username" extensions:"x-order=0"`
	Email      string `json:"email" extensions:"x-order=1"`
	Phone      string `json:"phone" extensions:"x-order=2"`
	DivisionID string `json:"divisionID" extensions:"x-order=3"`
	Password   string `json:"password" extensions:"x-order=4"`
}

type UserLoginRequest struct {
	Data LoginUserData `json:"data"`
}

type LoginUserData struct {
	User *UserLoginDataPayload `json:"user,omitempty"`
}

type UserLoginDataPayload struct {
	Username string `json:"username" example:"ahmad" extensions:"x-order=0"`
	Password string `json:"password" example:"sahroni" extensions:"x-order=1"`
}

type UserLogoutRequest struct {
	Data LogoutUserData `json:"data"`
}

type LogoutUserData struct {
	User *UserLogoutDataPayload `json:"user,omitempty"`
}

type UserLogoutDataPayload struct {
	Username string `json:"username"`
}

type UserRelogRequest struct {
	Data RelogUserData `json:"data"`
}

type RelogUserData struct {
	User *UserRelogDataPayload `json:"user,omitempty"`
}

type UserRelogDataPayload struct {
	RefreshToken string `json:"refreshToken"`
}
