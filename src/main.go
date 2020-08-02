package main

import (
    "pachong/src/engine"
    "pachong/src/parse"
)

func main() {
    engine.Run(engine.Request{
        Url: "https://book.douban.com/",
        ParseFunc: parse.ParseContent,
    })

}

/*func parseContene(content []byte){
    re := regexp.MustCompile(`<a href="([^"]+)" class="tag">([^"]+)</a>`)
    submatch := re.FindAllSubmatch(content, -1)
    for _,result:= range submatch{
        fmt.Printf("%s \n",result[0],result[1],result[2])
    }
}*/
