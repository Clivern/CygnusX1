#!/bin/bash

function peacock {
    echo "Installing peacock ..."

    apt-get install jq -y

    cd /etc/peacock

    LATEST_VERSION=$(curl --silent "https://api.github.com/repos/Clivern/peacock/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

    curl -sL https://github.com/Clivern/peacock/releases/download/v{$LATEST_VERSION}/peacock_{$LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz


    echo "[Unit]
Description=Peacock
Documentation=https://github.com/Clivern/peacock

[Service]
ExecStart=/etc/peacock/peacock server -c /etc/peacock/config.prod.yml
Restart=on-failure
RestartSec=2

[Install]
WantedBy=multi-user.target" > /etc/systemd/system/peacock.service

    systemctl daemon-reload
    systemctl enable peacock.service
    systemctl start peacock.service

    echo "peacock installation done!"
}

peacock
