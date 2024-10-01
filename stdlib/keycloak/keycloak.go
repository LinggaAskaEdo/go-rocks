package keycloak

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
	"github.com/rs/zerolog"

	"github.com/linggaaskaedo/go-rocks/stdlib/preference"
)

type Keycloak struct {
	Gocloak      *gocloak.GoCloak
	ClientId     string
	ClientSecret string
	Realm        string
}

type Options struct {
	URL          string
	Username     string
	Password     preference.EncryptedValue
	ClientID     string
	ClientSecret preference.EncryptedValue
	Realm        string
}

func Init(log zerolog.Logger, opt Options) *Keycloak {
	client := gocloak.NewClient(opt.URL)

	// Checking connection
	_, err := client.LoginAdmin(context.Background(), opt.Username, opt.Password.String(), "master")
	if err != nil {
		log.Panic().Err(err).Msg("Failed to connect to Keycloak !!!")
	}

	log.Debug().Msg(fmt.Sprintf("Keycloak status: OK"))

	return &Keycloak{
		Gocloak:      client,
		ClientId:     opt.ClientID,
		ClientSecret: opt.ClientSecret.String(),
		Realm:        opt.Realm,
	}
}
