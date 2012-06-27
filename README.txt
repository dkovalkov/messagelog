Message log server

This is simple logging server in Go. It can concurrently write up to 500 messages.

Messages wrote with http request.
Like
http://localhost:8080/log?msg=some_data
Or using HTTP POST method.
