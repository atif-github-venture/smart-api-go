FROM golang:alpine AS build-env
RUN apk add git
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
ADD . /go/src/smartapigo/creare-project
WORKDIR /go/src/smartapigo/creare-project
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o creare-project .
#RUN go build main.go
#COPY . /go/src/smartapigo/creare-project

#CMD ["go", "run", "./creare-project/main.go"]
#EXPOSE 8080
RUN apk --no-cache add curl
EXPOSE 8080
#COPY --from=builder ${SOURCE_LOCATION} .
ENTRYPOINT ["./creare-project"]



# cd to this directory and then
    # docker build -t creare-project .
# Run the image
    # docker run -p 8080:8080 -it creare-project
# To check the container id
    # docker ps -> this will give list of containers running (get container id)
    #docker logs -f c78e59810694
#**** but mongo db should be up and you need to change the mongo db config on each service to point to your local and not the container


#  docker login
#  docker tag creare-project atifdockerventure/creare-project
#  docker push atifdockerventure/creare-project:latest