FROM golang:latest
COPY . /src
WORKDIR /src
RUN go mod download
CMD ["go", "run", "main.go"]
