FROM golang:1.17-alpine AS build
WORKDIR /src/
COPY . /src/
RUN mkdir -p /build
RUN CGO_ENABLED=0 go build -o /build/bardview5

from scratch
COPY --from=build /build/bardview5 /bin/bardview5
EXPOSE 8080
ENTRYPOINT ["/bin/bardview5", "serve"]
