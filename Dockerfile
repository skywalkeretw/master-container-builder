FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /api

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...


FROM quay.io/podman/stable

COPY --from=build-stage /api /api

ENTRYPOINT ["/api"]