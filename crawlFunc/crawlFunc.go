package main

import (        
  "fmt"
  "net/http"
  "io"      
  "io/ioutil"     
  "time"
)

type Lang struct {
  Name string
  URL string
  bytes int64
  time time.Duration

}
func retrieve(uri string) ( int64, time.Duration) {  
                           
  now := time.Now()
  resp1, err := http.Get(uri)
  if err != nil { 
    t := time.Since(now)           
    return 0, t                
  }
  num,_ :=io.Copy(ioutil.Discard, resp1.Body)
  t := time.Since(now)
  return num, t
}

func pfunc(l Lang) {
    fmt.Println("name :    xyz")
    fmt.Println("URL is  %v:",l.URL)
      fmt.Println("number of lytes loded are %v",l.bytes)
      fmt.Println("time taken is : %v", l.time)
  }

func crawl(pfunc func(Lang), lang Lang){
  var ttime time.Duration
  url:= [3]string{"https://www.python.org/", "https://www.ruby-lang.org/en/","https://golang.org/"}
  for _,url_val := range url {
    lang.URL = url_val
    lang.bytes,lang.time = retrieve(url_val)
    ttime=ttime+lang.time
    pfunc(lang)

  }
  fmt.Println("total time to crawl 3 sites id %v", ttime)

}

func main() {
  
  var l Lang

      crawl(pfunc,l)

  }
