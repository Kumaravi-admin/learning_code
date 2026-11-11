/*
                  ServeMux


• ServeMux is the router provided by Go standard lib net/http package.
• ServeMux a route multiplexer allows you to match different request handlers to different URLs served by your server.
• Registers patterns with their respective handlers.

                  Patterns


• Patterns are the URL patterns that ServeMux will accept.
• The general pattern is of form: `[METHOD ][HOST]/[PATH]`.
• `/hello` matches the path `/hello`.
• `GET /hello` matches only a GET request.
• `GET /hello/` matches only GET requests for all paths that begin with `/hello/`.
• `GET /hello/{user_id}/message` will match the path with segment `/hello` while the `{user}` denotes the path param.


                  Wildcards


• The `{user_id}` wildcard matches only a single path segment till the next forward slash.
• GET request made to URL http://server-address/hello/1/message will match the pattern`GET /hello/{user_id}/message`.
• The `{user_id}` wildcard will match 1 in the URL. You can use the PathValue method on the request to get this value.
• If the wildcard ends with ellipses i.e. `...` then it matches the rest of the URL.
• A `...` wildcard is only allowed at the end of the URL.
• Example: `GET /hello/{user_id...}` in this case PathValue `user_id` will be the rest of the URL after hello/.
• The pattern `/{$}` matches only the path "/", whereas the pattern "/" matches every path.

                  Path Conflict


• If two or more patterns match a request, then the most specific pattern takes precedence.
• As an example of the general rule, `/hello/there/` is more specific than `/hello/`, so both can be registered.
• The former matches paths beginning with `/hello/there/` and the latter will match any other path in the `/hello/` subtree.

Package net/http: [https://pkg.go.dev/net/http]
*/

package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Step 1: Create our OWN router (ServeMux) instead of using the default one.
	// Good practice: the default mux is shared globally, so other code could add routes to it.
	router := http.NewServeMux()

	// Step 2: Register a route with a METHOD + PATH pattern: "GET /hello".
	// It matches ONLY GET requests to exactly /hello.
	router.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	// Step 3: Register a route with a wildcard: "POST /hello/{user_id}".
	// {user_id} matches ONE path segment (anything between two slashes).
	// It matches POST requests like /hello/1 or /hello/anjesh.
	router.HandleFunc("POST /hello/{user_id}", func(w http.ResponseWriter, r *http.Request) {

		// Step 4: PathValue reads the wildcard value from the URL.
		// The name inside {} (user_id) is the name we pass here.
		userID := r.PathValue("user_id")

		// Sprintf builds the text "hello <userID>" and returns it (does not print)
		w.Write([]byte(fmt.Sprintf("hello %s", userID)))
	})

	// Step 5: Start the server and give it OUR router (not nil).
	// The router decides which function handles each request.
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
	}

}

/*  How to test (run the program first: go run .):

1. GET  /hello
   curl http://localhost:8080/hello
   Output: hello

2. POST /hello/{user_id}
   curl -X POST http://localhost:8080/hello/anjesh
   Output: hello anjesh

3. Wrong method (GET on a POST-only route)
   curl http://localhost:8080/hello/anjesh
   Output: 405 Method Not Allowed

4. Path that is not registered
   curl http://localhost:8080/abc
   Output: 404 page not found

Note: opening a URL in the browser sends a GET request, so the POST route
cannot be tested from the browser address bar. Use curl or a tool like Postman.

*/