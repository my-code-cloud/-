pragma solidity >=0.5.0;

/// @notice invariant __verifier_sum_uint(balances) == totalSupply
contract Token {
    mapping (address => uint) balances;
    uint totalSupply;

    constructor() {
        totalSupply = 7000000000 * (10**18);
        balances[msg.sender] = totalSupply; // Give the creator all initial tokens
    }

    function transfer(address receiver, uint amount) public {
        //balances[msg.sender] += amount;
        balances[receiver] -= amount;
    }
}