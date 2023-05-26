# ibcswap
A protocol that implement IBC atomic swap 

## run it as a Docker
build the image:

```
docker build -t simd .

```

run the image for tests

```
docker run --rm -it -p 26656:26656 -p 26657:26657 simd

```
inspect the container if needed. 

```
docker run --rm -it -p 26656:26656 -p 26657:26657l--entrypoint /bin/bash simd

```
