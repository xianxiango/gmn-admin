
pragma solidity ^0.4.0;
contract MyToken {
    /* This creates an array with all balances */
    mapping (address => uint256) public balanceOf;
    /* Public variables of the token */
    string public name = "SOTE";
    string public symbol = "S";
    uint8 public decimals = 2;

    /* This generates a public event on the blockchain that will notify clients */
    event Transfer(address indexed from, address indexed to, uint256 value);

    /* Initializes contract with initial supply tokens to the creator of the contract */
    function MyToken() public {

    }

    /* Send coins */
    function transfer(address _to, uint256 _value) public {
        /* if the sender doenst have enough balance then stop */
        if (balanceOf[msg.sender] < _value) return;
        if (balanceOf[_to] + _value < balanceOf[_to]) return;

        /* Add and subtract new balances */
        balanceOf[msg.sender] -= _value;
        balanceOf[_to] += _value;

        /* Notifiy anyone listening that this transfer took place */
        emit Transfer(msg.sender, _to, _value);
    }
	
	function transferFrom(address _from, address _to, uint256 _value) public {
	    if (balanceOf[_from] < _value) return;    
	    if (balanceOf[_to] + _value < balanceOf[_to]) return;
	    
	    balanceOf[_from] -= _value;
	    balanceOf[_to] += _value;
	    /* Notifiy anyone listening that this transfer took place */
        emit Transfer(_from, _to, _value);
	}
  
	
}