# spike-signature-server

## Build and install spike-signature-server
1. Clone the repository
```shell
cd  /root/go/src/github.com
git clone https://github.com/spike-engine/spike-signature-server.git
cd  spike-signature-server
```
2. Install all the dependencies
```shell
go mod tidy
```
3. Make build
```shell
go build -o spike-signature-server ./main.go
```
4. Startup script
```shell
vim /etc/systemd/system/spike-signature-server.service
```
Specify the path to the binary
```markdown
[Service] 
ExecStart=/root/go/src/github.com/spike-blockchain-server/spike-frame-server
Environment=CONFIG_PATH=/etc/config.toml
Restart=always
RestartSec=5
```
```shell
systemctl daemon-reload
systemctl start spike-frame-server
```
Check Program Exec Log 
```shell
journalctl -u spike-signature-server.service  -f 
```

Of course, you can click the build icon in your IDE to run the project instead of startup script.
But, we recommend using system script in mainnet.

### config
If you don't specify the path to the configuration file in the environment variable in the startup script,
config.toml is the default.And config.toml is also a demo.

You should configure some information about system port, GM wallet.

Regarding GM wallet, if you don't prepare keystore.json in advance, the program will be automatically generated according to your configuration.
