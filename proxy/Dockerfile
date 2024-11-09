#syntax=docker/dockerfile:1

# This Dockerfile uses the service folder as context.


# --
# Upstream images

FROM golang:1.23-alpine AS go_upstream


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
# Prod image

FROM app_runtime_base AS app_prod

# Set runtime environment
ENV APP_ENV=prod

# Build application
RUN go build -o main .

# Set exposed port
ARG PORT=8080
ENV PORT=${PORT}
EXPOSE ${PORT}

CMD [ "./main" ]
