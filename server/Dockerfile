FROM golang:1.19-buster as go-server
ADD . /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/server .

# Base image
FROM node:14-alpine AS base
COPY --from=go-server /app/app ./
# Set working directory

# Copy over package files
COPY package.json yarn.lock ./

# Install dependencies
RUN yarn install --frozen-lockfile


# Copy over source code
COPY . .

# Build app
RUN yarn build

# Production image
FROM alpine:latest

# Copy over package files
COPY package.json yarn.lock ./

# Install dependencies
RUN yarn install --frozen-lockfile --production=true
RUN yarn global add serve

# Copy over compiled code from base image
COPY --from=base /app/dist ./dist

# Expose port
EXPOSE 3000

# Start app
CMD ["yarn", "start"]
# CMD ["serve", "-s", "/dist", "-l", "3000"]
