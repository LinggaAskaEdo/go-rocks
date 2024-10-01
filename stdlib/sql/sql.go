package sql

import (
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/linggaaskaedo/go-rocks/stdlib/preference"
)

type Options struct {
	Enabled     bool
	Driver      string
	Host        string
	Port        string
	DB          string
	User        string
	Password    preference.EncryptedValue
	SSL         bool
	ConnOptions ConnOptions
}

type ConnOptions struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
	ConnMaxIdleTime int
}

func Init(log zerolog.Logger, opt Options) *sqlx.DB {
	if !opt.Enabled {
		return nil
	}

	driver, host, err := getURI(opt)
	if err != nil {
		log.Panic().Err(err).Msg(fmt.Sprintf("%s status: FAILED", strings.ToUpper(opt.Driver)))
	}

	db, err := sqlx.Connect(driver, host)
	if err != nil {
		log.Panic().Err(err).Msg(fmt.Sprintf("%s status: FAILED", strings.ToUpper(opt.Driver)))
	}

	log.Debug().Msg(fmt.Sprintf("%s status: OK", strings.ToUpper(opt.Driver)))

	db.SetMaxOpenConns(opt.ConnOptions.MaxOpenConns)
	db.SetMaxIdleConns(opt.ConnOptions.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(opt.ConnOptions.ConnMaxLifetime) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(opt.ConnOptions.ConnMaxIdleTime) * time.Minute)

	return db
}

func getURI(opt Options) (string, string, error) {
	switch opt.Driver {
	case preference.POSTGRES:
		ssl := `disable`
		if opt.SSL {
			ssl = `require`
		}

		return opt.Driver, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", opt.Host, opt.Port, opt.User, opt.Password, opt.DB, ssl), nil

	case preference.MYSQL:
		ssl := `false`
		if opt.SSL {
			ssl = `true`
		}

		return opt.Driver, fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?tls=%s&parseTime=%t", opt.User, opt.Password, opt.Host, opt.Port, opt.DB, ssl, true), nil

	default:
		return "", "", errors.New("DB Driver is not supported ")
	}
}
