# bench_go

# installation
```bash
wget https://golang.org/dl/go1.16.5.linux-amd64.tar.gz
tar xvf go1.16.5.linux-amd64.tar.gz 
sudo chown -R root:root ./go
sudo mv go /usr/local
sudo nano ~/.profile
## Add the following at the end: ##
export GOPATH=$HOME/work
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

source ~/.profile
mkdir $HOME/work

go env -w GO111MODULE=auto

go get -u github.com/gin-gonic/gin
go get -u github.com/go-pg/pg
go get github.com/go-pg/pg/v9
go get github.com/go-pg/pg/v9/orm

go run main.go
```

# Help Docs
- Go Gin Core:
- Go Gin Framework: https://github.com/gin-gonic/gin
- HTML Rendering in Gin: https://github.com/gin-gonic/gin#html-rendering
- Go Goroutines: https://github.com/gin-gonic/gin#goroutines-inside-a-middleware
- Run Multiple Services: https://github.com/gin-gonic/gin#run-multiple-service-using-gin

DB:
- PostgreSQL: https://pg.uptrace.dev/
- ArangoDB: https://www.arangodb.com/docs/stable/drivers/go-getting-started.html
- MongoDB: https://docs.mongodb.com/drivers/go/
