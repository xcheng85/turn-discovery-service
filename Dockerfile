FROM gcr.io/k8s-skaffold/skaffold:v1.39.0 as skaffold
FROM golang:1.17-buster as builder

RUN apt-get update -yq && apt-get upgrade -yq

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN go build -v -o server

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates bash curl && \
    rm -rf /var/lib/apt/lists/*
RUN curl http://stedolan.github.io/jq/download/linux64/jq -o /bin/jq && chmod a+x /bin/jq 

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /app/server

# Install kubectl
COPY --from=skaffold /usr/local/bin/kubectl /usr/local/bin/kubectl

COPY entryPoint.sh /
RUN chmod +x /entryPoint.sh
ENTRYPOINT ["/entryPoint.sh"]
