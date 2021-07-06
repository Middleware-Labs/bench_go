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
## Go Gin Core:
- Go Gin Framework: https://github.com/gin-gonic/gin
- HTML Rendering in Gin: https://github.com/gin-gonic/gin#html-rendering
- Go Goroutines in Gin: https://github.com/gin-gonic/gin#goroutines-inside-a-middleware
- Run Multiple Services: https://github.com/gin-gonic/gin#run-multiple-service-using-gin

## Benefits of Goroutines
- https://www.geeksforgeeks.org/goroutines-concurrency-in-golang/
- Goroutines are cheaper than threads.
- Goroutine are stored in the stack and the size of the stack can grow and shrink according to the requirement of the program. But in threads, the size of the stack is fixed.
- Goroutines can communicate using the channel and these channels are specially designed to prevent race conditions when accessing shared memory using Goroutines.
- Suppose a program has one thread, and that thread has many Goroutines associated with it. If any of Goroutine blocks the thread due to resource requirement then all the remaining Goroutines will assign to a newly created OS thread. All these details are hidden from the programmers.

## DB:
- PostgreSQL (ORM Based): https://pg.uptrace.dev/
- Pure PostgreSQL Driver: https://pkg.go.dev/github.com/lib/pq

- ArangoDB: https://www.arangodb.com/docs/stable/drivers/go-getting-started.html
- - ArangoDB Example Usage: https://www.arangodb.com/docs/stable/drivers/go-example-requests.html

- MongoDB: https://docs.mongodb.com/drivers/go/
- - MongoDB Example Usage: https://github.com/mongodb/mongo-go-driver#usage
