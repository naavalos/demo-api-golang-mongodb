# Demo API REST - Golang - MongoDB

Template para desarrollar una **API REST** con **Golang** + **MySQL**.\
Utilizamos la entidad **Persona** para realizar operaciones **CRUD** (Create, Read, Update, Delete) sobre una Base de Datos NoSQL.

## API REST

- ### GET --> [http://localhost:8080/people](http://localhost:8080/people)
Devuelve **todas** las personas de la base de datos

>###### Response
>```sh
>200-OK
>```
>```sh
>[
>   {
>      "_id":"1",
>      "firstname":"Juan",
>      "lastname":"Perez",
>      "adress":{
>         "street":"Rivadavia",
>         "number":"3355"
>      }
>   },
>   {
>      "_id":"2",
>      "firstname":"Ana",
>      "lastname":"Lia",
>      "adress":{
>         "street":"Nazca",
>         "number":"1824"
>      }
>   }
>]
>```

- ### GET --> [http://localhost:8080/people/{_id}](http://localhost:8080/people/{_id})
Devuelve la persona con **_id**=*1*

>###### Response
>```sh
>200-OK
>```
>```sh
>{
>   "_id":"1",
>   "firstname":"Juan",
>   "lastname":"Perez",
>   "adress":{
>      "street":"Rivadavia",
>      "number":"3355"
>   }
>}
>```

- ### POST --> [http://localhost:8080/people](http://localhost:8080/people)
Crea una persona

>###### Request
>```sh
>{
>   "firstname":"Nahuel",
>   "lastname":"Avalos",
>   "adress":{
>      "street":"Avenida Siempreviva",
>      "number":"742"
>   }
>}
>```

>###### Response
>```sh
>200-Ok
>```

## Tecnología

- [Golang](https://golang.org/dl/)
- [Visual Studio Code](https://code.visualstudio.com/download)
- [MongoDB](https://www.mongodb.com/es)
- [Git](http://https://git-scm.com/)
- [Postman](https://www.postman.com/downloads/)

## Environment

Comandos para configurar el ambiente en Linux

```sh
## Golang
$ sudo tar -C /usr/local -xzf go1.15.6.linux-amd64.tar.gz
$ vi /home/USER_NAME/.profile
  export PATH=$PATH:/usr/local/go/bin
$ source /home/USER_NAME/.profile
$ go version

## Plugins
$ go get -v -u github.com/gorilla/mux
$ go get -v -u go.mongodb.org/mongo-driver/mongo
$ go get -v -u go.mongodb.org/mongo-driver/mongo/options
$ go get -v -u go.mongodb.org/mongo-driver/bson
$ go get -v -u go.mongodb.org/mongo-driver/bson/primitive

## MondoDB
$ sudo apt install -y mongodb
$ sudo systemctl status mongodb
$ mongo --eval 'db.runCommand({ connectionStatus: 1 })'

## API
$ git clone https://github.com/nahuelavalos/demo-api-golang-mongodb.git
$ cd demo-api-golang-mongodb/
$ go build main.go
$ ./main
```

## Run

Comandos para correr la API REST. 

Por defecto queda levantada en [http://localhost/8080/](http://localhost/8080/)

```sh
$ nohup go main &
```


## Base de Datos

**MongoDB** (NoSQL) Local.

Por defecto queda levantada en **mongodb://localhost:27017**
