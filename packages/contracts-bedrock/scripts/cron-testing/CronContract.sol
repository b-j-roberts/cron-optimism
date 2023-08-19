pragma solidity >=0.8.2 <0.9.0;

interface IJob {
    function executeCron() external;
}

contract CronContract is IJob {
    uint256 public value;

    constructor() {
        value = 0;
    }

    function executeCron() external override {
        value = value + 1;
    }

    function getCronContractValue() external view returns (uint256) {
        return value;
    }
}
