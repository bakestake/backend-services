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
exports.BeraStateUpdate = void 0;
const ethers_1 = require("ethers");
const getProviderUrl_1 = require("./getProviderUrl");
const setter_1 = require("../artifacts/setter");
const axios_1 = __importDefault(require("axios"));
const dotenv_1 = __importDefault(require("dotenv"));
dotenv_1.default.config({ path: "../../.env" });
const BeraStateUpdate = () => __awaiter(void 0, void 0, void 0, function* () {
    var _a, _b, _c, _d, _e, _f, _g, _h;
    console.log("updating bera");
    const chains = [
        { chains: "fuji", chainId: 6, rpc: (0, getProviderUrl_1.getProviderURLs)("fuji") },
        { chains: "arbSepolia", chainId: 10003, rpc: (0, getProviderUrl_1.getProviderURLs)("arbSepolia") },
        { chains: "amoy", chainId: 10007, rpc: (0, getProviderUrl_1.getProviderURLs)("amoy") },
        { chains: "bscTestnet", chainId: 4, rpc: (0, getProviderUrl_1.getProviderURLs)("bscTestnet") },
        { chains: "beraTestnet", chainId: 39, rpc: (0, getProviderUrl_1.getProviderURLs)("beraTestnet") },
        // { chains: "coreTestnet", chainId: 4, rpc: getProviderURLs("coreTestnet") },
        { chains: "baseSepolia", chainId: 10004, rpc: (0, getProviderUrl_1.getProviderURLs)("baseSepolia") },
    ];
    console.log("getting responses from all chains");
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
                params: [{ to: "0xB2A338Fb022365Aa40a2c7ADA3Bbf1Ae001D6dbe", data: "0x4269e94c" }, "latest"],
            },
        ])
            .catch((error) => {
            console.error(`Error fetching data for rpc: ${rpc}`, error);
            return null;
        })
        : Promise.reject(new Error(`RPC URL is undefined for chain ${chainId}`))));
    console.log("summing up global");
    let globalBudsCount = 0;
    for (let i = 0; i < responses.length; i++) {
        console.log(parseInt((_d = (_c = (_b = (_a = responses[i]) === null || _a === void 0 ? void 0 : _a.data) === null || _b === void 0 ? void 0 : _b[0]) === null || _c === void 0 ? void 0 : _c.result) === null || _d === void 0 ? void 0 : _d.number));
        globalBudsCount += parseInt((_h = (_g = (_f = (_e = responses[i]) === null || _e === void 0 ? void 0 : _e.data) === null || _f === void 0 ? void 0 : _f[0]) === null || _g === void 0 ? void 0 : _g.result) === null || _h === void 0 ? void 0 : _h.number);
    }
    console.log(globalBudsCount);
    console.log("updating bera state");
    const setterInst = new ethers_1.ethers.Contract("0x054e0D70FfaefE4d7329B5077B4EcdE8062bfc08", setter_1.ABI, new ethers_1.ethers.Wallet(process.env.STATE_PRIVATE_KEY || "", new ethers_1.ethers.JsonRpcProvider((0, getProviderUrl_1.getProviderURLs)("beraTestnet") || "")));
    console.log("doing tx");
    const tx = yield setterInst.setGlobalStakedBuds(globalBudsCount);
    yield tx.wait();
});
exports.BeraStateUpdate = BeraStateUpdate;
