"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const wormhole_query_sdk_1 = require("@wormhole-foundation/wormhole-query-sdk");
const axios_1 = __importDefault(require("axios"));
const getProviderUrl_1 = require("./getProviderUrl");
const dotenv_1 = __importDefault(require("dotenv"));
const ethers_1 = require("ethers");
const beraStateUpdate_1 = require("./beraStateUpdate");
dotenv_1.default.config({ path: "../../.env" });
const CCQ = (chain) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        if (chain == "beraTestnet") {
            yield (0, beraStateUpdate_1.BeraStateUpdate)();
        }
        else {
            yield stateUpdate(chain);
        }
    }
    catch (error) {
        console.log(error);
    }
});
const stateUpdate = (chain) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const contractAddress = "0xB2A338Fb022365Aa40a2c7ADA3Bbf1Ae001D6dbe";
        const selector = "0x4269e94c";
        const chains = [
            { chains: "fuji", chainId: 6, rpc: (0, getProviderUrl_1.getProviderURLs)("fuji") },
            { chains: "arbSepolia", chainId: 10003, rpc: (0, getProviderUrl_1.getProviderURLs)("arbSepolia") },
            { chains: "amoy", chainId: 10007, rpc: (0, getProviderUrl_1.getProviderURLs)("amoy") },
            { chains: "bscTestnet", chainId: 4, rpc: (0, getProviderUrl_1.getProviderURLs)("bscTestnet") },
            // { chains: "beraTestnet", chainId: 39, rpc: getProviderURLs("beraTestnet") },
            // { chains: "coreTestnet", chainId: 4, rpc: getProviderURLs("coreTestnet") },
            { chains: "baseSepolia", chainId: 10004, rpc: (0, getProviderUrl_1.getProviderURLs)("baseSepolia") },
        ];
        console.log("Eth calls and block number calls getting recorded");
        const responses = yield Promise.all(chains.map(({ rpc, chainId }) => rpc
            ? axios_1.default
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
            : Promise.reject(new Error(`RPC URL is undefined for chain ${chainId}`))));
        console.log("Preparing eth call data");
        const callData = {
            to: contractAddress,
            data: selector,
        };
        console.log("Preparing queries for all chains");
        let perChainQueries = chains.map(({ chainId }, idx) => {
            var _a, _b, _c, _d, _e;
            if (!responses[idx] || !((_a = responses[idx]) === null || _a === void 0 ? void 0 : _a.data)) {
                console.error(`no response data for chain ID: ${chainId}`);
                throw new Error(`no response data for chain ID: ${chainId}`);
            }
            return new wormhole_query_sdk_1.PerChainQueryRequest(chainId, new wormhole_query_sdk_1.EthCallQueryRequest((_e = (_d = (_c = (_b = responses[idx]) === null || _b === void 0 ? void 0 : _b.data) === null || _c === void 0 ? void 0 : _c[0]) === null || _d === void 0 ? void 0 : _d.result) === null || _e === void 0 ? void 0 : _e.number, [callData]));
        });
        const nonce = 2;
        const request = new wormhole_query_sdk_1.QueryRequest(nonce, perChainQueries);
        const serialized = request.serialize();
        console.log("Querying cross chain");
        const response = yield axios_1.default
            .put("https://testnet.query.wormhole.com/v1/query", {
            bytes: Buffer.from(serialized).toString("hex"),
        }, { headers: { "X-API-Key": process.env.WORMHOLE_API_KEY } })
            .catch((error) => {
            console.error("error querying cross chain", error);
            throw error;
        });
        const mock = new wormhole_query_sdk_1.QueryProxyMock({ 6: chains[0].rpc || "", 10003: chains[1].rpc || "", 10007: chains[2].rpc || "", 4: chains[3].rpc || "", 10004: chains[4].rpc || "" });
        const mockData = yield mock.mock(request);
        const mockQueryResponse = wormhole_query_sdk_1.QueryResponse.from(mockData.bytes);
        let global = 0;
        for (let i = 0; i < responses.length; i++) {
            global += parseInt(mockQueryResponse.responses[i].response.results[0]);
        }
        const mockQueryResult = mockQueryResponse.responses[0].response.results[0];
        console.log(`Mock Query Result: ${mockQueryResult} (${BigInt(mockQueryResult)})`);
        const bytes = `0x${response.data.bytes}`;
        const signatures = response.data.signatures.map((s) => ({
            r: `0x${s.substring(0, 64)}`,
            s: `0x${s.substring(64, 128)}`,
            v: `0x${(parseInt(s.substring(128, 130), 16) + 27).toString(16)}`,
            guardianIndex: `0x${s.substring(130, 132)}`,
        }));
        console.log(global / 1e18, typeof (global));
        let currentState = 0;
        const contractInst = new ethers_1.ethers.Contract("0xB2A338Fb022365Aa40a2c7ADA3Bbf1Ae001D6dbe", [{
                "inputs": [],
                "name": "getGlobalStakedBuds",
                "outputs": [
                    {
                        "internalType": "uint256",
                        "name": "",
                        "type": "uint256"
                    }
                ],
                "stateMutability": "view",
                "type": "function"
            }], new ethers_1.ethers.Wallet(process.env.PRIVATE_KEY || "", new ethers_1.ethers.JsonRpcProvider((yield (0, getProviderUrl_1.getProviderURLs)(chain)) || "")));
        currentState = yield contractInst.getGlobalStakedBuds();
        const parsedState = parseInt(currentState.toString()) / 1e18;
        global = global / 1e18;
        console.log(parsedState, typeof (parsedState));
        if (parsedState != global) {
            console.log("State update needed");
            const stateContractInst = new ethers_1.ethers.Contract("0xB2A338Fb022365Aa40a2c7ADA3Bbf1Ae001D6dbe", [{
                    "inputs": [
                        {
                            "internalType": "bytes",
                            "name": "response",
                            "type": "bytes"
                        },
                        {
                            "components": [
                                {
                                    "internalType": "bytes32",
                                    "name": "r",
                                    "type": "bytes32"
                                },
                                {
                                    "internalType": "bytes32",
                                    "name": "s",
                                    "type": "bytes32"
                                },
                                {
                                    "internalType": "uint8",
                                    "name": "v",
                                    "type": "uint8"
                                },
                                {
                                    "internalType": "uint8",
                                    "name": "guardianIndex",
                                    "type": "uint8"
                                }
                            ],
                            "internalType": "struct IWormhole.Signature[]",
                            "name": "signatures",
                            "type": "tuple[]"
                        }
                    ],
                    "name": "updateState",
                    "outputs": [],
                    "stateMutability": "nonpayable",
                    "type": "function"
                }], new ethers_1.ethers.Wallet(process.env.PRIVATE_KEY || "", new ethers_1.ethers.JsonRpcProvider((0, getProviderUrl_1.getProviderURLs)(chain) || "")));
            yield stateContractInst.updateState(bytes, signatures);
            console.log("State updated");
        }
        else {
            console.log("State update not needed");
        }
    }
    catch (Error) {
        console.error("an error occurred during the cross-chain query process", Error);
    }
});
exports.default = CCQ;
