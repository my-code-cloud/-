pragma solidity >=0.5.0;

contract C {
    uint x;
    uint y;

    /// @notice postcondition x == __verifier_old_uint(y) + 2
    function add_to_x(uint n) public {
        x = y + 1;
        require(x >= y); // Ensures that there is no overflow
        y = y - 3;
    }

    function add(uint n) public {
        require(n >= 0);
        add_to_x(n);
        while (y < x) {
            y = y + 1;
        }
    }
}