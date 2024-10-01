package sqlx

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
	"time"
)

// NullString is an alias for sql.NullString data type
type NullString sql.NullString

// Scan implements the Scanner interface for NullString
func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ns = NullString{s.String, false}
	} else {
		*ns = NullString{s.String, true}
	}

	return nil
}

func (ns NullString) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}

	return ns.String, nil
}

// NullBool is an alias for sql.NullBool data type
type NullBool sql.NullBool

// Scan implements the Scanner interface for NullBool
func (nb *NullBool) Scan(value interface{}) error {
	var b sql.NullBool
	if err := b.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nb = NullBool{b.Bool, false}
	} else {
		*nb = NullBool{b.Bool, true}
	}

	return nil
}

func (nb NullBool) Value() (driver.Value, error) {
	if !nb.Valid {
		return nil, nil
	}

	return nb.Bool, nil
}

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 sql.NullInt64

// Scan implements the Scanner interface for NullInt64
func (ni *NullInt64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NullInt64{i.Int64, false}
	} else {
		*ni = NullInt64{i.Int64, true}
	}

	return nil
}

func (ni NullInt64) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}

	return ni.Int64, nil
}

// NullFloat64 is an alias for sql.NullFloat64 data type
type NullFloat64 sql.NullFloat64

// Scan implements the Scanner interface for NullFloat64
func (nf *NullFloat64) Scan(value interface{}) error {
	var f sql.NullFloat64
	if err := f.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nf = NullFloat64{f.Float64, false}
	} else {
		*nf = NullFloat64{f.Float64, true}
	}

	return nil
}

func (nf NullFloat64) Value() (driver.Value, error) {
	if !nf.Valid {
		return nil, nil
	}

	return nf.Float64, nil
}

// NullTime is an alias for sql.NullTime data type
type NullTime sql.NullTime

// Scan implements the Scanner interface for NullTime
func (nt *NullTime) Scan(value interface{}) error {
	var t sql.NullTime
	if err := t.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nt = NullTime{t.Time, false}
	} else {
		*nt = NullTime{t.Time, true}
	}

	return nil
}

func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}

	return nt.Time, nil
}

// NonStdTime is an alias for time.Time data type when format is not standard
type NonStdTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan - Implementation of scanner for database/sql
func (t *NonStdTime) Scan(value interface{}) error {
	// Check nil value
	if value == nil {
		t.Time, t.Valid = time.Time{}, false
		return nil
	}

	// Try to parse as std format
	t.Time, t.Valid = value.(time.Time)
	if t.Valid {
		return nil
	}

	// Should be more strictly to check this type.
	var (
		vt  time.Time
		err error
	)
	if vt, err = time.Parse(time.RFC3339, string(value.([]byte))); err == nil {
		t.Time, t.Valid = vt, true
		return nil
	}
	if vt, err = time.Parse("2006-01-02 15:04:05", string(value.([]byte))); err == nil {
		t.Time, t.Valid = vt, true
		return nil
	}

	t.Time, t.Valid = time.Time{}, false
	return err
}

// Value - Implementation of valuer for database/sql
func (t NonStdTime) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}

	return t.Time, nil
}

// NullID is a custom type for the database ID column
// It accepts Int64 and string, always resulting String
// Implements SQL Scanner, JSON Marshaler, JSON Unmarshaler
type NullID struct {
	ID    string
	Valid bool // Valid is true if ID is not NULL
}

// Scan implements the Scanner interface for NullID
func (nid *NullID) Scan(value interface{}) error {
	switch value := value.(type) {
	case int64, float64:
		// convert to string
		return nid.Scan(fmt.Sprint(value))
	case []byte:
		// convert to string
		// https://github.com/go-sql-driver/mysql/issues/407
		return nid.Scan(string(value))
	case string:
		str := new(NullString)
		if err := str.Scan(value); err != nil {
			return err
		}
		*nid = NullID{ID: str.String, Valid: true}
	case nil:
		return nil
	default:
		return fmt.Errorf("invalid type %T for NullID.Scan; only support int64 & string", value)
	}

	return nil
}

func (nid NullID) Value() (driver.Value, error) {
	if !nid.Valid {
		return nil, nil
	}

	return nid.ID, nil
}
