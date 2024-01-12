# FROM golang:1.21 AS build-stage

# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . ./

# RUN CGO_ENABLED=0 GOOS=linux go build -o /api *.go

# # Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...


FROM quay.io/podman/stable

COPY deployment/registries.conf /etc/containers/registries.conf
COPY dockerfiles /dockerfiles

# COPY --from=build-stage /api /api

# ENTRYPOINT ["/api"]
ENTRYPOINT ["sleep", "infinity"]