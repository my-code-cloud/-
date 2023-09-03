pragma solidity ^0.4.15;

contract SimpleSuicide {

  function sudicideAnyone() {
    selfdestruct(msg.sender);
  }

}