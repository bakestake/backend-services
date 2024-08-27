import { ethers } from "ethers";
import dotenv from "dotenv"
dotenv.config({path: "../.env"});
import { eventAbi } from "../artifacts/eventAbi";
import { getProviderURLs } from "./getProviderUrl";

export const getLastUnstakeEvent = async (chain:string, startBlock:string, userAddress:string) => {
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

    var filter = contract.filters.UnStaked(userAddress, null, null, null, null, null);

    console.log("filter", filter);

    data2 = await contract.queryFilter(
      filter,
      Number(startBlock),
      latestBlockNumber
    );

    console.log("data 2 filter :", data2);

    data2 = data2.map((el) => {
      const decodedEvent = contract.interface.decodeEventLog("UnStaked", el.data, el.topics);
          return {
              sender: decodedEvent[0],
              tokenId: Number(decodedEvent[1].toString()),
              amount: Number(decodedEvent[2].toString()) / 1e18,
              currentBlock: latestBlockNumber,
              transactionBlock: el.blockNumber,
          };
      });

    console.log("data 2", data2);
    
    for (let i = 0; i < data2.length; i++) {
      if (data2[i].sender == userAddress) {
        console.log("data[i] in events", data2[i]);
        console.log("true");
        return { eventOccurred: true, amount: data2[i].amount, tokenId: data2[i].tokenId };
      }
    }

    console.log("false");
    return { eventOccurred: false, amount: 0, tokenId:0};

  } catch (error) {
    console.error("Error in getLastEvent:", error);
    throw error;
  }
};
