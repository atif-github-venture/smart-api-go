FROM golang:alpine AS build-env
RUN apk add git
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
ADD . /go/src/smartapigo/identifier-reponere
WORKDIR /go/src/smartapigo/identifier-reponere
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o identifier-reponere .
#RUN go build main.go
#COPY . /go/src/smartapigo/identifier-reponere

#CMD ["go", "run", "./identifier-reponere/main.go"]
#EXPOSE 8080
RUN apk --no-cache add curl
EXPOSE 8080
#COPY --from=builder ${SOURCE_LOCATION} .
ENTRYPOINT ["./identifier-reponere"]
#


# cd to this directory and then
    # docker build -t identifier-reponere .
# Run the image
    # docker run -p 8080:8080 -it identifier-reponere
# To check the container id
    # docker ps -> this will give list of containers running (get container id)
    #docker logs -f c78e59810694