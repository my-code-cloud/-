pragma solidity >=0.5.0;

contract C {
    int x;
    int y;

    /// @notice postcondition x == (y + 1)
    function add_to_x(int n) internal {
        x = y + 2;
        require(x >= y); // Ensures that there is no overflow
    }

    function add(int n) public {
        require(n >= 0);
        add_to_x(n);
        while (y < x) {
            y = y + 1;
        }
    }
}