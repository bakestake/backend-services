import ethers from 'ethers';
import { getProviderURLs } from './getProviderUrl';
import {faucetABI} from "../artifacts/faucet"

export const getClaimTS = async (chain : string, userAddress:string) =>{
    try{
        const contractInst = new ethers.Contract(
            "0xB2A338Fb022365Aa40a2c7ADA3Bbf1Ae001D6dbe", 
            faucetABI,
            new ethers.Wallet(process.env.PRIVATE_KEY || "", new ethers.JsonRpcProvider(await getProviderURLs(chain) || ""))
        )

        const ts =  await contractInst.nextClaimTimeInSeconds(userAddress);

        return ts;

    }catch(error){
        console.log(error)
    }

}