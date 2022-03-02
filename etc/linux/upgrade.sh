#!/bin/bash

function copper {
    echo "Upgrade copper ..."

    cd /etc/copper
    mv config.prod.yml config.back.yml

    LATEST_VERSION=$(curl --silent "https://api.github.com/repos/Clivern/copper/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

    curl -sL https://github.com/Clivern/copper/releases/download/v{$LATEST_VERSION}/copper_{$LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz

    rm config.prod.yml
    mv config.back.yml config.prod.yml

    systemctl restart copper

    echo "copper upgrade done!"
}

copper
