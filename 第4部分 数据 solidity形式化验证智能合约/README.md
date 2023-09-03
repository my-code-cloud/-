# Solidity智能合约形式化验证

* 合约来源：https://github.com/SRI-CSL/solidity/tree/boogie/test/solc-verify/examples

|合约|验证条件|预期结果|
|:--|:--|:--|
|function_precondition1.sol|precondition x == y|验证通过|
|function_precondition2.sol|precondition x == y|验证不通过|
|function_postcondition1.sol|postcondition x == (y + 1)|验证通过|
|function_postcondition2.sol|postcondition x == (y + 1)|验证不通过|
|contract_invariant1.sol|invariant x == y|验证通过|
|contract_invariant2.sol|invariant x == y|验证不通过|
|loop_invariant1.sol|invariant y <= x|验证通过|
|loop_invariant2.sol|invariant y <= x|验证不通过|
|modification_specifier1.sol|modifies x|验证通过|
|modification_specifier2.sol|modifies x|验证不通过|
|modification_specifier_condition1.sol|modifies y if n < 0|验证通过|
|modification_specifier_condition2.sol|modifies y if n < 0|验证不通过|
|uint_sum_function1.sol|invariant __verifier_sum_uint(balances) == totalSupply|验证通过|
|uint_sum_function2.sol|invariant __verifier_sum_uint(balances) == totalSupply|验证不通过|
|int_sum_function1.sol|invariant __verifier_sum_int(balances) == totalSupply|验证通过|
|int_sum_function2.sol|invariant __verifier_sum_int(balances) == totalSupply|验证不通过|
|uint_old_value_function1.sol|postcondition x == __verifier_old_uint(y) + 1|验证通过|
|uint_old_value_function2.sol|postcondition x == __verifier_old_uint(y) + 1|验证不通过|
|int_old_value_function1.sol|postcondition x == __verifier_old_int(y) + 1|验证通过|
|int_old_value_function2.sol|postcondition x == __verifier_old_int(y) + 1|验证不通过|
|inter_function1.sol|invariant enableTransfer ==> transferFrom|验证通过|
|inter_function2.sol|invariant enableTransfer ==> transferFrom|验证不通过|
|emit_specifier1.sol|emits e|验证通过|
|emit_specifier2.sol|emits e|验证不通过|
|ERC20_event1.sol|tracks-changes-in balanceOf|验证通过|
|ERC20_event2.sol|tracks-changes-in balanceOf|验证不通过|
|ERC20_totalsuppply1.sol|invariant __verifier_sum_uint(balanceOf) == totalSupply|验证通过|
|ERC20_totalsuppply2.sol|invariant __verifier_sum_uint(balanceOf) == totalSupply|验证不通过|
|ERC20_modifies1.sol|modifies balanceOf if balanceOf[msg.sender] >= amount|验证通过|
|ERC20_modifies2.sol|modifies balanceOf if balanceOf[msg.sender] >= amount|验证不通过|
