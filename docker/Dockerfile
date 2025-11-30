FROM golang:1.20-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /ecommerce ./cmd/api


FROM scratch
COPY --from=build /ecommerce /ecommerce
EXPOSE 8080
ENTRYPOINT ["/ecommerce"]