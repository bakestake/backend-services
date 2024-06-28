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
const StateUpdate_1 = require("../artifacts/StateUpdate");
dotenv_1.default.config({ path: "../../.env" });
const CCQ = (curNetwork) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const contractAddress = "0x9e014aAE147D85d5764641e773dE9C29aC0141e9";
        const selector = "0x4269e94c";
        const chains = [
            { chains: "fuji", chainId: 6, rpc: (0, getProviderUrl_1.getProviderURLs)("fuji") },
            { chains: "arbSepolia", chainId: 10003, rpc: (0, getProviderUrl_1.getProviderURLs)("arbSepolia") },
            { chains: "amoy", chainId: 10007, rpc: (0, getProviderUrl_1.getProviderURLs)("amoy") },
            { chains: "bscTestnet", chainId: 4, rpc: (0, getProviderUrl_1.getProviderURLs)("bscTestnet") }
        ];
        // First step is getting responses from all chains
        console.log("Eth calls and block number calls getting recorded");
        const responses = yield Promise.all(chains.map(({ rpc }) => rpc
            ? axios_1.default.post(rpc, [
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
            : Promise.reject()));
        // second step is construct per chain query request
        console.log("preping eth call data");
        const callData = {
            to: contractAddress,
            data: selector,
        };
        console.log("Preping queries for all chains");
        let perChainQueries = chains.map(({ chainId }, idx) => {
            var _a, _b, _c, _d;
            return new wormhole_query_sdk_1.PerChainQueryRequest(chainId, new wormhole_query_sdk_1.EthCallQueryRequest((_d = (_c = (_b = (_a = responses[idx]) === null || _a === void 0 ? void 0 : _a.data) === null || _b === void 0 ? void 0 : _b[0]) === null || _c === void 0 ? void 0 : _c.result) === null || _d === void 0 ? void 0 : _d.number, [callData]));
        });
        const nonce = 2;
        const request = new wormhole_query_sdk_1.QueryRequest(nonce, perChainQueries);
        const serialized = request.serialize();
        console.log("querying cross chain");
        const response = yield axios_1.default.put("https://testnet.query.wormhole.com/v1/query", {
            bytes: Buffer.from(serialized).toString("hex"),
        }, { headers: { "X-API-Key": process.env.WORMHOLE_API_KEY } });
        console.log("broadcasting to chain");
        const contract = new ethers_1.ethers.Contract(contractAddress, StateUpdate_1.ABI, new ethers_1.ethers.Wallet(process.env.PRIVATE_KEY || "", new ethers_1.ethers.JsonRpcProvider((0, getProviderUrl_1.getProviderURLs)(curNetwork))));
        const tx = yield contract.updateState(`0x${response.data.bytes}`, response.data.signatures.map((s) => ({
            r: `0x${s.substring(0, 64)}`,
            s: `0x${s.substring(64, 128)}`,
            v: `0x${(parseInt(s.substring(128, 130), 16) + 27).toString(16)}`,
            guardianIndex: `0x${s.substring(130, 132)}`,
        })));
        yield tx.wait();
    }
    catch (err) {
        console.log(err);
    }
});
