FROM golang:1.17-buster  

WORKDIR /go

ENV POSTGRES_URL=postgres://bhdhkulz:2e6RiBcN2HpdIy9bGWvDSnkVH3NAUWw6@tyke.db.elephantsql.com/bhdhkulz
  
COPY . .
 
RUN go mod tidy

RUN go build -v -o main

CMD ["/go/main" "monolith"]
 