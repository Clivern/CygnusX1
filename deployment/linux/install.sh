#!/bin/bash

function kevent {
    echo "Installing kevent ..."

    apt-get install jq -y

    cd /etc/kevent

    LATEST_VERSION=$(curl --silent "https://api.github.com/repos/Clivern/kevent/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

    curl -sL https://github.com/Clivern/kevent/releases/download/v{$LATEST_VERSION}/kevent_{$LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz


    echo "[Unit]
Description=Kevent
Documentation=https://github.com/Clivern/kevent

[Service]
ExecStart=/etc/kevent/kevent server -c /etc/kevent/config.prod.yml
Restart=on-failure
RestartSec=2

[Install]
WantedBy=multi-user.target" > /etc/systemd/system/kevent.service

    systemctl daemon-reload
    systemctl enable kevent.service
    systemctl start kevent.service

    echo "kevent installation done!"
}

kevent
