# Golang智能合约漏洞

* 合约来源：https://github.com/my-code-cloud/CCrevive/tree/master/testdata/chaincode

|合约|漏洞|
|:--|:--|
|global-variable.go|chaincode-global-variable, chaincode-read-after-write|
|blacklist-imports.go|io, math/rand, net, os, time|
|range-over-map.go|chaincode-math-overflow, chaincode-range-over-map, chaincode-global-variable|
|goroutine.go|chaincode-go-routine|
|address.go|chaincode-pointer, chaincode-math-overflow|
|channel.go|chaincode-invoke-chaincode|
|phrantom.go|chaincode-phandom-read|
|read-write.go|chaincode-read-after-write|
|overflow2.go|chaincode-conversion-overflow|
