package engine

import (
    "log"
    "pachong/src/fetcher"
)

func Run(seeds ...Request) {
    var request []Request
    for _, r := range seeds {
        request = append(request, r)
    }
    for len(request) > 0 {
        r := request[0]
        request = request[1:]
        log.Printf("Fetching url:%s \n", r.Url)
        fetch, err := fetcher.Fetch(r.Url)
        if err != nil {
            log.Printf("异常：%s", err)
        }
        parseFunc := r.ParseFunc(fetch)
        request = append(request, parseFunc.Request...)
      /*  for _, item := range parseFunc.Items {
           /* fmt.Printf("Got item:%s \n", item)
        }*/
    }

}
