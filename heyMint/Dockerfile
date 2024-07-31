FROM node:20-alpine as build

# Set the working directory
WORKDIR /usr/src/app/

# Copy package.json and install dependencies
COPY package.json .
COPY tsconfig.json .

RUN npm install

# Copy the rest of the application code
COPY . .

# Expose the application port
EXPOSE 6000

# Start the application
CMD ["sh", "entrypoint.sh"]
