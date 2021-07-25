package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "encoding/json"
    "net"
    "os"
    "flag"
    "regexp"
    // regexp/syntax
)



type Message struct {
    ipversion string
    ipaddr string
}

func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    // fmt.Println(localAddr.IP)
    return localAddr.IP
}

func getHostname () string {
hostname, err := os.Hostname()
 if err != nil {
     panic(err)
 }
return hostname
}

// func getURL(url string) string {
func getURL(url string) string {

    // resp, err := http.Get("http://webcode.me")
    resp, err := http.Get(url)

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
  responseBody := string(body)
   // e := json.Unmarshal(body, &Message)
   // fmt.Println(tr)
   var data map[string]interface{}
    err = json.Unmarshal([]byte(responseBody), &data)
    if err != nil {
        panic(err)
    }
    fmt.Println(data["ip"])

    if err != nil {

        log.Fatal(err)
    }
    sbody := string(body)
    fmt.Println(string(body))
    return sbody
}

func IsValidUUID(uuid string) bool {
    r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
    return r.MatchString(uuid)
}

func main() {
  var nFlag = flag.Int("n", 1234, "help message for flag n")
  fmt.Println("flagvar has value ", nFlag)
  fmt.Println(len(os.Args), os.Args)
  if len(os.Args) < 2 {
    fmt.Println(len(os.Args), os.Args)
    // msg := fmt.Sprintf("Usage: %s uuid", os.Args[0])
    usage_msg := "Usage: iutil <uuid>"
    fmt.Println(usage_msg)
    example_msg := "Example: iutil 13c023b2-ed03-11eb-b237-00163ebb406c"
    fmt.Println(example_msg)
    return
  }
    uuid :=  os.Args[1]
    fmt.Printf("UUID %q match? %v \n", uuid, IsValidUUID(uuid))
    return
  //getURL("http://example.org")
  // url := "http://example.org"
  // url := "https://api.ipify.org?format=json"
  // url := "https://api.ipify.org"
  url := "https://api64.ipify.org?format=json"
  r := getURL(url)
  // j, e := json.Marshal(r)
  // a := json.Unmarshal([]byte(r), &ip)
  fmt.Println(r)
  b, err := json.Marshal(r)
  fmt.Println(b)
  fmt.Println(err)
  internet_ipaddr := GetOutboundIP()
  fmt.Println(internet_ipaddr)
  hostname := getHostname()
  fmt.Println(hostname)


}
