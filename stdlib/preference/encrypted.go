package preference

import (
	"os"

	"github.com/linggaaskaedo/go-rocks/stdlib/util"
)

type EncryptedValue string

func (s *EncryptedValue) Decode(cfgValue string) error {
	if !util.RegexEncryptedValue.MatchString(cfgValue) {
		*s = EncryptedValue(cfgValue)
		return nil
	}

	var (
		secretKey = os.Getenv("SECRETKEY")
		encrypted = util.ExtractEncryptedValue(cfgValue)
	)

	key := util.HashingSHA256(secretKey)

	decrypted, err := util.Decrypt(encrypted, key[32:], "")
	if err != nil {
		return err
	}

	*s = EncryptedValue(decrypted)

	return nil
}

func (s *EncryptedValue) String() string {
	return string(*s)
}
