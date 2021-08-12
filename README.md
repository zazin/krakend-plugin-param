## krakend plugin with custom param

1. docker run -d --name krakend -p 8000:8000 nurzazin/krakend-plugin:latest
2. curl -L -X GET 'http://localhost:8000/test?param=2'
3. curl -L -X GET 'http://localhost:8000/test?param=1'
