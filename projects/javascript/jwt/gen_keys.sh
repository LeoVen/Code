rm jwtRS256.key jwtRS256.key.pub

ssh-keygen -t rsa -b 4096 -m PKCS8 -f jwtRS256.key -N ''

# Don't add passphrase
openssl rsa -in jwtRS256.key -pubout -outform PEM -out jwtRS256.key.pub
