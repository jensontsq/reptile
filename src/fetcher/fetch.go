package fetcher

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func Fetch(url string) ([]byte, error) {
    client := http.Client{}
    request, err := http.NewRequest("GET", url, nil)
    request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36")
    resp, err := client.Do(request)
    if resp.StatusCode != http.StatusOK {
        fmt.Printf("Error status code:%d", resp.StatusCode)
    }
    if nil != err {
        fmt.Print(err)
    }
    all, err := ioutil.ReadAll(resp.Body)
    return all, err

}
