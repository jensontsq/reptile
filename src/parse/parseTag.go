package parse

import (
    "fmt"
    "pachong/src/engine"
    "regexp"
)

const regStr = `<a href="([^"]+)" class="tag">([^"]+)</a>`

func ParseContent(content []byte) engine.ParseResult {
    re := regexp.MustCompile(regStr)
    submatch := re.FindAllSubmatch(content, -1)
    result := engine.ParseResult{}
    for _, m := range submatch {
        fmt.Printf("m[0]:%s m[1]:%s m[3]:%s \n", m[0],m[1],m[2])
        result.Items = append(result.Items, m[2])
        result.Request = append(result.Request, engine.Request{
            Url:       "https://book.douban.com/" + string(m[1]),
            ParseFunc: ParseBookList,
        })
    }
    return result
}
