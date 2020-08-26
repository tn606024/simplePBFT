Simple PBFT
------

This repository contains the golang code of simple pbft consensus implementation.

  
How to run
------

## Build

```shell script
go build 
```

### Start four pbft node

```shell script
./simplePBFT pbft node -id 0
./simplePBFT pbft node -id 1
./simplePBFT pbft node -id 2
./simplePBFT pbft node -id 3
```

### Start pbft client to send message

```shell script
./simplePBFT pbft client
```


### Reference

- https://www.jianshu.com/p/78e2b3d3af62
