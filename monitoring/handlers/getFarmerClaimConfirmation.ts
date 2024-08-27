import { ethers } from "ethers";
import dotenv from "dotenv"
dotenv.config({path: "../.env"});
import { getProviderURLs } from "./getProviderUrl";
import { nftFaucetAbi } from "../artifacts/nftFaucet";

export const getFarmerClaimConfirmation = async (chain:string, userAddress:string) => {
  const curUrl = await getProviderURLs(chain);
  console.log(curUrl);

  try {

    const provider = new ethers.JsonRpcProvider(curUrl);
    console.log("provider", provider);
    const contract = new ethers.Contract(
      "0x16c2c9B38F59Ca3749D29a428d426558619c85E8",
      nftFaucetAbi,
      provider
    );

    const latestBlockNumber = await provider.getBlockNumber();

    console.log("latest block", latestBlockNumber);

    const claimed = await contract.farmerClaimedBy(userAddress);

    return claimed;
   

  } catch (error) {
    console.error("Error in getLastEvent:", error);
    throw error;
  }
};
