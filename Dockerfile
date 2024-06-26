FROM golang:alpine3.19
LABEL maintainer="Vincenzo Palazzo (@vincenzopalazzo) vincenzopalazzodev@gmail.com"

workdir lnlambda

COPY . .

RUN go get -d ./...

CMD [ "go", "run", "main.go" ]
