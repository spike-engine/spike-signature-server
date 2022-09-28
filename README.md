# spike-signature-server
Signature service

## Build and install spike-signature-server
1. Clone the repository
```shell
cd  $(go env GOPATH)/src/github.com
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
4. Update Config
```
cp config.toml.example config.toml
```
5. Run
```
./spike-signature-server
```

## Config 

1. Config file localtion.  
By default, spike-signature-server reads configuration from config.toml under the current folder. If it is not created, you may copy from config.config.example. If you need to specify the config file, you may set the enviornment as follows:
```
export ENV_CONFIG=~/spike_home/config-signature.toml
```
2. keystore file.  
If spike-signature-server cannot find keystore file with *walletFile* config under folder with *walletFolder* config, it will generate one for you. You may replace it with your own keystore file if needed.
3. Machine ID  
To run multiple signature services, you may assign each server a different machine ID with *machineId* config.

## Register spike as a system service
1. Link server into system binary path
```
sudo ln -s ./spike-signature-server /usr/local/bin
```
2. Copy config file into spike home
```
sudo mkdir -p /etc/spike/
sudo cp ./config.toml /etc/spike/config-signature.toml
```
3. Startup script
```shell
sudo vim /etc/systemd/system/spike-signature-server.service
```
Specify the path to the binary
```ini
[Service] 
ExecStart=/usr/local/bin/spike-signature-server
Environment=ENV_CONFIG=/etc/spike/config-signature.toml
Restart=always
RestartSec=5
```
```shell
systemctl daemon-reload
systemctl start spike-signature-server
```
Check Program Exec Log 
```shell
journalctl -u spike-signature-server.service  -f 
```
