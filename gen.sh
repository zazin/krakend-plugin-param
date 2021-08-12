go build -buildmode=plugin -o plugin/test-plugin.so ./plugin

docker build -t nurzazin/krakend-plugin .
