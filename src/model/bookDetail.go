package model

import "strconv"

type BookDetail struct {
    Name      string
    Author    string
    Publicer  string
    BookPages int
    Price     string
    Score     string
    Info      string
}

func (b *BookDetail) String() string {
    return "书名：" + b.Name + " 作者：" + b.Author + "出版社：" + b.Publicer + "页数：" + strconv.Itoa(b.BookPages) + "价格：" + b.Price + "评分：" + b.Score + "简介：" + b.Info
}
