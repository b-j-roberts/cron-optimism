import { deploy } from './deploy-deps.js'
import fs from 'fs'

(async () => {
  try {
      console.log("Deploying Cron Contract onto l2 at", "http://localhost:5055, from address", process.env.FROM_ADDRESS)
     // const result =
      await deploy('CronContract', [], "http://localhost:5055", '0x' + process.env.FROM_ADDRESS)
      //console.log(result)
      //console.log("Deployed Cron contract to : ", result.address)
      //var jsonOutput = "{\"address\": \"" + result.address + "\"}"
      // Write the contract address to a file
      //fs.writeFileSync('./builds/cron-contract.txt', jsonOutput)
  } catch (e) {
      console.log(e.message)
  }

  process.exit(0)
})()

