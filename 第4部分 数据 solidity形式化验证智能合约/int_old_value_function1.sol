pragma solidity >=0.5.0;

contract C {
    int x;
    int y;

    /// @notice postcondition x == __verifier_old_int(y) + 1
    function add_to_x(int n) public {
        x = y + 1;
        require(x >= y); // Ensures that there is no overflow
        y = y - 3;
    }

    function add(int n) public {
        require(n >= 0);
        add_to_x(n);
        while (y < x) {
            y = y + 1;
        }
    }
}