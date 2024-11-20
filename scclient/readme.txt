https://registry.hub.docker.com

https://hub.docker.com/
316203121@qq.com
123456@Aa

docker build -t ssqkhh/scclient:latest .

cd /app/jk/envoyfilter/dedup/psbc_sidecar_multi/scclient/
/app/jk/dedup/dedup/psbc_sidecar_multi/scclient
#kubectl create ns edaapp
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o scclient/scclient scclient/main_cleint.go
cd scclient/
docker build -t scclient .
docker build -t scclient Dockerfile-uid
kind load docker-image scclient:latest --name c1


kubectl delete -f deployment.yaml
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml

curl http://172.18.0.2:30921/send -H "dc:f"


apt-get install s3fs
apk add s3fs
echo FPSM0DNI8KBL59VQSAAH:qQ4MO3lN9sUJUVFvBRAgpqZTvmMwuVhP5EHqFJFI > /root/.passwd-s3fs
chmod 600 ~/.passwd-s3fs
mkdir /app/mountpoint
#`use_path_request_style`：指定请求的类型，对于某些 S3 兼容服务这是必需的。
s3fs new-bucket /app/mountpoint -o url=https://minio.edaapp:443 -o use_path_request_style -o passwd_file=~/.passwd-s3fs -o no_check_certificate
#调试模式
s3fs new-bucket /app/mountpoint -o url=https://minio.edaapp:443 -o use_path_request_style -o passwd_file=/root/.passwd-s3fs -o no_check_certificate -d -f
./mc ls minio-tenant-1/new-bucket --insecure


需在宿主机上执行
modprobe fuse
modprobe: FATAL: Module fuse not found in directory /lib/modules/6.5.0-14-generic
apt-get install fuse






#dnsutils 在 CentOS 中相应的包是 bind-utils，它包含 nslookup 等工具。
RUN yum install -y \
    bash \
    curl \
    unzip \
    vim \
    git \
    wget \
    net-tools \
    python3 \
    telnet \
    s3fs-fuse \
    fuse \
    bind-utils








https://github.com/moby/buildkit/releases/download/v0.15.1/buildkit-v0.15.1.linux-arm64.tar.gz
wget https://github.com/moby/buildkit/releases/download/v0.15.1/buildkit-v0.15.1.linux-amd64.tar.gz

tar -xvf buildkit-v0.15.1.linux-amd64.tar.gz
sudo cp bin/buildctl /usr/local/bin/
sudo cp bin/buildkitd /usr/local/bin/

sudo chmod +x /usr/local/bin/buildctl
sudo chmod +x /usr/local/bin/buildkitd

sudo buildkitd &

buildctl --version
buildkitd --version


nerdctl build -t scclient .




docker build -t scclient -f Dockerfile-uid .
/Users/lz/Downloads/e/javautils/go/dedup/dedup/psbc_sidecar_multi/scclient/
/app/jk/envoyfilter/dedup/psbc_sidecar_multi/scserver



https://docs.nvidia.com/deeplearning/triton-inference-server/user-guide/docs/customization_guide/compose.html






wrk -t1 -c10 -d60s -R10 --latency http://10.244.0.13/get

docker run -it --rm ubuntu:23.04 bash

vi /etc/docker/daemon.json
{
  "insecure-registries" : ["localhost:32000"],
  "dns": ["8.8.8.8", "8.8.4.4"]
}


docker run --network host -it --rm ubuntu:23.04 bash
export use_proxy=on
export no_proxy=localhost,127.0.0.1,*.example.com,10.0.0.0/8,172.0.0.0/8,192.168.0.0/16
export http_proxy=http://120.244.127.116:22222
export https_proxy=http://120.244.127.116:22222

apt update
apt install -y bash curl unzip vim git wget net-tools python3 telnet s3fs fuse dnsutils wrk




kubectl create deployment scclient-deployment --image ssqkhh/scclient:latest
