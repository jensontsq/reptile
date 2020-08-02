package parse

import (
    "pachong/src/engine"
    "regexp"
)

const BookListReg = `<a href="([^"]+)" title="([^"]+)"`

func ParseBookList(content []byte) engine.ParseResult {
    reg := regexp.MustCompile(BookListReg)
    findSubmatch := reg.FindAllSubmatch(content, -1)
    result := engine.ParseResult{}
    for _, m := range findSubmatch {
        bookName:=string(m[2])
        result.Items = append(result.Items, m[2])
        result.Request = append(result.Request, engine.Request{
            Url:       string(m[1]),
            ParseFunc: func(bytes []byte) engine.ParseResult {
                return ParseBookdDetail(bytes,bookName)
            },
        })
    }
    return result
}
