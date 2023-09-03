pragma solidity ^0.7.0;

/// @notice invariant enableTransfer ==> transferFrom
contract DaiToken {
    mapping(address => uint) public balances;
    bool public transferEnabled;
    
    function enableTransfer() public {
        transferEnabled = true;
    }

    /// @notice precondition transferEnabled == false
    function transferFrom(address _from, address _to, uint _value) public {
        require(transferEnabled);
        balances[_from] -= _value;
        balances[_to] += _value;
    }
}
