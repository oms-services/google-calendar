FROM golang

RUN go get github.com/gorilla/mux

RUN go get -v golang.org/x/oauth2/google

RUN go get -v google.golang.org/api/calendar/v3

WORKDIR /go/src/github.com/oms-services/google-calendar

ADD . /go/src/github.com/oms-services/google-calendar

RUN go install github.com/oms-services/google-calendar

ENTRYPOINT google-calendar

EXPOSE 3000