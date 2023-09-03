# Solidity智能合约形式化验证

* 合约来源：https://github.com/SRI-CSL/solidity/tree/boogie/test/solc-verify/examples

|合约|验证条件|预期结果|
|:--|:--|:--|
|ERC20_event1.sol|tracks-changes-in balanceOf|验证通过|
|ERC20_event2.sol|tracks-changes-in balanceOf|验证不通过|
|ERC20_totalsuppply1.sol|invariant __verifier_sum_uint(balanceOf) == totalSupply|验证通过|
|ERC20_totalsuppply2.sol|invariant __verifier_sum_uint(balanceOf) == totalSupply|验证不通过|
|ERC20_modifies1.sol|modifies balanceOf if balanceOf[msg.sender] >= amount|验证通过|
|ERC20_modifies2.sol|modifies balanceOf if balanceOf[msg.sender] >= amount|验证不通过|
