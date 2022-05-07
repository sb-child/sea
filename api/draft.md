
```
water id(save in DB) -> token(save in DB) -> session(save in memory) -> data
unique server id        unique peer id for   unique id for every        raw data binary
                        p2p connection       tcp connection, delete
                                             after connection closed
```
