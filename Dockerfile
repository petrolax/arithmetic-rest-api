FROM golang:1.15.6
RUN mkdir /res
ADD . /res
WORKDIR /res
RUN go build -o main .
CMD [ "./main" ]
