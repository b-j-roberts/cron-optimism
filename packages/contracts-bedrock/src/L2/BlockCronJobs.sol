pragma solidity 0.8.15;

import { Semver } from "../universal/Semver.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";

interface IJob {
    function executeCron() external;
}

//TODO: Do a seperate time based cron
// Predeploy Contract : 0x42000000000000000000000000000000000000c0
contract BlockCronJobs is Semver, Ownable, IJob {
    address public constant DEPOSITOR_ACCOUNT = 0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001;

    uint256 public MINIMUM_BLOCK_INTERVAL = 10;

    address[] public jobs;
    uint256[] public lastExecutedAt;
    uint256[] public jobIntervals;

    constructor(address _owner) Ownable() Semver(0, 0, 1) {
        transferOwnership(_owner);
    }

    function setMinimumBlockInterval(uint256 _minimumBlockInterval) external onlyOwner {
        MINIMUM_BLOCK_INTERVAL = _minimumBlockInterval;
    }

    function addJob(address _job, uint256 _interval) external onlyOwner {
        require(_job != address(0), "Cronjobs: Job address cannot be 0");
        require(_interval >= MINIMUM_BLOCK_INTERVAL, "Cronjobs: Interval must be greater than minimum interval");

        jobs.push(_job);
        jobIntervals.push(_interval);
        lastExecutedAt.push(block.number);
    }

    function removeJob(address _job) external onlyOwner {
        require(_job != address(0), "Cronjobs: Job address cannot be 0");

        for (uint256 i = 0; i < jobs.length; i++) {
            if (jobs[i] == _job) {
                // Overwrite the job with the last job in the array & pop the last job away
                // Note : This means the order of jobs will change
                jobs[i] = jobs[jobs.length - 1];
                jobs.pop();

                lastExecutedAt[i] = lastExecutedAt[lastExecutedAt.length - 1];
                lastExecutedAt.pop();

                jobIntervals[i] = jobIntervals[jobIntervals.length - 1];
                jobIntervals.pop();
                break;
            }
        }
    }

    function executeCron() external override {
        //uint256 gasStart = gasleft();
        require(msg.sender == DEPOSITOR_ACCOUNT, "Cronjobs: Only depositor account can executeCron cron");

        for (uint256 i = 0; i < jobs.length; i++) {
            address job = jobs[i];
            uint256 interval = jobIntervals[i];
            uint256 lastExecuted = lastExecutedAt[i];

            //TODO: Stop and save position for next block if full on gas? or allow pageing & client controlled usage?
            // https://ethereum.stackexchange.com/questions/134968/is-possible-to-control-gasleft-and-exit-of-a-loop-before-out-of-gas-return
            if (block.number >= lastExecuted + interval) {
                IJob(job).executeCron();
                lastExecutedAt[i] = block.number;
            }
        }
    }

    function getJobs() external view returns (address[] memory) {
        return jobs;
    }
}
