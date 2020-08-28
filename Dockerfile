FROM golang:1.9.2 
ADD . /go/src/Genity-Demo
WORKDIR /go/src/Genity-Demo
RUN go get Genity-Demo
RUN go install
EXPOSE 5000
ENTRYPOINT ["/go/bin/Genity-Demo"]