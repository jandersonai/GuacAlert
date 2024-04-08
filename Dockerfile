FROM alpine:latest
WORKDIR /app
RUN apk update
COPY /bin/GuacAlert /app/GuacAlert
RUN chmod +x /app/GuacAlert
ENTRYPOINT ["./GuacAlert"]