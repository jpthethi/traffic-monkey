FROM golang:latest
RUN mkdir -p /go/src/traffic-monkey
WORKDIR /go/src/traffic-monkey
ADD . /go/src/traffic-monkey
RUN go get
RUN go build
# RUN chmod 755 ./traffic-monkey
CMD ["./traffic-monkey"]
EXPOSE 6000/udp
EXPOSE 3060
