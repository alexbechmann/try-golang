# Note: You can use any Debian/Ubuntu based image you want. 
FROM mcr.microsoft.com/devcontainers/base:bullseye

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

RUN apt-get update && apt-get install -y \
    build-essential \
    cmake \
    libssl-dev \
    libsasl2-dev \
    libzstd-dev \
    zlib1g-dev \
    pkg-config

# Install NodeJS
RUN apt-get update && apt-get install \
  apt-transport-https \
  curl \
  -y
RUN set -uex; \
    apt-get update; \
    apt-get install -y ca-certificates curl gnupg; \
    mkdir -p /etc/apt/keyrings; \
    curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key \
     | gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg; \
    NODE_MAJOR=18; \
    echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_$NODE_MAJOR.x nodistro main" \
     > /etc/apt/sources.list.d/nodesource.list; \
    apt-get -qy update; \
    apt-get -qy install nodejs;
RUN apt-get update && apt-get install nodejs -y

RUN SNIPPET="export PROMPT_COMMAND='history -a' && export HISTFILE=/commandhistory/.bash_history" \
  && mkdir /commandhistory \
  && touch /commandhistory/.bash_history \
  && chown -R vscode /commandhistory \
  && echo $SNIPPET >> "/home/vscode/.bashrc"

# RUN apt-get update && apt install libssl-dev -y
# RUN apt-get update && apt-get install pkg-config -y
RUN npm i -g quicktype
RUN npm i -g foreman

RUN apt-get install bison mercurial -y

# Protoc
COPY ./install-protoc.sh ./
RUN bash ./install-protoc.sh

USER vscode

# ENV GVM_ROOT /home/vscode/.gvm
# RUN curl https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer | bash
# RUN /home/vscode/.gvm/bin/gvm install go1.20.5 -B
# RUN /home/vscode/.gvm/bin/gvm use go1.20.5 --default