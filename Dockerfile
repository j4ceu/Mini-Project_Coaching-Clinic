#build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .

ENV DB_URI = postgres://wkqyugttfjqguu:929050c0ebe575fd9631720fc66873b17f991731f824e79752569313951aff1b@ec2-44-195-132-31.compute-1.amazonaws.com:5432/dejj3hp60kc8pc
ENV SECRET_JWT = jace
ENV EMAIL =	j4ceucoachingclinic@gmail.com
ENV EMAIL_PASSWORD = fssqilhrgqwrdzhl

RUN go mod download

RUN go build -o main.app 

#final stage
FROM alpine:latest
COPY --from=builder app/main.app /app
ENTRYPOINT /app
LABEL Name=coaching-clinic Version=1.0
EXPOSE 8080
