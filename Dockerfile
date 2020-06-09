FROM golang:alpine
USER root
RUN mkdir -p $GOPATH/src/rating-service
COPY . $GOPATH/src/rating-service
WORKDIR $GOPATH/src/rating-service/
ENTRYPOINT go run project.go
EXPOSE 6666