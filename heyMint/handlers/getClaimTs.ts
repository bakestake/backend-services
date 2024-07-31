import ethers from 'ethers';
import { getProviderURLs } from './getProviderUrl';
import {faucetABI} from "../artifacts/faucet"

export const getClaimTS = async (chain : string, userAddress:string) =>{
    try{
        const contractInst = new ethers.Contract(
            "0x26705ad938791e61aa64a2a9d808378805aec819", 
            faucetABI,
            new ethers.Wallet(process.env.PRIVATE_KEY || "", new ethers.JsonRpcProvider(await getProviderURLs(chain) || ""))
        )

        const ts =  await contractInst.nextClaimTimeInSeconds(userAddress);

        return ts;

    }catch(error){
        console.log(error)
    }

}