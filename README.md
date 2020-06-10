# EasyBuggy Go
EasyBuggy clone build on Gin. [EasyBuggy](https://github.com/k-tamura/easybuggy) is a broken web application in order to understand behavior of bugs and vulnerabilities, for example, memory leak, deadlock, JVM crash, SQL injection and so on.

Note: Under Constructing.

Quick Start
-
First you need to install Go and Gin.
```bash
$ dnf -y install go # If you use CentOS 8 and have not installed Go yet.
$ go get -u github.com/gin-gonic/gin
```
Start up EasyBuggy Go.
```bash
$ go run main.go
```

Access to:

    http://localhost:8080