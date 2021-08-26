package utils
import (
  "strings"
  "errors"
  "log"

  "github.com/Yousiph1/gurl/Colors/colors"
)
//ParseHeaders converts flag -H into a key: value map
func ParseHeaders(headers string) (map[string]string,error) {
  headersMap := make(map[string]string)
  keyValueArr := strings.Split(headers,",")
  for _,kv := range keyValueArr {
     kv = strings.TrimSpace(kv)
     keyValue := strings.Split(kv,":")
     if len(keyValue) != 2 {
       return nil, errors.New(colors.Yellow + "Invalid headers format, headers will be dropped" + colors.Reset)
     }else {
       key := strings.TrimSpace(keyValue[0])
       headersMap[key] = strings.TrimSpace(keyValue[1])
     }
  }
  return headersMap,nil
}

//ParseCookies converts flag -C into a key: value map
func ParseCookies(cookies string) (map[string]string, error) {

  cookiesMap := make(map[string]string)
  if cookies == "" {
    return cookiesMap,nil
  }
  keyValueArr := strings.Split(cookies,";")
  for _,kv := range keyValueArr {
     kv = strings.TrimSpace(kv)
     keyValue := strings.Split(kv,"=")
     if len(keyValue) != 2 {
       return nil, errors.New(colors.Yellow + "Invalid cookies format, all cookies will be dropped" + colors.Reset)
     }else {
       key := strings.TrimSpace(keyValue[0])
       headersMap[key] = strings.TrimSpace(keyValue[1])
     }
  }
  return cookiesMap,nil
}

func FailOnError(err error, msg string) {
  if err != nil {
    log.Fatalf(colors.Red + "%v: %s",err,"couldn't create a new request")
  }
}
