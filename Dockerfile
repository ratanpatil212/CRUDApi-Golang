# Base Image
FROM golang:latest

#Specifying working directory
WORKDIR /usr/src/app

COPY . .

# Specifying container's internal port
EXPOSE 4000

# Entry point of the application
CMD ["go","run","main.go"]