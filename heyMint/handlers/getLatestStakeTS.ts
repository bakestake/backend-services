import { ethers } from 'ethers';
import { getProviderURLs } from './getProviderUrl';
import { getterSetter } from '../artifacts/getterSetter';

export const getLatestStakeTS = async (chain: string, userAddress: string): Promise<number | null> => {
  try {
    const providerUrl = await getProviderURLs(chain);
    if (!providerUrl) {
      throw new Error(`Unable to get provider URL for chain: ${chain}`);
    }

    const provider = new ethers.JsonRpcProvider(providerUrl);
    const wallet = new ethers.Wallet(process.env.PRIVATE_KEY || "", provider);

    const contractInst = new ethers.Contract(
      "0x26705ad938791e61aa64a2a9d808378805aec819",
      getterSetter,
      wallet
    );

    const stakes = await contractInst.getUserStakes(userAddress);

    if(stakes.length == 0){
        return 0;
    }

    return Number(stakes[stakes.length-1].timeStamp); // Convert BigInt to Number
  } catch (error) {
    console.error('Error in getUserStake:', error);
    return null; // or throw error if you want to handle it higher up
  }
};
