# view cert details
`````
openssl x509 -noout -text -in CA.crt
openssl x509 -noout -text -in Server.crt
`````

# generate certs for server
`````
openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes \
  -keyout Server.key -out Server.crt \
  -subj '/C=DE/ST=Berlin/L=Berlin/O=wsva/CN=Server' -addext 'subjectAltName=IP:10.0.0.1'
`````

# Notes
浏览器报错：此服务器无法证实它就是 10.0.0.1 - 它的安全证书没有指定主题备用名称。

解决方法：
需要在生成证书时除了配置CN，还要配置subjectAltName
网上有许多方法，这里使用openssl 1.1.1才支持的-addtext方式