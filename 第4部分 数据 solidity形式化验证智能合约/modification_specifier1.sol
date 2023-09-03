pragma solidity >=0.5.0;

contract C {
    int x;
    int y;

    /// @notice modifies x
    function add_to_x(int n) internal {
        x = x + n;
        require(x >= y); // Ensures that there is no overflow
    }

    /// @notice modifies x
    /// @notice modifies y
    function add(int n) public {
        require(n >= 0);
        add_to_x(n);
        while (y < x) {
            y = y + 1;
        }
    }
}