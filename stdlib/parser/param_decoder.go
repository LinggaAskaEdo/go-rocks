package parser

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/linggaaskaedo/go-rocks/stdlib/preference/sqlx"
)

func (p *paramparser) InitDecoder() {
	nullString, nullBool, nullInt64, nullFloat64, nullTime := sqlx.NullString{}, sqlx.NullBool{}, sqlx.NullInt64{}, sqlx.NullFloat64{}, sqlx.NullTime{}
	objectID := primitive.ObjectID{}
	p.decoder.RegisterConverter(nullString, convertsqlxNullString)
	p.decoder.RegisterConverter(nullBool, convertsqlxNullBool)
	p.decoder.RegisterConverter(nullInt64, convertsqlxNullInt64)
	p.decoder.RegisterConverter(nullFloat64, convertsqlxNullFloat64)
	p.decoder.RegisterConverter(nullTime, p.convertsqlxNullTime)
	p.decoder.RegisterConverter(sqlx.NonStdTime{}, p.convertsqlxNullTime)
	p.decoder.RegisterConverter(objectID, convertPrimitiveObjectID)
	p.decoder.RegisterConverter(sqlx.NullID{}, convertsqlxNullID)
}

func convertsqlxNullString(value string) reflect.Value {
	v := sqlx.NullString{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}

func convertsqlxNullBool(value string) reflect.Value {
	v := sqlx.NullBool{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}

func convertsqlxNullInt64(value string) reflect.Value {
	v := sqlx.NullInt64{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}

func convertsqlxNullFloat64(value string) reflect.Value {
	v := sqlx.NullFloat64{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}

func (p *paramparser) convertsqlxNullTime(value string) reflect.Value {
	v := sqlx.NullTime{}

	if t0, err := time.Parse(time.RFC3339, value); err == nil {
		v.Valid = true
		v.Time = t0
		return reflect.ValueOf(v)
	}

	if t0, err := time.Parse(`2006-01-02`, value); err == nil {
		v.Valid = true
		v.Time = t0
		return reflect.ValueOf(v)
	}

	if t0, err := time.Parse(`2006-01-02 15:04:05`, value); err == nil {
		v.Valid = true
		v.Time = t0
		return reflect.ValueOf(v)
	}

	if t0, err := time.Parse(`2006-01-02T15:04:05`, value); err == nil {
		v.Valid = true
		v.Time = t0
		return reflect.ValueOf(v)
	}

	if t0, err := time.Parse(`2006-01-02T15:04:05.000Z`, value); err == nil {
		v.Valid = true
		v.Time = t0
		return reflect.ValueOf(v)
	}

	if t0, err := time.Parse(`2006-01-02 15:04:05.000Z`, value); err == nil {
		v.Valid = true
		v.Time = t0
		return reflect.ValueOf(v)
	}

	if t0, err := time.Parse(`2006-01-02 15:04:05-07:00`, value); err == nil {
		v.Valid = true
		v.Time = t0
		return reflect.ValueOf(v)

	}

	if t0, err := time.Parse(`2006-01-02T15:04:05-07:00`, value); err == nil {
		v.Valid = true
		v.Time = t0
		return reflect.ValueOf(v)
	}

	if t0, err := time.Parse(`2006-01-02 15:04:05 -07:00 MST`, value); err == nil {
		v.Valid = true
		v.Time = t0
		return reflect.ValueOf(v)
	}

	if t0, err := time.Parse(`2006-01-02T15:04:05 -07:00 MST`, value); err == nil {
		v.Valid = true
		v.Time = t0
		return reflect.ValueOf(v)
	}

	if t0, err := time.Parse(`2006-01-02T15:04:05 -07:00MST`, value); err == nil {
		v.Valid = true
		v.Time = t0
		return reflect.ValueOf(v)
	}

	if t0, err := time.Parse(`2006-01-02 15:04:05 -07:00MST`, value); err == nil {
		v.Valid = true
		v.Time = t0
		return reflect.ValueOf(v)
	}

	return reflect.Value{}
}

func convertPrimitiveObjectID(value string) reflect.Value {
	v, err := primitive.ObjectIDFromHex(value)
	if err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}

func convertsqlxNullID(value string) reflect.Value {
	v := sqlx.NullID{}
	if err := v.Scan(value); err != nil {
		return reflect.Value{}
	}

	return reflect.ValueOf(v)
}
