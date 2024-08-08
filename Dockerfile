FROM node:alpine3.20

# Install necessary packages for Puppeteer
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
      tor \
      udev \
      ttf-opensans

# Create app directory
WORKDIR /app

# Copy app artifacts and dependencies
COPY . .
RUN npm install

# Set Puppeteer environment variable to use installed Chromium
ENV PUPPETEER_EXECUTABLE_PATH=/usr/bin/chromium-browser

EXPOSE 3000
CMD ["npm", "start"]
