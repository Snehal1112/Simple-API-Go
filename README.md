# Docs

[Task Management API](#) is a simple API for task management.
Developed under Go (Golang) platform.
This project developed using [gin](https://github.com/gin-gonic/gin) framework to ease the development.
See [Source](#source) For all the library used

## Index

* [Overview](#overview)
* [Support](#support)
* [Quick Start](#quick-start)
* [Source](#source)


## Overview

1. Import the SQL to your MySql Database. 

2. Run the project

3. Open  <http://localhost:9090/ninja/docs/> for API documentation 

4. Integrate with your apps Android , Angular, React
    
See [Quick Start](#quick-start) for example implementations.


## Support


You can also email <iman.tumorang@gmail.com> or file an [Issue](https://github.com/bxcodec/Simple-API-Go/issues/new).


## Quick Start

### Command Line

Run your project using this command

```bash
# download the project
git clone https://github.com/bxcodec/Simple-API-Go.git
cd simple-api-go

# run the project
go run server.go

```

## Source

Below listed all the library used here.

### [gin](https://github.com/gin-gonic/gin) 
Gin is a web framework written in Go (Golang). It features a martini-like API with much better performance, up to 40 times faster thanks to [httprouter](https://github.com/julienschmidt/httprouter). If you need performance and good productivity, you will love Gin.

### [sqlboiler](https://github.com/vattle/sqlboiler)
SQLBoiler is a tool to generate a Go ORM tailored to your database schema.

It is a "database-first" ORM as opposed to "code-first" (like gorm/gorp). That means you must first create your database schema. Please use something like [goose](https://bitbucket.org/liamstask/goose), [sql-migrate](https://github.com/rubenv/sql-migrate) or some other migration tool to manage this part of the database's life-cycle.

### [jwt-go](https://github.com/dgrijalva/jwt-go)
A [go](http://www.golang.org/) (or 'golang' for search engine friendliness) implementation of JSON Web Tokens
