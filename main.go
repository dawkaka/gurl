package main

import (
 "net/http"
 "fmt"
 "flag"
 "io/ioutil"
 "log"

 "github.com/Yousiph1/gurl/Flags/flags"
 "github.com/Yousiph1/gurl/Utils/utils"
)

//Gurl main struc

type Gurl struct {
  method,target,path,headers,cookies,data string;
  respLogLength int
}


func(g *Gurl) Dial() {
  client := &http.Client{}
  req,err := http.NewRequest(g.method,utils.Sanitize(g.target)+g.path,g.data)
  utils.FailOnError(err)
  headers,err := utils.ParseHeaders(g.headers)
  if err != nil {
    log.Println(err)
  }else {
  for key,value := range headers {
    req.Headers.Add(key,value)
  }
  }
  cookies,err := utils.ParseCookies
  if err != nil {
    log.Println(err)
  }else {
  for key,value := range cookies {
    req.Cookies.Add(key,value)
  }
  }
  resp,err := client.Do(req)

}

func main() {
  g := &Gurl{flags.Method,flags.Target,flags.Path,flags.Headers,
             flags.Cookies,flags.RespLogLength}
  g.Dial()
}
