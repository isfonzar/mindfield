# build backend
FROM golang:1.22.0 as backend-builder

COPY server /server
WORKDIR /server

# Disable CGO and statically compile the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# build frontend
COPY resources /resources

# run application
FROM scratch as app
COPY --from=backend-builder /server/server /server/
COPY --from=backend-builder /resources/html /server/

WORKDIR /server
CMD ["./server"]
