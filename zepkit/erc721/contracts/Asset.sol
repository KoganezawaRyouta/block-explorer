pragma solidity ^0.5.0;

// 1. ERC721Full.solのインポート
import "openzeppelin-solidity/contracts/token/ERC721/ERC721Full.sol";

// 2. ERC721Fullの継承
// 1
// 2
// 3
contract Asset is ERC721Full {
  constructor(string memory name, string memory symbol, uint tokenId, string memory tokenURI)
    ERC721Full(name, symbol) public {
      _mint(msg.sender, tokenId);
      _setTokenURI(tokenId, tokenURI);
    }

  function mint(address to, uint256 tokenId) public {
    _mint(to, tokenId);
  }
}
