FROM devopsfaith/krakend:1.4.1 as krakend

WORKDIR /etc/krakend

COPY plugin/test-plugin.so /etc/krakend/plugin/
COPY krakend.json /etc/krakend/

ENTRYPOINT [ "/usr/bin/krakend" ]
CMD [ "run", "-c", "/etc/krakend/krakend.json" ]

EXPOSE 8000 8090
