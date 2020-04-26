package main

func main() {
	server := NewServer(":8080")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("POST", "/api", server.AddMiddleware(HandleApi, CheckAuth(), Logging()))
	server.Listen()
}
