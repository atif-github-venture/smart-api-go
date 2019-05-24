FROM golang:alpine AS build-env
RUN apk add git
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go get gopkg.in/yaml.v2
ADD . /go/src/smartapigo/create-project
WORKDIR /go/src/smartapigo/create-project
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o create-project .
RUN apk --no-cache add curl
EXPOSE 8080
ENTRYPOINT ["./create-project"]



# cd to this directory and then
    # docker build -t create-project .
# Run the image
    # docker run -p 8080:8080 -it create-project
# To check the container id
    # docker ps -> this will give list of containers running (get container id)
    #docker logs -f c78e59810694
#**** but mongo db should be up and you need to change the mongo db config on each service to point to your local and not the container


#  docker login
#  docker tag create-project atifdockerventure/create-project
#  docker push atifdockerventure/create-project:latest