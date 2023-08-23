contract ERC20 {
    uint public totalSupply;
    mapping(address => uint) public balanceOf;
    mapping(address => mapping(address => uint)) public allowance;
    string public name = "Solidity by Example";
    string public symbol = "SOLBYEX";
    uint8 public decimals = 18;

    event Transfer(address indexed from, address indexed to, uint value);

    event Approval(address indexed owner, address indexed spender, uint value);

    /// @notice modifies balanceOf if balanceOf[msg.sender] >= amount
    function transfer(address recipient, uint amount) external returns (bool) {
        require(balanceOf[msg.sender] >= amount);
        balanceOf[msg.sender] -= amount;
        balanceOf[recipient] += amount;
        emit Transfer(msg.sender, recipient, amount);
        return true;
    }
    
    /// @notice modifies allowance[msg.sender][spender]
    function approve(address spender, uint amount) external returns (bool) {
        allowance[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    /// @notice modifies balanceOf if allowance[sender][msg.sender] >= amount
    /// @notice modifies balanceOf if balanceOf[sender] >= amount
    /// @notice modifies allowance if allowance[sender][msg.sender] >= amount
    /// @notice modifies allowance if balanceOf[sender] >= amount
    function transferFrom(
        address sender,
        address recipient,
        uint amount
    ) external returns (bool) {
        require(allowance[sender][msg.sender] >= amount);
        require(balanceOf[sender] >= amount);
        allowance[sender][msg.sender] -= amount;
        balanceOf[sender] -= amount;
        balanceOf[recipient] += amount;
        emit Transfer(sender, recipient, amount);
        return true;
    }
    
    /// @notice modifies balanceOf[msg.sender]
    /// @notice modifies totalSupply
    function mint(uint amount) external {
        balanceOf[msg.sender] += amount;
        totalSupply += amount;
        emit Transfer(address(0), msg.sender, amount);
    }

    /// @notice modifies balanceOf if balanceOf[msg.sender] >= amount
    /// @notice modifies totalSupply if balanceOf[msg.sender] >= amount
    function burn(uint amount) external {
        require(balanceOf[msg.sender] >= amount);
        balanceOf[msg.sender] -= amount;
        totalSupply -= amount;
        emit Transfer(msg.sender, address(0), amount);
    }
}
