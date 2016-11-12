# WTA-58 Go Network Tutorial

Sample GoLang code to satisfy tutorial exercise.

## Installation
Clone files from github.com/bryantvalverde/gotest into your local machine $GOPATH

## Running
Requires 2 terminals to run program

### Terminal 1 (Server)
Navigate to $GOPATH/github.com/bryantvalverde/gotest/server and run receiver.
```
go run server.go
```


### Terminal 2 (Client)
Navigate to $GOPATH/github.com/bryantvalverde/gotest/client and run receiver.
```
go run client.go
```
Receiver.go will close once the packet has been sent. You will see the output on terminal 1.