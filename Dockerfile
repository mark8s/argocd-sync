FROM golang:1.16.13
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build . 
CMD ["/app/argocd-sync"]
