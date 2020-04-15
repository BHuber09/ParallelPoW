# ParallelPoW
Golang implementation of a Parallel PoW algorithm

--- Parallel Run --- 
$ go run main.go
Number Miners = 5
Num Blocks = 6

Block: 0
Timestamp: 1586893758
Data: 0
Hash: nil
PrevHash: GENESIS
Header: 0
Miner: nil

Block: 1
Timestamp: 1586893758
Data: 1
Hash: 00000RuivUgUwJF60k1h5csdNHc=
PrevHash: nil
Header: 1375542074
Miner: Miner3

Block: 2
Timestamp: 1586893965
Data: 2
Hash: 00000p8nQjhwsSYo2qvLW0djvec=
PrevHash: 00000RuivUgUwJF60k1h5csdNHc=
Header: 91024741
Miner: Miner0

Block: 3
Timestamp: 1586894358
Data: 3
Hash: 00000c8g6kKlT4rikI-hiT4FiAo=
PrevHash: 00000p8nQjhwsSYo2qvLW0djvec=
Header: 1216954569
Miner: Miner2

Block: 4
Timestamp: 1586895990
Data: 4
Hash: 00000m39D3T0dCAl9_I5RuCwgIQ=
PrevHash: 00000c8g6kKlT4rikI-hiT4FiAo=
Header: 1587034130
Miner: Miner3

Block: 5
Timestamp: 1586897086
Data: 5
Hash: 00000y11n61hZqITjIDNK_opOV4=
PrevHash: 00000m39D3T0dCAl9_I5RuCwgIQ=
Header: 1756825109
Miner: Miner4



--- Serial Run ---
$ go run main.go
Number Miners = 5
Num Blocks = 6

Block: 0
Timestamp: 1586893756
Data: 0
Hash: nil
PrevHash: GENESIS
Header: 0
Miner: nil

Block: 1
Timestamp: 1586893756
Data: 1
Hash: 000005q4LoQPZMGNzVgxTAKAGCo=
PrevHash: nil
Header: -2003160881
Miner: Miner0

Block: 2
Timestamp: 1586894105
Data: 2
Hash: 00000n_liH0_Zv1nO99lgvbNbgQ=
PrevHash: 000005q4LoQPZMGNzVgxTAKAGCo=
Header: -2082123139
Miner: Miner1

Block: 3
Timestamp: 1586894393
Data: 3
Hash: 00000C8q4wnh1vlyY43IbqPqaFI=
PrevHash: 00000n_liH0_Zv1nO99lgvbNbgQ=
Header: -2147475996
Miner: Miner1

Block: 4
Timestamp: 1586894393
Data: 4
Hash: 00000lhArduyK1OgafzmQFz0F78=
PrevHash: 00000C8q4wnh1vlyY43IbqPqaFI=
Header: -1494300028
Miner: Miner1

Block: 5
Timestamp: 1586897471
Data: 5
Hash: 00000Hl3xSzfY6QjkYaqevwyhwU=
PrevHash: 00000lhArduyK1OgafzmQFz0F78=
Header: -2090114942
Miner: Miner2
