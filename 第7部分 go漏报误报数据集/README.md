# Golang智能合约数据集

* 合约来源：GitHub

|合约|下载地址|漏洞|
|:--|:--|:--|
|0.go||chaincode-conversion-overflow，2 chaincode-math-overflow, black-listed-import, chaincode-pointer, chaincode-read-after-write|
|1.go||无漏洞|
|3.go||3 chaincode-pointer, chaincode-global-variable|
|4.go||2 chaincode-read-after-write, 3 chaincode-math-overflow（全误报）|
|6.go||2 chaincode-pointer, 3 chaincode-global-variable|
|11.go||无漏洞|
|14.go||chaincode-read-after-write, 4 chaincode-math-overflow（1误报）|
|17.go||2 chaincode-pointer, chaincode-global-variable|
|20.go||chaincode-global-variable|
|37.go||3 chaincode-global-variable, chaincode-global-variable|
|38.go||chaincode-global-variable, chaincode-read-after-write, chaincode-invoke-chaincode, 2 chaincode-privacy-leakage-in-arg, chaincode-privacy-leakage-in-ret|
|56.go||4 chaincode-math-overflow（全误报）, 5 chaincode-pointer, chaincode-go-routine|
|65.go||3 chaincode-pointer, chaincode-math-overflow（误报）, black-listed-import|
|76.go||4 chaincode-read-after-write, 6 chaincode-pointer|
|81.go||4 chaincode-pointer|
|92.go||4 chaincode-pointer, 2 chaincode-go-routine|
|103.go||3 chaincode-pointer, chaincode-global-variable|
|140.go||black-listed-import， chaincode-global-variable|
|147.go||4 chaincode-pointer, 2 chaincode-go-routine|
|148.go||4 chaincode-pointer|
|149.go||chaincode-privacy-leakage-in-arg, chaincode-read-after-write, chaincode-privacy-leakage-in-ret, chaincode-global-variable, chaincode-invoke-chaincode|
|153.go||2 black-listed-import|
|154.go||无漏洞|
|160.go||4 chaincode-pointer, 2 chaincode-math-overflow（全误报）, black-listed-import, chaincode-go-routine|
|163.go||无漏洞|
|168.go||5 chaincode-pointer, chaincode-privacy-leakage-in-ret, chaincode-read-after-write|
|169.go||无漏洞|
|175.go||chaincode-global-variable, 2 chaincode-go-routine|
|184.go||4 chaincode-conversion-overflow, 4 chaincode-math-overflow（2误报）, black-listed-import, chaincode-pointer|
|185.go||chaincode-conversion-overflow, 3 black-listed-import, 2 chaincode-pointer, chaincode-global-variable|
|187.go||chaincode-privacy-leakage-in-arg, 2 chaincode-privacy-leakage-in-ret, chaincode-read-after-write, chaincode-phandom-read, black-listed-import|
|188.go||无漏洞|
|189.go||chaincode-global-variable, 6 chaincode-pointer|
|190.go||5 chaincode-pointer, 4 chaincode-math-overflow （全误报）, chaincode-go-routine|
|191.go||2 chaincode-math-overflow, 2 chaincode-global-variable, chaincode-read-after-write, 2 chaincode-conversion-overflow, black-listed-import|
