import { ethers } from "ethers";
import dotenv from "dotenv"
dotenv.config({path: "../.env"});
import { eventAbi } from "../artifacts/eventAbi";
import { getProviderURLs } from "./getProviderUrl";

export const getBurnConfirmation = async (chain:string, startBlock:string, userAddress:string) => {
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

    var filter = contract.filters.Burned(null, null, null);

    console.log(filter);

    data2 = await contract.queryFilter(
        filter,
        Number(startBlock),
        parseInt(startBlock)+20
    );

    // Decode event data
    data2 = data2.map((el) => {
        const decodedEvent = contract.interface.decodeEventLog("Burned", el.data, el.topics);
        return {
            booster: decodedEvent[0],
            user: decodedEvent[1],
            tokenId: decodedEvent[2],
            currentBlock: latestBlockNumber,
            transactionBlock: el.blockNumber,
        };
    });

    console.log("data 2: ",data2)

    for (let i = 0; i < data2.length; i++) {
    if (data2[i].user == userAddress) {
        console.log("true");
        const tokenID = data2[i].tokenId
        return { boosterMinted: data2[i].booster, tokenId: tokenID.toString()};
    }
    }
    
    return { boosterMinted: "", tokenId: 0};

  } catch (error) {
    console.error("Error in getLastEvent:", error);
    throw error;
  }
};
