## price_exporter
Prometheus exporter for cryptocurrency prices.  
(default serivce port: 61988)


## List of APIs
- USD
  - USD/KRW(Dunamu)
- BTC
  - BTC/KRW(Upbit)
  - BTC/UDST(Upbit, Binance, HoubiGlobal)
- ATOM
  - ATOM/USDT(Binance, HuobiGlobal)
  - ATOM/KRW(Coinone, Upbit)
  - ATOM/BTC(Binance, Upbit, HuobiGlobal)
- IRIS
  - IRIS/USDT(HuobiGlobal)
  - IRIS/BTC(HuobiGlobal)
- KAVA
  - KAVA/USDT(Binance)
  - KAVA/KRW(Coinone)
  - KAVA/BTC(Binance)
- SOL
  - SOL/BUSD(Binance)
  - SOL/BTC(Binance)
  

## Install
```bash
mkdir price_exporter && cd price_exporter 

wget https://github.com/node-a-team/price_exporter/releases/download/v0.2.2/price_exporter.tar.gz  && sha256sum price_exporter.tar.gz | fgrep 2642adbeee786f9b752ff6a4d909a3a38647cb89ce918b2635d8e347766c358e && tar -zxvf price_exporter.tar.gz ||  echo "Bad Binary!"
```


## Start
  
```bash
./price_exporter {path to config.toml}

// ex)
./price_exporter /data/monitoring/price_exporter
```


## Use systemd service
  
```sh
# $HOME: /data/monitoring
# Path to config.toml: /data/monitoring/price_exporter
sudo tee /etc/systemd/system/price_exporter.service > /dev/null <<EOF
[Unit]
Description=Price Exporter
After=network-online.target

[Service]
User=monitoring
WorkingDirectory=/data/monitoring/price_exporter
ExecStart=/data/monitoring/price_exporter/price_exporter \
        /data/monitoring/price_exporter
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=price-exporter
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
EOF

sudo systemctl enable price_exporter.service
sudo systemctl start price_exporter.service
