// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract FileRegistry {
    mapping(string => string) private files;

    function save(string memory filePath, string memory cid) public {
        files[filePath] = cid;
    }

    function get(string memory filePath) public view returns (string memory) {
        return files[filePath];
    }
}
