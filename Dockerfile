FROM golang:buster

RUN echo $GOBIN
# Set the Current Working Directory inside the container
WORKDIR $GOBIN/codetuna/busybarbers

RUN export PATH=$PATH:$GOPATH/src/codetuna/busybarbers
# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

ENV POSTGRES_URL=postgres://bhdhkulz:2e6RiBcN2HpdIy9bGWvDSnkVH3NAUWw6@tyke.db.elephantsql.com/bhdhkulz
 
# This container exposes port 8080 to the outside world
EXPOSE 8080
EXPOSE 5000

RUN nohup ./main monolith &
# Run the executable