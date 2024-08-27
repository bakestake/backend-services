import { ethers } from "ethers";
import dotenv from "dotenv"
dotenv.config({path: "../.env"});
import { getProviderURLs } from "./getProviderUrl";
import { faucetABI } from "../artifacts/faucet";

export const getBudsClaimConfirmation = async (chain:string, userAddress:string) => {
  const curUrl = await getProviderURLs(chain);
  console.log(curUrl);

  try {

    const provider = new ethers.JsonRpcProvider(curUrl);
    console.log("provider", provider);
    const contract = new ethers.Contract(
      "0xc01bC8B25245FfcA5950DDA2bC1e2e7B3279F822",
      faucetABI,
      provider
    );

    const latestBlockNumber = await provider.getBlockNumber();

    console.log("latest block", latestBlockNumber);

    const claimTs = await contract.nextClaimTimeInSeconds(userAddress);

    if(claimTs != 0){
        return true;
    }

    return false;

  } catch (error) {
    console.error("Error in getLastEvent:", error);
    throw error;
  }
};
