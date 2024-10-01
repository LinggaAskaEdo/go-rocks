package parser

import "github.com/rs/zerolog"

type Parser interface {
	ParamParser() ParamParser
}

type parser struct {
	param ParamParser
	opt   Options
}

type Options struct {
	Param ParamOptions
}

func Init(logger zerolog.Logger, opt Options) Parser {
	return &parser{
		param: InitParamP(logger, opt.Param),
		opt:   opt,
	}
}

func (p *parser) ParamParser() ParamParser {
	return p.param
}
