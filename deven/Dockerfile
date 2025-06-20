FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive

ARG UPSTREAM_REPO="zHanyo/ultimate-setup"
ARG RELEASE_NAME="ultimate-setup-linux-amd64"

COPY requirements-apt.txt /tmp/requirements-apt.txt
COPY requirements-pip.txt /tmp/requirements-pip.txt
COPY requirements-tools.sh /tmp/requirements-tools.sh

RUN apt-get update && \
    xargs -a /tmp/requirements-apt.txt apt-get install -y && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

RUN pip3 install --no-cache-dir -r /tmp/requirements-pip.txt

RUN bash /tmp/requirements-tools.sh

RUN set -eux; \
    url=$(curl -sL https://api.github.com/repos/${UPSTREAM_REPO}/releases/latest | \
    grep browser_download_url | grep ${RELEASE_NAME} | cut -d '"' -f 4); \
    curl -L "$url" -o /usr/local/bin/ultimate-setup && chmod +x /usr/local/bin/ultimate-setup

WORKDIR /workspace
CMD ["/bin/bash"]
