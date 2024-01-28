# lavaprovider-monitor
Send discord alert when your lava provider is frozen or jailed

# Features
It will send alert in discord channel when your lava provider is frozen or jailed.
It will check the status every 10 minutes.
# Pre-requisites
1. You have go 1.20+ installed.
2. You have a discord channel and enable webhook You can follow this guide: https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks.
3. Prepare a working lava grpc
 
# How to use
Edit alert.toml

* webhook: Discord webhook url. https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks
* alertuserid: Discord user id. If you want be @ when receive alert in discord, set this. https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID-
* alertroleid: Discord role id. If you want the people who have special discord role be @ when receive alert in discord, set this. https://www.itgeared.com/how-to-get-role-id-on-discord/

Edit config.toml

* lavagrpc: A working lava grpc. For example: ip:9090
* chainid: Lava chain id. For current testnet, it is lava-testnet-2
* lavaprovideraddress: The lava provider address you wanna monitor
* chains: The chains you wanna monitor. For example: `["LAV1", "EVMOST", "EVMOS", "AXELAR", "AXELART"]`

# Build
go build -o lavaprovider-monitor  ./cmd
# Run
./lavaprovider-monitor
You can check the log in alert.log

