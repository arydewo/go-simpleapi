# go-simpleapi

'simpleapi' is a simple RESTful API which response with dummy response or records retrieved from users table on DB


## Database Preparation

Run the file scripts/db.sql to create table and populate data

## Install Go runtime

Download [go](https://golang.org/dl/).

```bash
 $ wget https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
```

Install (default)
```bash
 $ sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```

Set PATH and GOPATH
```bash
 $ export PATH=$PATH:/usr/local/go/bin
 $ export GOPATH=/home/myuser/mygopath
```
Replace GOPATH with the working directory you're using. 
Your application codes can be copied to folder under the GOPATH

## Install and Build Your Application

Install Dependencies
```bash
 $ go get github.com/go-sql-driver/mysql
 $ go get github.com/gorilla/mux
```

Build the application. Go to your source directory containing main.go

```bash
 $ go build
```

## Usage
Set dbtype and dburl into your Environment Variables. Here's the sample
```bash
 $  export dbtype='mysql'
 $  export dburl='mydbuser:mydbpassword@tcp(127.0.0.1:3306)/mydbschema'
```

Run the application
```bash
 $ ./simpleapi
```

Open new terminal and go to your application directory. Run test using curl
```bash
 $ curl http://localhost:8888/users
```


## Contributing
Fell free to pull these codes.

## License
[MIT](https://choosealicense.com/licenses/mit/)
