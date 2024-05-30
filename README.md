# lavaprovider-monitor
Send discord alert when your lava provider is frozen or jailed

# Features
It will send alert in discord channel when your lava provider is frozen or jailed.
It will check the status every 10 minutes.
# Pre-requisites
1. You have go 1.20+ installed.
2. You have a discord channel and enable webhook. You can follow this guide: https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks.
3. Prepare a working lava grpc
 
# How to use

## Build
```bash
cd $HOME
git clone https://github.com/silentnoname/lavaprovider-monitor
cd lavaprovider-monitor
go build -o lavaprovider-monitor  ./cmd
```

## Fill the config
```bash
cp config.toml.example config.toml
cp alert.toml.example alert.toml
```

Edit alert.toml

* webhook: Discord webhook url. https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks
* alertuserid: Discord user id. If you want be @ when receive alert in discord, set this. https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID-
* alertroleid: Discord role id. If you want the people who have special discord role be @ when receive alert in discord, set this. https://www.itgeared.com/how-to-get-role-id-on-discord/

Edit config.toml

* lavagrpc: A working lava grpc. For example: ip:9090
* chainid: Lava chain id. For current testnet, it is lava-testnet-2
* lavaprovideraddress: The lava provider address you wanna monitor
* chains: The chains you wanna monitor. For example: `["LAV1", "EVMOST", "EVMOS", "AXELAR", "AXELART"]`

## Run as a service
``` 
sudo tee /etc/systemd/system/lavaprovider-monitor.service > /dev/null << EOF
[Unit]
Description=lavaprovider-monitor
After=network-online.target

[Service]
User=$USER
WorkingDirectory=$HOME/lavaprovider-monitor
ExecStart=$HOME/lavaprovider-monitor/lavaprovider-monitor
Restart=always
RestartSec=3
LimitNOFILE=65535

[Install]
WantedBy=multi-user.target

EOF

```
start
```shell
sudo systemctl enable lavaprovider-monitor
sudo systemctl daemon-reload
sudo systemctl start lavaprovider-monitor
```

check logs
```shell
sudo journalctl -fu  lavaprovider-monitor -o cat
```


If your provider is frozen, you will see this in your discord channel

![](https://i.imgur.com/9wVbRl7.jpg)