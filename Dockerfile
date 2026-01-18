FROM golang:1.25-alpine AS build

WORKDIR /app
COPY . .

RUN apk add --no-cache make

RUN make build-server &&\
    make build-migrator

FROM alpine:latest AS run

RUN apk add --no-cache make

WORKDIR /app
COPY --from=build /app/lilpapa-server ./lilpapa-server
COPY --from=build /app/lilpapa-migrator ./lilpapa-migrator
COPY --from=build /app/Makefile ./Makefile

EXPOSE 3000

CMD ["make", "lilpapa-server"]
