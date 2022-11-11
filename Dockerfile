#build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .


RUN go mod download

RUN go build -o main.app 

#final stage
ARG COMPANY_LOGO=images/company-logo.png
ARG INVOICE_HTML=html/invoice_email.html

FROM alpine:latest
COPY --from=builder app/main.app /app
COPY ${COMPANY_LOGO} /app
COPY ${INVOICE_HTML} /app
ENTRYPOINT /app
LABEL Name=coaching-clinic Version=1.0
EXPOSE 8080
