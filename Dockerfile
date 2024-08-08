FROM node:alpine3.20

# Installs latest Chromium (77) package.
RUN apk add --no-cache \
      chromium \
      nss \
      freetype \
      freetype-dev \
      harfbuzz \
      ca-certificates \
      ttf-freefont \
      nodejs \
      npm \
      tor

# Create app directory
WORKDIR /app

# Copy app artifacts and dependencies
COPY . .
RUN npm install

EXPOSE 3000
CMD ["node", "/app/src/main.js"]