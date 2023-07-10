Steps to generate custom CA private key and certificate
1. openssl genpkey -algorithm RSA -out ca.key
2. openssl req -x509 -new -nodes -key ca.key -sha256 -days 365 -out ca.crt -subj "/CN=Custom CA"

Steps to generate server private key and certificate
1. openssl req -newkey rsa:2048 -nodes -keyout server.key -out server.csr -subj "/CN=localhost"
2. echo "subjectAltName = DNS:localhost" > server.ext
3. openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365 -extfile server.ext

Steps to generate client private key and certificate (needed for mTLS)
1. openssl req -newkey rsa:2048 -nodes -keyout client.key -out client.csr -subj "/CN=localhost"
2. echo "subjectAltName = DNS:localhost" > client.ext
3. openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365 -extfile client.ext