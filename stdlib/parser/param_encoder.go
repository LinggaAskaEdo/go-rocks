package parser

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/linggaaskaedo/go-rocks/stdlib/preference/sqlx"
)

func (p *paramparser) InitEncoder() {
	p.encoder.RegisterEncoder(sqlx.NullString{}, encodesqlxNullString)
	p.encoder.RegisterEncoder(sqlx.NullBool{}, encodesqlxNullBool)
	p.encoder.RegisterEncoder(sqlx.NullInt64{}, encodesqlxNullInt64)
	p.encoder.RegisterEncoder(sqlx.NullFloat64{}, encodesqlxNullFloat64)
	p.encoder.RegisterEncoder(sqlx.NullTime{}, encodesqlxNullTime)
	p.encoder.RegisterEncoder(sqlx.NonStdTime{}, encodesqlxNullTime)
	p.encoder.RegisterEncoder(primitive.ObjectID{}, encodePrimitiveObjectID)
	p.encoder.RegisterEncoder(time.Time{}, encodeTime)
	p.encoder.RegisterEncoder(sqlx.NullID{}, encodesqlxNullID)
}

func encodesqlxNullString(v reflect.Value) string {
	nullString, ok := v.Interface().(sqlx.NullString)
	if !ok {
		return ""
	}

	if !nullString.Valid {
		return ""
	}

	return nullString.String
}

func encodesqlxNullBool(v reflect.Value) string {
	nullBool, ok := v.Interface().(sqlx.NullBool)
	if !ok {
		return ""
	}

	if !nullBool.Valid {
		return ""
	}

	return strconv.FormatBool(nullBool.Bool)
}

func encodesqlxNullInt64(v reflect.Value) string {
	nullInt, ok := v.Interface().(sqlx.NullInt64)
	if !ok {
		return ""
	}

	if !nullInt.Valid {
		return ""
	}

	return strconv.FormatInt(nullInt.Int64, 10)
}

func encodesqlxNullFloat64(v reflect.Value) string {
	nullFloat, ok := v.Interface().(sqlx.NullFloat64)
	if !ok {
		return ""
	}

	if !nullFloat.Valid {
		return ""
	}

	return fmt.Sprintf("%.2f", nullFloat.Float64)
}

func encodesqlxNullTime(v reflect.Value) string {
	nullTime, ok := v.Interface().(sqlx.NullTime)
	if !ok {
		return ""
	}

	if !nullTime.Valid {
		return ""
	}

	return nullTime.Time.Format(time.RFC3339)
}

func encodePrimitiveObjectID(v reflect.Value) string {
	objID, ok := v.Interface().(primitive.ObjectID)
	if !ok {
		return ""
	}

	if objID.IsZero() {
		return ""
	}

	return objID.Hex()
}

func encodeTime(v reflect.Value) string {
	currTime, ok := v.Interface().(time.Time)
	if !ok {
		return ""
	}

	if currTime.IsZero() {
		return ""
	}

	return currTime.Format(time.RFC3339)
}

func encodesqlxNullID(v reflect.Value) string {
	nullString, ok := v.Interface().(sqlx.NullID)
	if !ok {
		return ""
	}

	if !nullString.Valid {
		return ""
	}

	return nullString.ID
}
