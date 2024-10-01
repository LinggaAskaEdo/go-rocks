package common

import (
	"strings"

	"github.com/foolin/mixer"

	x "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"
)

const (
	key = "3m35eS"
)

func MixerEncode(data int64) string {
	return strings.ToLower(mixer.EncodeNumber(key, data))
}

func MixerDecode(data string) (int64, error) {
	decodePaddingData, err := mixer.DecodeNumber(key, strings.ToUpper(data))
	if err != nil {
		return decodePaddingData, x.Wrap(err, "decode_id")
	}

	return decodePaddingData, nil
}
