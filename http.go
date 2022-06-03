package main

import (
	"fmt"
	"net/http"
	"strings"
)

func httpPackage() {
	http.HandleFunc("/login", myLogin)
	http.HandleFunc("/welcome/", myWelcome)

	fmt.Println("Listening on Port 5000...")
	http.ListenAndServe("localhost:5000", nil)
}

func myLogin(writer http.ResponseWriter, request *http.Request) {
	if strings.ToLower(request.Method) == "get" {
		fmt.Fprintf(writer, `
			<html>
				<body>
					<h1>Please enter your username and pasword</h1>
				</body>
			</html>
		`)
	}
}

func myWelcome(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, `
		<html>
			<body>
				<h1>Welcome</h1>
			</body>
		</html>
	`)
}
