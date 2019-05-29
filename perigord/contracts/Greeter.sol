pragma solidity ^0.5.0;

contract Greeter {
    string public greeting;

    event Result(address from, string stored);

    function setGreeting(string memory _greeting) public {
        greeting = _greeting;
        emit Result(msg.sender, greeting);
    }

    function greet() public view returns (string memory) {
        return greeting;
    }
}