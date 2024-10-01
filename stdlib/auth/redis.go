package auth

import (
	"context"
	"fmt"
	"strings"

	x "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"
)

func (a *auth) saveToRedis(ctx context.Context, publicID string, td *TokenDetails) error {
	respAccess := a.redis.Set(ctx, td.AccessUUID, publicID, a.expiredToken)
	if respAccess.Err() != nil {
		return x.Wrap(respAccess.Err(), "store_redis_access_token")
	}

	respRefresh := a.redis.Set(ctx, td.RefreshUUID, publicID, a.expiredRefreshToken)
	if respRefresh.Err() != nil {
		return x.Wrap(respRefresh.Err(), "store_redis_refresh_token")
	}

	return nil
}

func (a *auth) deleteToRedis(ctx context.Context, authD *AccessDetails, isRefresh bool) error {
	var accessUUID, refreshUUID string

	// get uuid
	if !isRefresh {
		accessUUID = authD.AccessUUID
		refreshUUID = fmt.Sprintf("%s++%s", authD.AccessUUID, authD.UserID)
	} else {
		accessUUID = strings.ReplaceAll(authD.RefreshUUID, "++"+authD.UserID, "")
		refreshUUID = authD.RefreshUUID
	}

	// delete access token
	err := a.redis.Del(ctx, accessUUID).Err()
	if err != nil {
		return x.Wrap(err, "delete_redis_access_uuid")
	}

	// delete refresh token
	err = a.redis.Del(ctx, refreshUUID).Err()
	if err != nil {
		return x.Wrap(err, "delete_redis_refresh_uuid")
	}

	return nil
}
