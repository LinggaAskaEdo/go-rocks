package division

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/golang/snappy"

	"github.com/linggaaskaedo/go-rocks/src/business/entity"
	commonerr "github.com/linggaaskaedo/go-rocks/stdlib/errors/common"
	x "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"
)

func (d *division) setCacheDivision(ctx context.Context, param entity.DivisionParam, result []entity.Division, pagination entity.Pagination) error {
	// serialize query param to string
	rawKey, err := json.Marshal(param)
	if err != nil {
		return x.WrapWithCode(err, commonerr.CodeCacheMarshal, "set_cache_division")
	}

	field := string(rawKey)

	rawJSON, err := json.Marshal(result)
	if err != nil {
		return x.WrapWithCode(err, commonerr.CodeCacheMarshal, "set_cache_division")
	}

	// snappy compression on top up callback history
	var encJSON []byte
	encJSON = snappy.Encode(encJSON, rawJSON)

	// set key expiration
	if err := d.redis.HSet(ctx, divisionByParamHashKey, field, encJSON).Err(); err != nil {
		return x.WrapWithCode(err, commonerr.CodeCacheSetHashKey, "set_cache_division")
	}

	if err := d.redis.Expire(ctx, divisionByParamHashKey, durationDivisionExpiration).Err(); err != nil {
		return x.WrapWithCode(err, commonerr.CodeCacheSetExpiration, "set_cache_division")
	}

	rawJSON, err = json.Marshal(pagination)
	if err != nil {
		return x.WrapWithCode(err, commonerr.CodeCacheMarshal, "set_cache_division")
	}

	// snappy compression on pagination
	encJSON = []byte{}
	encJSON = snappy.Encode(encJSON, rawJSON)

	if err := d.redis.HSet(ctx, divisionPaginationByParamHashKey, field, encJSON).Err(); err != nil {
		return x.WrapWithCode(err, commonerr.CodeCacheSetHashKey, "set_cache_division")
	}

	if err := d.redis.Expire(ctx, divisionPaginationByParamHashKey, durationDivisionExpiration).Err(); err != nil {
		return x.WrapWithCode(err, commonerr.CodeCacheSetExpiration, "set_cache_division")
	}

	return nil
}

func (d *division) getCacheDivision(ctx context.Context, param entity.DivisionParam) ([]entity.Division, entity.Pagination, error) {
	var (
		results    []entity.Division
		pagination entity.Pagination
	)

	// serialize query param to string
	rawKey, err := json.Marshal(param)
	if err != nil {
		return results, pagination, x.WrapWithCode(err, commonerr.CodeCacheMarshal, "get_cache_division")
	}

	field := string(rawKey)

	// fetch transaction
	resultRaw, err := d.redis.HGet(ctx, divisionByParamHashKey, field).Bytes()
	if err == redis.Nil {
		return results, pagination, err
	} else if err != nil {
		return results, pagination, x.WrapWithCode(err, commonerr.CodeCacheGetHashKey, "get_cache_division")
	}

	var decJSON []byte
	decJSON, err = snappy.Decode(decJSON, resultRaw)
	if err != nil {
		return results, pagination, x.WrapWithCode(err, commonerr.CodeCacheDecode, "get_cache_division")
	}

	if err := json.Unmarshal(decJSON, &results); err != nil {
		return results, pagination, x.WrapWithCode(err, commonerr.CodeCacheUnmarshal, "get_cache_division")
	}

	// fetch pagination
	paginationRaw, err := d.redis.HGet(ctx, divisionPaginationByParamHashKey, field).Bytes()
	if err == redis.Nil {
		return results, pagination, err
	} else if err != nil {
		return results, pagination, x.WrapWithCode(err, commonerr.CodeCacheGetHashKey, "get_cache_division")
	}

	// decode pagination (encoded json)
	decJSON = []byte{}
	decJSON, err = snappy.Decode(decJSON, paginationRaw)
	if err != nil {
		return results, pagination, x.WrapWithCode(err, commonerr.CodeCacheDecode, "get_cache_division")
	}

	// unmarshaling returned byte
	if err := json.Unmarshal(decJSON, &pagination); err != nil {
		return results, pagination, x.WrapWithCode(err, commonerr.CodeCacheUnmarshal, "get_cache_division")
	}

	return results, pagination, nil
}
