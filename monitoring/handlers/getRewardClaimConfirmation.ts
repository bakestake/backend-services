import { ethers } from "ethers";
import dotenv from "dotenv"
dotenv.config({path: "../.env"});
import { eventAbi } from "../artifacts/eventAbi";
import { getProviderURLs } from "./getProviderUrl";

export const getLastRewardClaimEvent = async (chain:string, startBlock:string, userAddress:string) => {
  const curUrl = await getProviderURLs(chain);
  console.log(curUrl);

  try {

    const provider = new ethers.JsonRpcProvider(curUrl);
    console.log("provider", provider);
    const contract = new ethers.Contract(
      "0xB2A338Fb022365Aa40a2c7ADA3Bbf1Ae001D6dbe",
      eventAbi,
      provider
    );

    const latestBlockNumber = await provider.getBlockNumber();

    console.log("latest block", latestBlockNumber);

    var data2 = [];

    var filter = contract.filters.rewardsClaimed(null, null);

    console.log(filter);

    data2 = await contract.queryFilter(
        filter,
        Number(startBlock),
        latestBlockNumber
    );

    // Decode event data
    data2 = data2.map((el) => {
        const decodedEvent = contract.interface.decodeEventLog("rewardsClaimed", el.data, el.topics);
        return {
            user: decodedEvent[0],
            amount: decodedEvent[1],
            currentBlock: latestBlockNumber,
            transactionBlock: el.blockNumber,
        };
    });

    console.log("data 2: ",data2)

    for (let i = 0; i < data2.length; i++) {
    if (data2[i].user == userAddress) {
        console.log("true");
        const amount = data2[i].amount
        return { eventOccurred: true, amount: amount.toString()};
    }
    }
    
    console.log("false");
    return {eventOccurred: false, amount: 0};

  } catch (error) {
    console.error("Error in getLastEvent:", error);
    throw error;
  }
};
