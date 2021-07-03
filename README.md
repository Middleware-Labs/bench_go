# bench_go

# installation
```bash
wget https://golang.org/dl/go1.16.5.linux-amd64.tar.gz
tar xvf go1.16.5.linux-amd64.tar.gz 
sudo chown -R root:root ./go
sudo mv go /usr/local
sudo nano ~/.profile
source ~/.profile
mkdir $HOME/work

go env -w GO111MODULE=auto

go get -u github.com/gin-gonic/gin
go get -u github.com/go-pg/pg
go get github.com/go-pg/pg/v9
go get github.com/go-pg/pg/v9/orm
```
