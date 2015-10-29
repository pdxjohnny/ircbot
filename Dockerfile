FROM busybox
ADD ./ircbot_linux-amd64 /app
CMD ["/app"]
