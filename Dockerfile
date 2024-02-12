# build backend
FROM golang:1.22.0 as backend-builder

COPY server /server
WORKDIR /server
RUN go build -o app

# build frontend

# run application
FROM scratch as app
COPY --from=backend-builder /server/app /server

CMD ["./server"]
