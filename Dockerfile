FROM golang:alpine as gobuild
WORKDIR /app

ADD ./ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -o ./build/agent -a ./tcpmirror.go

FROM scratch
WORKDIR /app
EXPOSE 8080
COPY --from=gobuild /app/build/agent ./

CMD [ "/app/agent" ]