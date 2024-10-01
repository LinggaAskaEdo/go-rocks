package user

import (
	"context"

	"github.com/linggaaskaedo/go-rocks/src/business/dto"
	"github.com/linggaaskaedo/go-rocks/src/business/entity"
	"github.com/linggaaskaedo/go-rocks/src/common"
)

func (u *user) CreateUser(ctx context.Context, userEntity entity.User) (dto.UserDTO, error) {
	var userDto dto.UserDTO

	result, err := u.user.CreateUser(ctx, userEntity)
	if err != nil {
		return userDto, err
	}

	userDto.PublicID = common.MixerEncode(result.ID)
	userDto.Username = result.Username
	userDto.Email = result.Email
	userDto.Phone = result.Phone
	userDto.IsDeleted = result.IsDeleted

	return userDto, nil
}

func (u *user) GetUserByUserID(ctx context.Context, userID int64) (dto.UserDTO, error) {
	var userDto dto.UserDTO

	result, err := u.user.GetUserByUserID(ctx, userID)
	if err != nil {
		return userDto, err
	}

	userDto.PublicID = common.MixerEncode(result.ID)
	userDto.Username = result.Username
	userDto.Email = result.Email
	userDto.Phone = result.Phone
	userDto.IsDeleted = result.IsDeleted

	return userDto, nil
}

func (u *user) GetUserByUsername(ctx context.Context, username string) (dto.UserDTO, error) {
	var userDto dto.UserDTO

	result, err := u.user.GetUserByUsername(ctx, username)
	if err != nil {
		return userDto, err
	}

	userDto.PublicID = common.MixerEncode(result.ID)
	userDto.Username = result.Username
	userDto.Email = result.Email
	userDto.Phone = result.Phone
	userDto.HashPassword = result.Password
	userDto.IsDeleted = result.IsDeleted

	return userDto, nil
}
