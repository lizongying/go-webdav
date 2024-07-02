# go-webdav

可以很方便地開啟一個本地webdav服務

使用場景：

* 跨設備文件管理
* 跨設備播放媒體資源。比如可以在電視上播放電腦上的生活視頻

## usage

* 下載可執行文件
* 修改配置文件，需要配置鏈接地址和文件夾地址
* 運行。可設置配置文件地址--config
* 可選https，具體配置可以參考example.yml

## build

```shell
make
```

## run

```shell
go run ./cmd/webdav_server/* --config example.yml
```

```shell
./releases/webdav_server --config example.yml
```

## test

```shell
# genera a tls
go run ./tools/tls_generator/*.go -s -h webdav -i 10.0.2.2

# http
curl -v -k -u username:password -X PROPFIND http://localhost:9876/movie/ -o out.xml

# https
curl -v -k -u username:password -X PROPFIND https://localhost:9876/music/ -o out.xml
curl -v -k -u username:password -X PROPFIND https://localhost:9876/movie/ -o out.xml
curl -v -k -u username:password https://localhost:9876/movie/小蒙牛满月了.m4v -o 小蒙牛满月了.m4v
ffplay https://username:password@localhost:9876/movie/小蒙牛满月了.m4v
```
