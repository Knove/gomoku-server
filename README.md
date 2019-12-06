# gomoku-server
 
### 部署指令

```bash
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build

docker build -t gomoku-server .

docker run -d -p 7777:7777 gomoku-server
```

