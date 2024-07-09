// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract OPEN is ERC20 {
    constructor(address initialAccount) ERC20("OPEN", "OPEN") {
        _mint(initialAccount, 1000 * 10 ** 8 * 10 ** 18);
    }
}
