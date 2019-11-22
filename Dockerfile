FROM golang:alpine AS build-env
RUN apk add git
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go get gopkg.in/yaml.v2
ADD . /go/src/smartapigo/action-list
WORKDIR /go/src/smartapigo/action-list
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o action-list .
RUN apk --no-cache add curl
EXPOSE 8080
ENTRYPOINT ["./action-list"]



# cd to this directory and then
    # docker build -t action-list .
# Run the image
    # docker run -p 8080:8080 -it action-list
# To check the container id
    # docker ps -> this will give list of containers running (get container id)
    #docker logs -f c78e59810694
#**** but mongo db should be up and you need to change the mongo db config on each service to point to your local and not the container


#  docker login
#  docker tag action-list atifdockerventure/action-list
#  docker push atifdockerventure/action-list:latest