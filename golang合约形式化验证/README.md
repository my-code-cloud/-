# Golang智能合约形式化验证

* 合约来源：https://github.com/my-code-cloud/CCrevive/tree/master/testdata

|合约|验证条件|预期结果|
|:--|:--|:--|
|context-key-type.go|context-keys-type|验证不通过|
|time-equal.go|time-equal|验证不通过|
|errorf.go|errorf|验证不通过|
|blank-import-lib.go|blank-imports|验证不通过|
|context-as-argument.go|context-as-argument|验证不通过|
|receiver-naming.go|receiver-naming|验证不通过|
|get-return.go|get-return|验证不通过|
|modifies-parameter.go|modifies-parameter|验证不通过|
|unused-parameter.go|unused-parameter|验证不通过|
|unreachable-code.go|unreachable-code|验证不通过|
|range-val-in-closure.go|range-val-in-closure|验证不通过|
|range-val-address.go|range-val-address|验证不通过|
|atmic.go|atomic|验证不通过|
|bare-return.go|bare-return|验证不通过|
|unhandled-error.go|unhandled-error|验证不通过|
|string-of-int.go|string-of-int|验证不通过|
|unconditional-recursion.go|unconditional-recursion|验证不通过|
|defer.go|defer|验证不通过|
|unexported-naming.go|unexported-naming|验证不通过|
|datarace.go|datarace|验证不通过|
|waitgroup-by-value.go|waitgroup-by-value|验证不通过|
|modifies-value-receiver.go|modifies-value-receiver|验证不通过|
|unused-receiver.go|unused-receiver|验证不通过|

