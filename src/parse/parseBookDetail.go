package parse

import (
    "fmt"
    "pachong/src/engine"
    "pachong/src/model"
    "regexp"
    "strconv"
)

var authorReg = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
var name = regexp.MustCompile(`<span property="v:itemreviewed">([^<]+)</span>`)
var public = regexp.MustCompile(`<span class="pl">出版社:</span>([^<]+)<br/>`)
var page = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
var price = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br/>`)
var Score = regexp.MustCompile(`<strong class="ll rating_num " property="v:average"> ([^<]+) </strong>`)
var Info = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)

func ParseBookdDetail(contents []byte,bookName string) engine.ParseResult {
    book := &model.BookDetail{}
    book.Name = bookName
    book.Author = ExtraString(contents, authorReg)
    book.Publicer = ExtraString(contents, public)
    atoi, err := strconv.Atoi(ExtraString(contents, page))
    if err == nil {
        book.BookPages = atoi
    }
    book.Score = ExtraString(contents, Score)
    book.Price = ExtraString(contents, price)
    book.Info = ExtraString(contents, Info)
    fmt.Print( book.String())

    return engine.ParseResult{
        Items: []interface{}{book},
    }

}

func ExtraString(contents []byte, re *regexp.Regexp) string {
    submatch := re.FindSubmatch(contents)
    if len(submatch) >= 2 {
        return string(submatch[1])
    } else {
        return ""
    }
}
