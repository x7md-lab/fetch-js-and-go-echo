# Simple example of Fetch.

## Server-Side
GoLang Server.

```sh
cd ./server-side
# install packages
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
# run server
go run ./main.js
```

## Broswer Side

WebPage (HTML + JavaScript)

that call Server through HTTP API.

```sh
# Serve Page To Browser, You can use anythig, ex:
cd ./broswer-side
python3 -m http.server 
```
