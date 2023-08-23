pragma solidity >=0.5.0;

contract C {
    int x;
    int y;

    function add_to_x(int n) internal {
        x = x + n;
        require(x >= y); // Ensures that there is no overflow
    }

    function add(int n) public {
        require(n >= 0);
        add_to_x(n);
        /// @notice invariant y <= x
        while (y < x) {
            y = y + 1;
        }
    }
}