#build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .


RUN go mod download

RUN go build -o main.app 

#final stage
FROM alpine:latest
COPY --from=builder app/main.app /app
COPY --from=builder app/images/company-logo.png /app
COPY --from=builder app/html/invoice_email.html /app
COPY --from=builder app/.env /app
ENTRYPOINT /app
LABEL Name=coaching-clinic Version=1.0
EXPOSE 8080
