
```
water id(save in DB) -> token(save in DB) -> session(save in memory) -> data
unique server id        unique peer id for   unique id for every        raw data binary
                        p2p connection       tcp connection, delete
                                             after connection closed
```

```
s1 <-> s2
[s1]    [s2]    [s3]
  |      |
  +--+---+
     \- water id > token > session
```

```
s1 <-> s3, s2 as a bridge
             /- water id > token
          +--+----+
          |       |
[s1]    [s2]    [s3]
  |      |
  +--+---+
     \- water id > token

s1, s3 have session key
```
