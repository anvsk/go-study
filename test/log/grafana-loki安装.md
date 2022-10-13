### Gafana  Loki
官方安装文档
https://grafana.com/docs/loki/latest/installation/local/
Loki is the logging engine.
Promtail sends logs to Loki.

1. install grafana 
```bash
wget https://dl.grafana.com/enterprise/release/grafana-enterprise-6.0.0-1.x86_64.rpm
sudo yum install grafana-enterprise-6.0.0-1.x86_64.rpm
```
> 3000 端口、用户名密码默认admin

2. install loki 

```bash
curl -O -L "https://github.com/grafana/loki/releases/download/v2.4.2/loki-linux-amd64.zip"
# extract the binary
unzip "loki-linux-amd64.zip"
# make sure it is executable
chmod a+x "loki-linux-amd64"
# get config file
wget https://raw.githubusercontent.com/grafana/loki/master/cmd/loki/loki-local-config.yaml

# start
./loki-linux-amd64 -config.file=loki-local-config.yaml

```
> 3100 端口


3. install promtail (采集客户端，多节点都要装)

```bash
sudo yum install -y systemd-devel
curl -O -L "https://github.com/grafana/loki/releases/download/v2.4.2/promtail-linux-amd64.zip"
# extract the binary
unzip "promtail-linux-amd64.zip"
# make sure it is executable
chmod a+x "promtail-linux-amd64"
# get config file
wget https://raw.githubusercontent.com/grafana/loki/main/clients/cmd/promtail/promtail-local-config.yaml

# start
./promtail-linux-amd64 -config.file=promtail-local-config.yaml

```

#### 配置文件demo(官方`promtail-local-config.yaml`有bug)
`loki-local-config.yaml`
```yaml
auth_enabled: false

server:
  http_listen_port: 3100
  grpc_listen_port: 9096

common:
  path_prefix: /tmp/loki
  storage:
    filesystem:
      chunks_directory: /tmp/loki/chunks
      rules_directory: /tmp/loki/rules
  replication_factor: 1
  ring:
    instance_addr: 127.0.0.1
    kvstore:
      store: inmemory

schema_config:
  configs:
    - from: 2020-10-24
      store: boltdb-shipper
      object_store: filesystem
      schema: v11
      index:
        prefix:
        period: 24h

ruler:
  alertmanager_url: http://localhost:9093

```

`promtail-local-config`
```yaml
server:
  http_listen_port: 9080
  grpc_listen_port: 0

positions:
  filename: /tmp/positions.yaml

clients:
  configs:
    - url: http://localhost:3100/loki/api/v1/push

scrape_configs:
  - job_name: system
    static_configs:
      - targets:
          - localhost
        labels:
          job: test
          __path__: /Users/pingan/workspace/go/src/go-study/cmd/test/log/tmp/*log

```

