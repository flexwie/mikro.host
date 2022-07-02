FROM registry.semaphoreci.com/golang:1.18 as builder

ENV APP_HOME /go/src/user
WORKDIR "$APP_HOME"
COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o user

FROM registry.semaphoreci.com/golang:1.18

ENV APP_HOME /go/src/user
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY --from=builder "$APP_HOME/user" $APP_HOME

EXPOSE 8080
CMD ["./user"]