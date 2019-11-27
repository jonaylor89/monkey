######## Start build phase of execution #######
FROM golang:latest as builder

LABEL maintainer="John Naylor <jonaylor89@gmail.com>"
 
WORKDIR /app

COPY . .
 
RUN go build -o monkey

######## Start a new stage from scratch #######
FROM alpine:latest 

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/monkey .

CMD ["./monkey"]
