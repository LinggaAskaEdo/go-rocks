package dto

type UserDTO struct {
	PublicID     string `json:"id,omitempty" extensions:"x-order=0"`
	Username     string `json:"username,omitempty" extensions:"x-order=1"`
	Email        string `json:"email,omitempty" extensions:"x-order=2"`
	Phone        string `json:"phone,omitempty" extensions:"x-order=3"`
	HashPassword string `json:"hashPassword,omitempty" extensions:"x-order=4" swaggerignore:"true"`
	IsDeleted    bool   `json:"isDeleted" extensions:"x-order=5"`
	Message      string `json:"message,omitempty" extesnsions:"x-order=6"`
}

type UserLoginDTO struct {
	AccessToken  string `json:"accessToken,omitempty" extensions:"x-order=0"`
	RefreshToken string `json:"refreshToken,omitempty" extensions:"x-order=1"`
	AccessUUID   string `json:"-"`
	RefreshUUID  string `json:"-"`
	ExpiresAt    int64  `json:"expiresAt,omitempty" extensions:"x-order=2"`
	ExpiresRt    int64  `json:"expiresRt,omitempty" extensions:"x-order=3"`
	ExpiresIn    int    `json:"expiresIn,omitempty" extensions:"x-order=4"`
}

type UserLogoutDTO struct {
	Message string `json:"message,omitempty" extensions:"x-order=0"`
}
