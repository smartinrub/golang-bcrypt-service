FROM golang:latest 
WORKDIR /go/src/bcrypt-service
# -d not to install packages; -v verbose
RUN go get -d -v golang.org/x/crypto/bcrypt
RUN go get -d -v github.com/gin-gonic/gin
# copy go file to current directory (/go/src/main)
COPY main.go .
# disable cgo, target operating system, -a force rebuilding packages
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
ENV ROUNDS=12
WORKDIR /root/
# copies the first build into this stage
COPY --from=0 /go/src/bcrypt-service/main .
CMD ["./main"]
