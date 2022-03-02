#!/bin/bash

function copper {
    echo "Installing copper ..."

    apt-get install jq -y

    cd /etc/copper

    LATEST_VERSION=$(curl --silent "https://api.github.com/repos/Clivern/copper/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

    curl -sL https://github.com/Clivern/copper/releases/download/v{$LATEST_VERSION}/copper_{$LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz


    echo "[Unit]
Description=Copper
Documentation=https://github.com/Clivern/copper

[Service]
ExecStart=/etc/copper/copper server -c /etc/copper/config.prod.yml
Restart=on-failure
RestartSec=2

[Install]
WantedBy=multi-user.target" > /etc/systemd/system/copper.service

    systemctl daemon-reload
    systemctl enable copper.service
    systemctl start copper.service

    echo "copper installation done!"
}

copper
