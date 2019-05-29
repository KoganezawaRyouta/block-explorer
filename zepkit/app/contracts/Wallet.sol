pragma solidity ^0.5.0;

import "openzeppelin-eth/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-eth/contracts/token/ERC20/StandaloneERC20.sol";
import "openzeppelin-eth/contracts/ownership/Ownable.sol";
import "zos-lib/contracts/Initializable.sol";
import "./IMff.sol";

contract Wallet is Ownable {

  function transferERC20(IERC20 token, address to, uint256 amount) public onlyOwner returns (bool) {
    require(token.transfer(to, amount));
  }

  function transferERC201(StandaloneERC20 token, address to, uint256 amount) public onlyOwner returns (bool) {
    token.transfer(to, amount);
    return true;
  }

  function transferERC2012(IERC20 tokena, IERC20 tokenb, uint256 amount, address to) public onlyOwner returns (bool) {
    tokena.transfer(to, amount);
    tokenb.transfer(to, amount);
    return true;
  }

  function kogane(IMff tokena, IMff tokenb, uint256 amount, address to) public onlyOwner returns (bool) {
    tokena.teee(to, amount);
    tokenb.teee(to, amount);
    return true;
  }

  function koganee(ERC20 tokena, ERC20 tokenb, uint256 amount, address to) public onlyOwner returns (bool) {
    tokena.transfer(to, amount);
    tokenb.transfer(to, amount);
    return true;
  }
}
