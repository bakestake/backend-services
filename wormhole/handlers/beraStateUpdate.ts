import { ethers } from "ethers"
import { getProviderURLs } from "./getProviderUrl"
import {ABI} from "../artifacts/setter"
import axios from "axios";
import dotenv from "dotenv";

dotenv.config({path: "../../.env"});

export const BeraStateUpdate = async() => {
    try{
        console.log("updating bera")

        const chains = [
        { chains: "fuji", chainId: 6, rpc: getProviderURLs("fuji") },
        { chains: "arbSepolia", chainId: 10003, rpc: getProviderURLs("arbSepolia") },
        { chains: "amoy", chainId: 10007, rpc: getProviderURLs("amoy") },
        { chains: "bscTestnet", chainId: 4, rpc: getProviderURLs("bscTestnet") },
        { chains: "beraTestnet", chainId: 39, rpc: getProviderURLs("beraTestnet") },
        // { chains: "coreTestnet", chainId: 4, rpc: getProviderURLs("coreTestnet") },
        { chains: "baseSepolia", chainId: 10004, rpc: getProviderURLs("baseSepolia") },
        ];

        console.log("getting responses from all chains")

        const responses = await Promise.all(
            chains.map(({ rpc, chainId }) =>
            rpc
                ? axios
                    .post(rpc, [
                    {
                        jsonrpc: "2.0",
                        id: 1,
                        method: "eth_getBlockByNumber",
                        params: ["latest", false],
                    },
                    {
                        jsonrpc: "2.0",
                        id: 2,
                        method: "eth_call",
                        params: [{ to: "0x26705aD938791e61Aa64a2a9D808378805aec819", data: "0x4269e94c" }, "latest"],
                    },
                    ])
                    .catch((error) => {
                    console.error(`Error fetching data for rpc: ${rpc}`, error);
                    return null;
                    })
                : Promise.reject(new Error(`RPC URL is undefined for chain ${chainId}`)),
            ),
        );

        console.log("summing up global")

        let globalBudsCount = 0;

        for(let i = 0; i < responses.length; i++){
            console.log(parseInt(responses[i]?.data?.[0]?.result?.number))
            globalBudsCount += parseInt(responses[i]?.data?.[0]?.result?.number);
        }

        console.log(globalBudsCount)

        console.log("updating bera state")

        const setterInst = new ethers.Contract(
            "0x26705ad938791e61aa64a2a9d808378805aec819", 
            ABI,
            new ethers.Wallet(process.env.STATE_PRIVATE_KEY || "", new ethers.JsonRpcProvider(getProviderURLs("beraTestnet") || ""))
        );

        console.log("doing tx")

        const tx = await setterInst.setGlobalBudsState(globalBudsCount);

        await tx.wait();

    }catch(error){
        console.log(error);
    }
    
}