# Build stage
FROM golang:alpine AS builder
WORKDIR /
ADD . .
ARG VIPS_VERSION="8.6.5"

# RUN wget https://github.com/libvips/libvips/releases/download/v${VIPS_VERSION}/vips-${VIPS_VERSION}.tar.gz
# RUN apk update && apk add automake build-base pkgconfig glib-dev gobject-introspection libxml2-dev expat-dev jpeg-dev libwebp-dev libpng-dev
# # Exit 0 added because warnings tend to exit the build at a non-zero status
# RUN tar -xf vips-${VIPS_VERSION}.tar.gz && cd vips-${VIPS_VERSION} && ./configure && make && make install && ldconfig; exit 0
# RUN apk del automake build-base
RUN go build -o images cmd/main.go


# Deployment package
FROM alpine:latest
WORKDIR /
COPY --from=builder images .
RUN mkdir -p public/assets/images/actual
RUN mkdir -p public/assets/images/thumbnail
RUN mkdir -p public/assets/images/small
RUN chmod +x images

EXPOSE 8091
ENTRYPOINT [ "./images" ]