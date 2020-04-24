package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("POST", "/api", server.AddMiddleware(HandleApi, CheckAuth(), Logging()))
	server.Listen()
}
