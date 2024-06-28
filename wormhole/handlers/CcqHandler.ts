import {
  EthCallData,
  EthCallQueryRequest,
  PerChainQueryRequest,
  QueryProxyQueryResponse,
  QueryRequest,
  QueryResponse,
  signaturesToEvmStruct,
} from "@wormhole-foundation/wormhole-query-sdk";
import axios from "axios";
import { getProviderURLs } from "./getProviderUrl";
import dotenv from "dotenv";
import { ethers } from "ethers";
import { ABI } from "../artifacts/StateUpdate";

dotenv.config({path:"../../.env"});

const CCQ = async (curNetwork : string) => {
  try {
    const contractAddress = "0x9e014aAE147D85d5764641e773dE9C29aC0141e9";
    const selector = "0x4269e94c";
    const chains = [
      {chains : "fuji", chainId : 6, rpc: getProviderURLs("fuji")}, 
      {chains : "arbSepolia", chainId : 10003, rpc: getProviderURLs("arbSepolia")}, 
      {chains : "amoy", chainId : 10007, rpc: getProviderURLs("amoy")}, 
      {chains : "bscTestnet", chainId : 4, rpc: getProviderURLs("bscTestnet")}
    ];

    // First step is getting responses from all chains

    console.log("Eth calls and block number calls getting recorded")

    const responses = await Promise.all(
      chains.map(({ rpc }) =>
        rpc
          ? axios.post(rpc, [
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
                params: [{ to: contractAddress, data: selector }, "latest"],
              },
            ])
          : Promise.reject(),
      ),
    );

    // second step is construct per chain query request

    console.log("preping eth call data")

    const callData: EthCallData = {
      to: contractAddress,
      data: selector,
    }; 

    console.log("Preping queries for all chains")

    let perChainQueries = chains.map( ({chainId}, idx) =>
        new PerChainQueryRequest(
          chainId,
          new EthCallQueryRequest(responses[idx]?.data?.[0]?.result?.number, [callData])
        ),

    );

    const nonce = 2;
    const request = new QueryRequest(nonce, perChainQueries);
    const serialized = request.serialize();

    console.log("querying cross chain")

    const response = await axios.put<QueryProxyQueryResponse>(
      "https://testnet.query.wormhole.com/v1/query",
      {
        bytes: Buffer.from(serialized).toString("hex"),
      },
      { headers: { "X-API-Key": process.env.WORMHOLE_API_KEY} },
    );

    console.log("broadcasting to chain")

    
    const contract = new ethers.Contract(
      contractAddress, 
      ABI, 
      new ethers.Wallet(
        process.env.PRIVATE_KEY || "", 
        new ethers.JsonRpcProvider(getProviderURLs(curNetwork))
      ) 
    )

    const tx = await contract.updateState(
      `0x${response.data.bytes}`,
      response.data.signatures.map((s) => ({
        r: `0x${s.substring(0, 64)}`,
        s: `0x${s.substring(64, 128)}`,
        v: `0x${(parseInt(s.substring(128, 130), 16) + 27).toString(16)}`,
        guardianIndex: `0x${s.substring(130, 132)}`,
      })),
    );

    await tx.wait()

  }catch(err){
    console.log(err);
  }
}

