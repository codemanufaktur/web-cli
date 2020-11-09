FROM alpine:3.10
COPY web-cli /web-cli
ENTRYPOINT ["./web-cli serve"]