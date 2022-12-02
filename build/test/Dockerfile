FROM golang:1.19

COPY . /app
WORKDIR /app/tests/integration

ENTRYPOINT ["go", "test", "-v"]
