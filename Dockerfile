FROM golang:latest as builder
WORKDIR /app
ADD . /app/
# Compile the application to a static binary
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o go-jokly .
RUN #go build -o go-jokly

#FROM scratch
#FROM alpine:latest
#WORKDIR /app
#COPY --from=builder /app/go-jokly .
#RUN ls -l /app
CMD ["./go-jokly"]