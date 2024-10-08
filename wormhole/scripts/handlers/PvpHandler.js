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
dotenv_1.default.config({ path: "../.env" });
const getChainConf = (chain) => __awaiter(void 0, void 0, void 0, function* () {
    switch (chain) {
        case "fuji":
            return { chains: "fuji", chainId: 6, rpc: (0, getProviderUrl_1.getProviderURLs)("fuji") || "" };
        case "arbSepolia":
            return { chains: "arbSepolia", chainId: 10003, rpc: (0, getProviderUrl_1.getProviderURLs)("arbSepolia") || "" };
        case "amoy":
            return { chains: "amoy", chainId: 10007, rpc: (0, getProviderUrl_1.getProviderURLs)("amoy") || "" };
        case "bscTestnet":
            return { chains: "bscTestnet", chainId: 4, rpc: (0, getProviderUrl_1.getProviderURLs)("bscTestnet") || "" };
        case 'beraTestnet':
            return { chains: "beraTestnet", chainId: 39, rpc: (0, getProviderUrl_1.getProviderURLs)("beraTestnet") || "" };
        case 'coreTestnet':
            return { chains: "coreTestnet", chainId: 4, rpc: (0, getProviderUrl_1.getProviderURLs)("coreTestnet") || "" };
        case 'baseSepolia':
            return { chains: "baseSepolia", chainId: 10004, rpc: (0, getProviderUrl_1.getProviderURLs)("baseSepolia") || "" };
        default:
            throw new Error("Not configuered");
    }
});
const PVP = (curNetwork) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const contractAddress = "0xB2A338Fb022365Aa40a2c7ADA3Bbf1Ae001D6dbe";
        const selector = "0x0ea901d2";
        const chain = yield getChainConf(curNetwork);
        const chains = [chain];
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
            :
                Promise.reject(new Error(`RPC URL is undefined for chain ${chainId}`))));
        if (responses == null) {
            throw new Error("Failed to make Eth calls");
        }
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
        }, { headers: { "X-API-Key": process.env.WORMHOLE_API_KEY } });
        if (response == null) {
            throw new Error("error querying cross chain");
        }
        const bytes = `0x${response.data.bytes}`;
        const signatures = response.data.signatures.map((s) => ({
            r: `0x${s.substring(0, 64)}`,
            s: `0x${s.substring(64, 128)}`,
            v: `0x${(parseInt(s.substring(128, 130), 16) + 27).toString(16)}`,
            guardianIndex: `0x${s.substring(130, 132)}`,
        }));
        return {
            "bytes": bytes,
            "sigs": signatures
        };
    }
    catch (err) {
        console.error("an error occurred during the cross-chain query process", err);
    }
});
exports.default = PVP;
