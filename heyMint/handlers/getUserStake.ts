import { ethers } from 'ethers';
import { getProviderURLs } from './getProviderUrl';
import { getterSetter } from '../artifacts/getterSetter';

export const getUserStake = async (chain: string, userAddress: string): Promise<number | null> => {
  try {
    const providerUrl = await getProviderURLs(chain);
    if (!providerUrl) {
      throw new Error(`Unable to get provider URL for chain: ${chain}`);
    }

    const provider = new ethers.JsonRpcProvider(providerUrl);
    const wallet = new ethers.Wallet(process.env.PRIVATE_KEY || "", provider);

    const contractInst = new ethers.Contract(
      "0xB2A338Fb022365Aa40a2c7ADA3Bbf1Ae001D6dbe",
      getterSetter,
      wallet
    );

    const stakes = await contractInst.getUserStakes(userAddress);
    let stakedAmount = BigInt(0);

    for (let i = 0; i < stakes[0].length; i++) {
      stakedAmount += BigInt(stakes[0][i].budsAmount);
    }

    return Number(stakedAmount)/1e18; // Convert BigInt to Number
  } catch (error) {
    console.error('Error in getUserStake:', error);
    return null; // or throw error if you want to handle it higher up
  }
};
