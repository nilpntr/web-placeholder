FROM sammobach/go:1.17 as build
LABEL maintainer="Sam Mobach <hello@sammobach.com>"
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]

FROM alpine:3.14.0
LABEL maintainer="Sam Mobach <hello@sammobach.com>"
RUN mkdir /app
COPY --from=build /app /app
WORKDIR /app
CMD ["/app/main"]