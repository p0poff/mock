# FROM ghcr.io/umputun/baseimage/buildgo:latest as build

# WORKDIR /build
# ADD . /build

# RUN CGO_ENABLED=1 go build -mod=vendor -o app ./app/main.go


FROM ghcr.io/umputun/baseimage/app:latest

# COPY --from=build /build/app /srv/app
# COPY --from=build /build/html /srv/html
RUN mkdir /srv/app
COPY ./html /srv/html
COPY ./app/main /srv/app

RUN chown -R 1001 /srv

EXPOSE 8080
WORKDIR /srv

CMD ["/srv/app/main"]
