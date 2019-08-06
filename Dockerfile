FROM golang

RUN go get github.com/gorilla/mux

RUN go get -v golang.org/x/oauth2/google

RUN go get -v google.golang.org/api/calendar/v3

WORKDIR /go/src/github.com/heaptracetechnology/google-calendar

ADD . /go/src/github.com/heaptracetechnology/google-calendar

RUN go install github.com/heaptracetechnology/google-calendar

ENTRYPOINT google-calendar

EXPOSE 3000