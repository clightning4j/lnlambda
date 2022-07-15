FROM golang:alpine3.16
LABEL maintainer="Vincenzo Palazzo (@vincenzopalazzo) vincenzopalazzodev@gmail.com"

RUN make dep

CMD [ "make", "run" ]