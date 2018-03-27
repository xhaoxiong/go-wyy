package engin

import (
	"io"
)

type Request struct {
	Url           string
	ParserComment func([]byte) ParseResult
	ParserSong    func(reader io.Reader) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
