// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

contract DeveloperList {
    bool public initialized;
    bool public enabled;
    address public admin;
    address public pendingAdmin;
    mapping(address => bool) private developers;

    // events
    event AdminChanging(address indexed newAdmin);
    event AdminChanged(address indexed newAdmin);
    event DeveloperAdded(address indexed addr);
    event DeveloperRemoved(address indexed addr);
    event EnableStateChanged(bool indexed newState);

    modifier onlyNotInitialized() {
        require(!initialized, "Already initialized");
        _;
    }

    modifier onlyAdmin() {
        require(msg.sender == admin, "Admin only");
        _;
    }

    function initialize(address _admin) external onlyNotInitialized {
        admin = _admin;
        initialized = true;
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

    function commitChangeAdmin(address newAdmin) external onlyAdmin {
        pendingAdmin = newAdmin;

        emit AdminChanging(newAdmin);
    }

    function confirmChangeAdmin() external {
        require(msg.sender == pendingAdmin, "New admin only");
        admin = pendingAdmin;
        pendingAdmin = address(0);

        emit AdminChanged(admin);
    }

    function addDeveloper(address addr) external onlyAdmin {
        require(!developers[addr], "Already added");
        developers[addr] = true;

        emit DeveloperAdded(addr);
    }

    function removeDeveloper(address addr) external onlyAdmin {
        require(developers[addr], "Not a developer");
        developers[addr] = false;

        emit DeveloperRemoved(addr);
    }

    function isDeveloper(address addr) external view returns (bool) {
        return developers[addr];
    }
}
