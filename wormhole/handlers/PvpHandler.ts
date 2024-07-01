import {
  EthCallData,
  EthCallQueryRequest,
  PerChainQueryRequest,
  QueryProxyQueryResponse,
  QueryRequest,
} from "@wormhole-foundation/wormhole-query-sdk";
import axios from "axios";
import { getProviderURLs } from "./getProviderUrl";
import dotenv from "dotenv";
import { ethers } from "ethers";
import { ABI } from "../artifacts/Escrow";

dotenv.config({path: "../.env"});

type chainConf = {
  chains:string,
  chainId:number,
  rpc:string
}

const getChainConf = async (chain:string) => {
  switch (chain) {
    case "fuji":
      return { chains: "fuji", chainId: 6, rpc: getProviderURLs("fuji") || "" };
    case "arbSepolia":
     return { chains: "arbSepolia", chainId: 10003, rpc: getProviderURLs("arbSepolia") || "" };
    case "amoy":
      return { chains: "amoy", chainId: 10007, rpc: getProviderURLs("amoy") || "" };
    case "bscTestnet":
      return { chains: "bscTestnet", chainId: 4, rpc: getProviderURLs("bscTestnet") || "" };
    case 'beraTestnet':
      return { chains: "beraTestnet", chainId: 39, rpc: getProviderURLs("beraTestnet") || "" }
    case 'coreTestnet':
      return  { chains: "coreTestnet", chainId: 4, rpc: getProviderURLs("coreTestnet") || "" }
    case 'baseSepolia':
      return { chains: "baseSepolia", chainId: 10004, rpc: getProviderURLs("baseSepolia") || "" }
    default:
      throw new Error("Not configuered");
  }
}

const PVP = async (curNetwork: string) => {
  try {
    const contractAddress = "0x9e014aAE147D85d5764641e773dE9C29aC0141e9";
    const selector = "0x0ea901d2";
    const chain : chainConf = await getChainConf(curNetwork);
    const chains = [chain];

    console.log("Eth calls and block number calls getting recorded");

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
                  params: [{ to: contractAddress, data: selector }, "latest"],
                },
              ])
              .catch((error) => {
                console.error(`Error fetching data for rpc: ${rpc}`, error);
                return null;
              })
          : Promise.reject(new Error(`RPC URL is undefined for chain ${chainId}`)),
      ),
    );

    console.log("Preparing eth call data");

    const callData: EthCallData = {
      to: contractAddress,
      data: selector,
    };

    console.log("Preparing queries for all chains");

    let perChainQueries = chains.map(({ chainId }, idx) => {
      if (!responses[idx] || !responses[idx]?.data) {
        console.error(`no response data for chain ID: ${chainId}`);
        throw new Error(`no response data for chain ID: ${chainId}`);
      }
      return new PerChainQueryRequest(
        chainId,
        new EthCallQueryRequest(responses[idx]?.data?.[0]?.result?.number, [callData]),
      );
    });

    const nonce = 2;
    const request = new QueryRequest(nonce, perChainQueries);
    const serialized = request.serialize();

    console.log("Querying cross chain");

    const response = await axios
      .put<QueryProxyQueryResponse>(
        "https://testnet.query.wormhole.com/v1/query",
        {
          bytes: Buffer.from(serialized).toString("hex"),
        },
        { headers: { "X-API-Key": process.env.WORMHOLE_API_KEY } },
      )
      .catch((error) => {
        console.error("error querying cross chain", error);
        throw error;
      });

  } catch (err) {
    console.error("an error occurred during the cross-chain query process", err);
  }
};

export default PVP;
