# Use a more secure Node.js base image
FROM node:23-slim

# Create and set the working directory
WORKDIR /backend

# Copy package.json and install dependencies
COPY package*.json ./
RUN npm install

# Copy the entire project
COPY . .

# Build the project
RUN npm run build

# Expose the NestJS app on port 3000
EXPOSE 3000

# Run the app
CMD ["npm", "run", "start:prod"]