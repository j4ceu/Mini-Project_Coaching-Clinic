#build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .


RUN go mod download

RUN go build -o main.app 

#final stage

FROM alpine:latest

FROM alpine:latest
COPY --from=builder app/main.app /app/main.app
COPY --from=builder app/images/company-logo.png /images/company-logo.png
COPY --from=builder app/html/invoice_email.html /html/invoice_email.html



ENTRYPOINT ["/app/main.app"]
LABEL Name=coaching-clinic Version=1.0
EXPOSE 8080
