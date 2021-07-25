package main

import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "net/http"
)

type ipifyAPIResp struct {
        IPv4Addr string `json:"ip"`
}

func getIpify(body []byte) (*ipifyAPIResp, error) {
        var s = new(ipifyAPIResp)
        err := json.Unmarshal(body, &s)
        if err != nil {
                fmt.Println("whoops:", err)
        }
        return s, err
}

func getURLJSON() string {
        resp, err := http.Get("https://api.ipify.org?format=json")
        // resp, err := http.Get(url)
        if err != nil {
                // handle err
        }
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                panic(err.Error())
        }

        s, err := getIpify([]byte(body))

        // fmt.Printf("%+v\n", *s)
        fmt.Printf("%+v\n", s.IPv4Addr)
        fmt.Printf(s.IPv4Addr)
        return s.IPv4Addr
}

func main() {
        //getURLJSON("https://api.ipify.org?format=json")
        getURLJSON()
}
