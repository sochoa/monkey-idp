FROM golang:latest
COPY . /src
WORKDIR /src
CMD ["go", "run", "main.go"]
