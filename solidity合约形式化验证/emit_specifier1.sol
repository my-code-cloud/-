pragma solidity >=0.5.0;

contract C {
    /// @notice tracks-changes-in x
    event e();
    uint x = 0;

    /// @notice emits e
    function f() public {
        x = 1;
        emit e();
    }
}