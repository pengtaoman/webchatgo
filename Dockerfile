FROM alpine:3.15.0
WORKDIR /
RUN mkdir -p /html/template
COPY main /
COPY html/* /html
COPY html/template/* /html/template
ENTRYPOINT ["/main"]