package jet

import (
	"github.com/go-jet/jet/v2/generator/mysql"
	"github.com/rs/zerolog"

	"github.com/linggaaskaedo/go-rocks/stdlib/preference"
)

type Options struct {
	Enabled  bool
	User     string
	Password preference.EncryptedValue
	Host     string
	Port     int
	DB       string
	Path     string
}

func Init(log zerolog.Logger, opt Options) {
	if opt.Enabled {
		err := mysql.Generate(opt.Path,
			mysql.DBConnection{
				Host:     opt.Host,
				Port:     opt.Port,
				User:     opt.User,
				Password: opt.Password.String(),
				DBName:   opt.DB,
			})

		if err != nil {
			log.Panic().Err(err).Msg("Generating table with Jet failed !!!")
		}

		log.Debug().Msg("Generating table with Jet success !!!")
	}
}
