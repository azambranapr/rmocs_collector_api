# Certificates

# Key Gen

## create private key:
````
openssl genrsa -out app.rsa 1024
openssl genrsa -out app.rsa 2048
openssl genrsa -out app.rsa 4096
````

## create public key:
````
openssl rsa -in app.rsa -out app.rsa.pub -pubout -outform PEM
````

## ssh key rsa (GitHub)
````
ssh-keygen -t rsa -f key.rsa -b 1024
ssh-keygen -t rsa -f key.rsa -b 2048
ssh-keygen -t rsa -f key.rsa -b 4096
````
