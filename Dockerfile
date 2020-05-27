FROM golang
LABEL maintainer="zhenghaoayng<zhenghaoayng@topsec.com.cn>"
RUN go get -d -v github.com/gin-gonic/gin
COPY main.go /go/src/app/
WORKDIR /go/src/app/
EXPOSE 8080
ENTRYPOINT [ "/bin/bash","-c","go run main.go"]
