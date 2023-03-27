# Base image
FROM node:14-alpine AS base

# Set working directory
WORKDIR /app

# Copy over package files
COPY package.json yarn.lock ./

# Install dependencies
RUN yarn install --frozen-lockfile


# Copy over source code
COPY . .

# Build app
RUN yarn build

# Production image
FROM node:14-alpine AS production

# Set working directory
WORKDIR /app

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
