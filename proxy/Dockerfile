#syntax=docker/dockerfile:1

# This Dockerfile uses the service folder as context.


# --
# Upstream images

FROM golang:1.23-alpine AS go_upstream
FROM alpine:3.20 AS alpine_upstream


# --
# Base image

FROM go_upstream AS app_base

# Set app directory
WORKDIR /app


# --
# Runtime base image

FROM app_base AS app_runtime_base

# Set runtime environment
ENV APP_ENV=dev

# Copy application source code
COPY --link ./app .


# --
# Dev image

FROM app_runtime_base AS app_dev

# Set exposed port
ARG PORT=8080
ENV PORT=${PORT}
EXPOSE ${PORT}

CMD [ "go", "run", "main.go" ]


# --
# Prod build image

FROM app_runtime_base AS app_prod_build

# Set runtime environment
ENV APP_ENV=prod

# Build application
RUN go build -o main .


# --
# Prod image

FROM alpine_upstream AS app_prod

# Set app directory
WORKDIR /app

# Set runtime environment
ENV APP_ENV=prod

# Copy application binary
COPY --from=app_prod_build --chmod=755 /app/main ./main

# Set exposed port
ARG PORT=8080
ENV PORT=${PORT}
EXPOSE ${PORT}

CMD [ "./main" ]
