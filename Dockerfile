FROM alpine:latest
COPY build/axiom /bin/app
WORKDIR /bin
ENTRYPOINT ["/bin/app"]
EXPOSE 8080
