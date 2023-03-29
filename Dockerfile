# Build image
FROM golang:1.19-buster as go-server

WORKDIR /app/server

# Copy go modules file
COPY ./server/go.mod .
COPY ./server/go.sum .

# Install dependencies
RUN go mod download

# Copy app source code
COPY ./server .

# Build app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .

# Base image
FROM node:14-alpine AS base

WORKDIR /app

# Copy over compiled code from go-server image

# Set working directory
WORKDIR /app

# Copy app source code
COPY ./app .

# Install dependencies
RUN yarn install --frozen-lockfile

# Build app
RUN yarn build

# Production image
FROM alpine:latest

WORKDIR /app
COPY ./server/app.env .
COPY ./app/.env .
# Install dependencies
RUN apk add --no-cache ca-certificates

# Copy over compiled code from base image
COPY --from=base /app/dist ./dist
COPY --from=go-server /main ./
RUN chmod +x ./main
EXPOSE 8080
CMD ./main

# Start app
# CMD ["./server"]
