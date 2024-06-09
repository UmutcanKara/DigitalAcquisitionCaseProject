# Use an official Node.js runtime as a parent image
FROM node:latest

# Set the working directory in the container
WORKDIR /app

# Copy package.json and yarn.lock into the working directory
COPY package.json yarn.lock ./

# Install project dependencies using Yarn
RUN yarn install

# Copy the rest of the application code into the working directory
COPY . .

# Expose the port that the app runs on
EXPOSE 5173

# Start the React application in development mode
CMD ["yarn", "dev"]