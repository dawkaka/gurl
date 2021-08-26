package flags
import(
  "flag"
)
//Method : for flag -M
var Method string
//Host : for flag -T
var Host string
//Path : for flag -P
var Path string
//Headers : for flag -H
var Headers string
//Cookies : for flag -C
var Cookies string
//Data : for flag -D
var Data string
//RespLogLength : for flag -L
var RespLogLength int


func init() {
  // initialize and parse flag variables
  flag.StringVar(&Method,"M","GET","request method (GET,POST,HEAD,PUT,DELETE)")
  flag.StringVar(&Host,"T","","target host name")
  flag.StringVar(&Path,"P","/","request path")
  flag.StringVar(&Headers,"H","accepts: */*","request headers in the form; key:value,key:value")
  flag.StringVar(&Cookies,"C","","request cookies")
  flag.StringVar(&Data,"D","","data to send with with alias req body")
  flag.IntVar(&RespLogLength,"L",0,"number of response body chars to log")
  flag.Parse()
}
