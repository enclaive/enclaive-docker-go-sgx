# Go-SGX Demonstration

The goal of the demonstration is to show that files in a confidential container are protected. In the example case we run an HTTPS server written in golang
and protect the TLS X509 server certificate private key `/tmp/server.key`. We show that in the confidential container the key is inaccessible while in the standard container it is. The rationality is that the confidential container mounts in the enclave the file system `/tmp/server.key` being accessible only within the enclave.

## Building and Running

```sh
docker-compose up
```

## observe that it's possible to extract the private key

```sh
docker exec golang cat /tmp/server.key
-----BEGIN PRIVATE KEY-----

```

## but not with the sgx encrypted container

```sh
docker exec golang-sgx cat /tmp/server.key
cat: /tmp/server.key: No such file or directory
```
