# Template Initial API

This template allows you to start with a basic configuration of an API in Go, using the ECHO framework.

It is implemented for MYSQL but it can be adapted to the implementation of another database engine, since it uses the DAO design pattern.

### Install the project dependencies
`go get ./...`

### Create the necessary configurations
In the config folder, create two .json files
* database.json
        `
          {
            "engine": "",
            "host": "",
            "port": "",
            "user": "",
            "password": "",
            "database": ""
          }
        `
* server.json
        `
          {
            "protocol": "",
            "host": "",
            "port": "",
          }
        `
**In each of the values ​​insert the information of your project**

### Create the certificates Openssl
In the folder auth
**if you use linux:**
* private certificate
  `openssl genrsa -out app.rsa 1024`
* public certificate
  `openssl rsa -in app.rsa -pubout > app.rsa.pub`
**if you use Windows:**
Search the internet how to generate them xD

