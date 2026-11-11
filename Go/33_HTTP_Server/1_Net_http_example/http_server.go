/*

                ListenAndServe


• ListenAndServe function provided by net/http to start a HTTP server.
• func ListenAndServe(addr string, handler Handler) error
• ListenAndServe listens on the provided address and uses handler to handle requests.
• It will use the default serve mux if the provided handler is nil.


                  http.Handler


• type Handler interface {
      ServeHTTP(ResponseWriter, *Request)
  }
• A Handler responds to an HTTP request or handles the incoming requests for the server.


                 Handling Requests


• The handler should write reply headers and data to the ResponseWriter and then return.
• It is not valid to write to the ResponseWriter or Read from the request body after the request has been completed.
• Read from the request before writing response.
• Request handlers should not modify incoming requests.
• Panics inside handlers are logged to server logs.
• Panic with ErrAbortHandler if you want the client to know.


                http.ResponseWriter


• type ResponseWriter interface {
      Header() Header
      Write([]byte) (int, error)
      WriteHeader(statusCode int)
  }
• The method Header() returns a Header map that is set on the Response.
• The Write() method writes the data to the connection as part of an HTTP reply.
• The WriteHeader() method sends an HTTP response header with the provided status code.


                  http.Request


• The http.Request struct represents an HTTP request received by a server or to be sent by a client.
• Field Method of type string specifies the HTTP method (GET, POST, PUT, etc.). For client requests, an empty string means GET.
• URL *url.URL specifies either the URI being requested when used on the server side or the URL to access when used on the client side.
• Header of type map[string][]string contains the request header fields either received by the server or to be sent by the client.


                  Request.Body


• For client requests, a nil body means the request has no body, such as a GET request.
• For server requests, the Request Body is always non-nil but will return EOF immediately when no body is present.
• The Server will close the request body. The ServeHTTP Handler does not need to.
• Form url.Values contains the parsed form data, including both the URL field's query parameters and the PATCH, POST, or PUT form data.
• This field is only available after ParseForm is called.


Package net/http: [https://pkg.go.dev/net/http]

*/

package main

import (
	"fmt"
	"net/http"
)

// Step 1: Create a type that will act as our handler.
// An empty struct is enough, because this handler stores no data.
type HelloHandler struct{}

// Step 2: Add the ServeHTTP method.
// http.Handler is an interface with ONE method: ServeHTTP(ResponseWriter, *Request).
// Because HelloHandler has this method, it automatically IS an http.Handler
// (implicit interface, same idea as Shape and Area() earlier).
//
//	w = ResponseWriter: used to WRITE the reply back to the client
//	r = *Request:       holds everything the client SENT (method, URL, headers, body)
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Step 3: Write the response body. Write needs []byte, so convert the string.
	// Status code defaults to 200 OK because we did not call WriteHeader.
	w.Write([]byte("Hello!"))
}

func main() {

	// Step 4: Start the server.
	// ":8080" = listen on port 8080 on this machine.
	// new(HelloHandler) = a pointer to HelloHandler. Every request on every path
	// is handled by our ServeHTTP method.
	// ListenAndServe blocks (keeps running) until the server fails,
	// and it only returns an error, never a success value.
	if err := http.ListenAndServe(":8080", new(HelloHandler)); err != nil {
		fmt.Println(err)
	}
}

/*  How to test:

1. Run the program:        go run .
2. Open in the browser:    http://localhost:8080
   or use curl:            curl http://localhost:8080
3. Output:                 Hello!

(Any path like /abc or /hello also shows "Hello!", because there is only one handler.)

*/