FROM golang:1.18-bullseye

ENV GO111MODULE=on

ENV APP_HOME /go/src/

RUN mkdir -p "$APP_HOME"

ENV POSTGRES_URL=postgres://bhdhkulz:2e6RiBcN2HpdIy9bGWvDSnkVH3NAUWw6@tyke.db.elephantsql.com/bhdhkulz

WORKDIR "$APP_HOME"

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build

EXPOSE 5000

ENTRYPOINT ["m"]
