0. Go to (mnm.sh), install mnm, config it right.

1. `cd /etc/systemd/system`

2. `sudo vim leafer.service`, edit PATH below 
```
[Unit]
Description=Leafer
After=network.target

[Service]
User=root
Type=simple
ExecStart=/usr/bin/mnm run '[PATH]/leafer'
Restart=on-failure
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
```

3. sudo systemctl daemon-reload

4. sudo systemctl start [your-service.service]

5. sudo systemctl status [your-service.service]
