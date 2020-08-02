package engine

type ParseResult struct {
    Request []Request
    Items   []interface{}
}

type Request struct {
    Url       string
    ParseFunc func([]byte) ParseResult
}

func NilParse([]byte) ParseResult {
    return ParseResult{}
}

