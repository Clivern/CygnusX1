#!/bin/bash

function peacock {
    echo "Upgrade peacock ..."

    cd /etc/peacock
    mv config.prod.yml config.back.yml

    LATEST_VERSION=$(curl --silent "https://api.github.com/repos/Clivern/peacock/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

    curl -sL https://github.com/Clivern/peacock/releases/download/v{$LATEST_VERSION}/peacock_{$LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz

    rm config.prod.yml
    mv config.back.yml config.prod.yml

    systemctl restart peacock

    echo "peacock upgrade done!"
}

peacock
