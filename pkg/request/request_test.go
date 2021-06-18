package request

import (
  "fmt"
  "net/http"
  "net/http/httptest"

  "github.com/ConsenSys/fc-retrieval-common/pkg/fcrmessages"
)

var c = NewHttpCommunicator()
type A struct {
  B string
}

func ExampleHttpCommunications_GetJSON_v01() {
  bytes, err := c.GetJSON("")
	fmt.Println(bytes, err)
	// Output: [] Get "": unsupported protocol scheme ""
}

func ExampleHttpCommunications_GetJSON_v02() {
	v02 := map[string]string{"k1":"v1"}
  _, err := c.GetJSON("")
	fmt.Println(v02, err)
	// Output: map[k1:v1] Get "": unsupported protocol scheme ""
}

func ExampleHttpCommunications_GetJSON_v03() {
  v03 := map[string]string{"k1":"v1"}
  _, err := c.GetJSON("http://127.0.0.1:80")
  fmt.Println(v03, err)
  // Output: map[k1:v1] Get "http://127.0.0.1:80": dial tcp 127.0.0.1:80: connect: connection refused
}

func ExampleHttpCommunications_SendMessage_v03() {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(404)
  }))
  defer ts.Close()
  tsHostPort := ts.URL[7:]
  var req fcrmessages.FCRMessage
  var resp *fcrmessages.FCRMessage
  resp, err := c.SendMessage(tsHostPort, &req)
  fmt.Println(resp, err)
  // Output:
  // <nil> SendMessage receive error code: 404 Not Found
}

func ExampleHttpCommunications_SendJSON_v01() {
  err := c.SendJSON("", 1)
  fmt.Println(err)
  // Output: Post "": unsupported protocol scheme ""
}



func ExampleHttpCommunications_SendMessage_v01() {
	var req fcrmessages.FCRMessage
	var resp *fcrmessages.FCRMessage
	resp, err := c.SendMessage("", &req)
	fmt.Println(resp, err)
	// Output: <nil> Post "http:///v1": http: no Host in request URL
}

// todo: James please fix or remove
//func ExampleGetJSON_03() {
//  v01 := make(map[string]string)
//
//  httpClient = utest.NewTestClientString(`{"`)
//  err := c.GetJSON("", &v01)
//  httpClient = &http.Client{}
//
//  fmt.Println(v01, err)
//  // Output: map[] unexpected EOF
//}
//
//func ExampleSendJSON_02() {
//  httpClient = utest.NewTestClientString(``)
//  err := c.SendJSON("", 1)
//  httpClient = &http.Client{}
//  fmt.Println(err)
//  // Output: <nil>
//}
//
//func ExampleSendMessage_02() {
//	var req fcrmessages.FCRMessage
//	var resp *fcrmessages.FCRMessage
//
//	httpClient = utest.NewTestClientString(`
//	{"message_type":1,
//	"protocol_version":2,
//	"protocol_supported":[3],
//	"message_body":[4],
//	"message_signature":""}`)
//	resp, err := c.SendMessage("0", &req)
//	httpClient = &http.Client{}
//
//	fmt.Println(resp, err)
//	// Output: &{1 2 [3] [4] } <nil>
//}
//
//
//
//func ExampleSendMessage_04() {
//	var req fcrmessages.FCRMessage
//	var resp *fcrmessages.FCRMessage
//
//	httpClient = utest.NewTestClientError()
//	resp, err := c.SendMessage("", &req)
//	httpClient = &http.Client{}
//
//	fmt.Println(resp, err)
//	// Output: &{0 0 [] [] } test error
//}
//
//func ExampleSendMessage_05() {
//	var req fcrmessages.FCRMessage
//	var resp *fcrmessages.FCRMessage
//
//	httpClient = utest.NewTestClientString(`{`)
//	resp, err := c.SendMessage("0", &req)
//	httpClient = &http.Client{}
//
//	fmt.Println(resp, err != nil)
//	// Output: <nil> true
//}
