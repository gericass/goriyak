# goriyak

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![Go Report Card](https://goreportcard.com/badge/github.com/gericass/goriyak)](https://goreportcard.com/report/github.com/gericass/goriyak)

<img height="300px" src="https://github.com/gericass/goriyak/blob/img/img/goriyak.png"
 alt="goritak logo" title="goriyak" align="right" />

goriyak is a blockchain based on hash calculation on smartphone

goriyak provides the following functions:

1. High speed communication interface by using gRPC.
2. Synchronization of information in the management network and fast recovery.
3. Consensus algorithm depends on low-performance clients such as smartphones.
4. Rational transaction management mechanism.

## Specifications
 
|                         |     Uses    | 
|-------------------------|-------------| 
| Language                | Go          | 
| Blockchain type         | private     | 
| Protocol                | gRPC        | 
| Consensus algorithm     | PoW(Custom) | 
| Data management(public) | Riak        | 
| Data management(local)  | MySQL       | 
| Hash                    | SHA-256     | 

## Requirements

- Go 1.9 or later
- Riak KV 2.2.3 or later
- gRPC(Protocol buffer)
- MySQL 5.7 or later

or 

- docker
- docker-compose

## Install

If you uses `docker-compose`

Run `docker-compose up`


