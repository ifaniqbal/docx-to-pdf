# Use a Node.js base image
FROM node:18

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the code to the working directory
COPY . .

# Expose the port on which the application will run (optional)
EXPOSE 80

# Define the command to start the application
CMD [ "node", "app.js" ]
