FROM golang:1.18
WORKDIR /src/backend

ADD go.mod go.sum /src/backend/
RUN go mod download

ADD *.go /src/backend/

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o /bin/backend

FROM scratch
WORKDIR /opt/backend/
ARG VERSION=dev

COPY --from=0 /bin/backend /opt/backend/backend

ENV APP_DICE_MAX="6"
ENV APP_BTN_ROLL_TITLE="Roll"

EXPOSE 8080

ENTRYPOINT ["/opt/backend/backend"]
