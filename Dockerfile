FROM golang
LABEL maintainer="zhenghaoayng<zhenghaoayng@topsec.com.cn>"
COPY main.go /app/
WORKDIR /app/
EXPOSE 8888
ENTRYPOINT [ "/bin/bash","-c","go run main.go"]
