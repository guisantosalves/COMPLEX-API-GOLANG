#inherit all that functions from 1.16-alpine
FROM golang:1.16-alpine
# it creates a directory inside the image that we are building.
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

#execute it inside the image
RUN go mod download

# copying our source code
COPY *.go ./

#compile
RUN go build -o /docker-gs-ping

EXPOSE 3005

