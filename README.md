## MongoDB + Beego Development Environment


#### Create a new beego project using the beemongo image
`docker run --rm -v "$(pwd)":/go/src/ -w /go/src beemongo bee new beego-docker`


> If you are using docker-machine, look for the machine IP address with the command `docker-machine ls`
## How to run the project
In the console run `docker-compose up` and it will do the heavylifting to start your app

## Where is the source code?
angular4/code, contains the source code for the frontend app
beego/code, contains the source code for the beego API app

## How to access the services

http://localhost:8080/ beego

http://localhost:8000/ mongo-express

http://localhost:4200/ angular4
