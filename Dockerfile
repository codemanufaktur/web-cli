FROM gcr.io/distroless/static:latest
COPY web-cli .
EXPOSE 8080/tcp
ENTRYPOINT ["/web-cli", "serve"]