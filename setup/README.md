# Steps

0. postgre.md

Install postgresql, create db and import table

1. mixin.md

Get conversationID
Edit AcceptAsset


2. telegram.md

Config telegram

3. signal.sh

Install signal-cli and rest-api

4. twilio.md

Config twilio audio


5. https.md

enable https


6. service.md

add a systemd service



# Depoly after editing 

## go

go build 
scp -r leafer config.toml schema.sql hk:/root/dev/leafer/


## web
scp -r web/dist hk:/root/dev/leafer/web/
