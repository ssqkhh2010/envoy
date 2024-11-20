FROM ubuntu:22.04
# 设置代理环境变量
ENV use_proxy=on \
    no_proxy=localhost,127.0.0.1,*.example.com \
    http_proxy=http://192.168.16.3:22222 \
    https_proxy=http://192.168.16.3:22222

RUN apt update && apt install -y \
    bash \
    curl \
    unzip \
    vim \
    git \
    wget \
    net-tools \
    python3 \
    telnet \
    s3fs \
    fuse \
    dnsutils \
    wrk

ENV use_proxy=off \
    no_proxy= \
    http_proxy= \
    https_proxy=

WORKDIR /app
COPY ./scclient /app/
RUN chmod +x /app/scclient
EXPOSE 8184
CMD ["/app/scclient"]