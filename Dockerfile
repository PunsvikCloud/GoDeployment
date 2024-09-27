FROM golang:1.23-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o GoDeployment .

FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/GoDeployment .
EXPOSE 8080
CMD ["./GoDeployment"]
