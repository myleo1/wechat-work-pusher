FROM alpine:latest
LABEL maintainer="leo <leo@leosgo.com>" \
	version="v1.0.0" \
	description="Wechat-Work-Pusher"
WORKDIR /root
ADD wechat-work-pusher /root/wechat-work-pusher
ADD localtime /etc/localtime
CMD ["/root/wechat-work-pusher"]