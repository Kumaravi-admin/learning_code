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

func main() {

	// Step 1: Register a handler for the path "/hello" on the DEFAULT serve mux.
	// (A serve mux is a router: it matches the URL path to the right handler.)
	//
	// http.HandlerFunc turns a plain function into an http.Handler.
	// So we do not need to create a struct with a ServeHTTP method.
	// A simple function is enough.
	http.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Step 2: Write the reply. Write needs []byte, so convert the string.
		w.Write([]byte("Hello!"))
	}))

	// Step 3: Start the server on port 8080.
	// The handler is nil, so Go uses the DEFAULT serve mux.
	// That is the same mux where http.Handle registered "/hello" above.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}

}

/*  How to test:

1. Run the program:        go run .
2. Open in the browser:    http://localhost:8080/hello
   or use curl:            curl http://localhost:8080/hello
3. Output:                 Hello!

Any other path (like / or /abc) returns "404 page not found",
because only "/hello" is registered.

*/