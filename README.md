SGX golang demonstration


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
