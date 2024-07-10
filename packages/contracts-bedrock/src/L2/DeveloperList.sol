// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

contract DeveloperList {
    address public immutable admin;

    bool public enabled;
    mapping(address => bool) internal _developers;

    // events
    event DeveloperAdded(address indexed addr);
    event DeveloperRemoved(address indexed addr);
    event EnableStateChanged(bool indexed newState);

    modifier onlyAdmin() {
        require(msg.sender == admin, "Admin only");
        _;
    }

    constructor(address _admin) {
        admin = _admin;
    }

    function enableDevVerify() external onlyAdmin {
        require(enabled == false, "Already enabled");
        enabled = true;

        emit EnableStateChanged(true);
    }

    function disableDevVerify() external onlyAdmin {
        require(enabled, "Already disabled");
        enabled = false;

        emit EnableStateChanged(false);
    }

    function addDeveloper(address addr) external onlyAdmin {
        require(!_developers[addr], "Already added");
        _developers[addr] = true;

        emit DeveloperAdded(addr);
    }

    function removeDeveloper(address addr) external onlyAdmin {
        require(_developers[addr], "Not a developer");
        _developers[addr] = false;

        emit DeveloperRemoved(addr);
    }

    function isDeveloper(address addr) external view returns (bool) {
        return _developers[addr];
    }
}
