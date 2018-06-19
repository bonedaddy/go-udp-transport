go-tcp-transport
==================

[![](https://img.shields.io/badge/made%20by-Protocol%20Labs-blue.svg?style=flat-square)](http://ipn.io)
[![](https://img.shields.io/badge/project-IPFS-blue.svg?style=flat-square)](http://libp2p.io/)
[![](https://img.shields.io/badge/freenode-%23ipfs-blue.svg?style=flat-square)](http://webchat.freenode.net/?channels=%23ipfs)
[![Coverage Status](https://coveralls.io/repos/github/libp2p/go-udp-transport/badge.svg?branch=master)](https://coveralls.io/github/libp2p/go-udp-transport?branch=master)
[![Travis CI](https://travis-ci.org/libp2p/go-udp-transport.svg?branch=master)](https://travis-ci.org/libp2p/go-udp-transport)

> A libp2p transport implementation for udp.


## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Contribute](#contribute)
- [License](#license)

## Install

go-udp-transport is a standard Go module which can be installed with:
```
go get github.com/libp2p/go-udp-transport
```

## Usage
This module is packaged with Gx. In order to use it in your own project it is recommended that you:

```
go get -u github.com/whyrusleeping/gx
go get -u github.com/whyrusleeping/gx-go
cd <your-project-repository>
gx init
gx import github.com/libp2p/go-udp-transport
gx install --global
gx-go --rewrite
```

## API
See https://godoc.org/github.com/libp2p/go-udp-transport

## Contribute

PRs are welcome!

Small note: If editing the Readme, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT
