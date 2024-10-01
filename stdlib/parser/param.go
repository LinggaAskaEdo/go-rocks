package parser

import (
	"reflect"

	"github.com/gorilla/schema"
	"github.com/rs/zerolog"
)

type ParamParser interface {
	Encode(src interface{}, dest map[string][]string) error
	Decode(dest interface{}, src map[string][]string) error
	RegisterDecoder(value interface{}, fn func(v reflect.Value) string)
	RegisterConverter(value interface{}, fn func(value string) reflect.Value)
}

type paramparser struct {
	logger  zerolog.Logger
	encoder *schema.Encoder
	decoder *schema.Decoder
	opt     ParamOptions
}

type ParamOptions struct {
	TagName           string
	ZeroEmpty         bool
	IgnoreUnknownKeys bool
}

func InitParamP(logger zerolog.Logger, opt ParamOptions) ParamParser {
	p := &paramparser{
		logger:  logger,
		encoder: schema.NewEncoder(),
		decoder: schema.NewDecoder(),
		opt:     opt,
	}

	p.InitDecoder()
	p.InitEncoder()

	if opt.TagName != "" {
		p.encoder.SetAliasTag(opt.TagName)
		p.decoder.SetAliasTag(opt.TagName)
	}

	p.decoder.IgnoreUnknownKeys(opt.IgnoreUnknownKeys)
	p.decoder.ZeroEmpty(opt.ZeroEmpty)

	return p
}

func (p *paramparser) Encode(src interface{}, dest map[string][]string) error {
	return p.encoder.Encode(src, dest)
}

func (p *paramparser) Decode(dest interface{}, src map[string][]string) error {
	return p.decoder.Decode(dest, src)
}

func (p *paramparser) RegisterConverter(value interface{}, fn func(value string) reflect.Value) {
	p.decoder.RegisterConverter(value, fn)
}

func (p *paramparser) RegisterDecoder(value interface{}, fn func(v reflect.Value) string) {
	p.encoder.RegisterEncoder(value, fn)
}
