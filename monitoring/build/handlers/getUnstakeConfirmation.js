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
exports.getLastUnstakeEvent = void 0;
const ethers_1 = require("ethers");
const dotenv_1 = __importDefault(require("dotenv"));
dotenv_1.default.config({ path: "../.env" });
const eventAbi_1 = require("../artifacts/eventAbi");
const getProviderUrl_1 = require("./getProviderUrl");
const getLastUnstakeEvent = (chain, startBlock, userAddress) => __awaiter(void 0, void 0, void 0, function* () {
    const curUrl = yield (0, getProviderUrl_1.getProviderURLs)(chain);
    console.log(curUrl);
    try {
        const provider = new ethers_1.ethers.JsonRpcProvider(curUrl);
        console.log("provider", provider);
        const contract = new ethers_1.ethers.Contract("0xB2A338Fb022365Aa40a2c7ADA3Bbf1Ae001D6dbe", eventAbi_1.eventAbi, provider);
        const latestBlockNumber = yield provider.getBlockNumber();
        console.log("latest block", latestBlockNumber);
        var data2 = [];
        var filter = contract.filters.UnStaked(null, null, null, null, null, null);
        console.log("filter", filter);
        data2 = yield contract.queryFilter(filter, Number(startBlock), latestBlockNumber);
        console.log("data 2 filter :", data2);
        data2 = data2.map((el) => {
            const decodedEvent = contract.interface.decodeEventLog("UnStaked", el.data, el.topics);
            return {
                sender: decodedEvent[0],
                tokenId: Number(decodedEvent[1].toString()),
                amount: Number(decodedEvent[2].toString()) / 1e18,
                currentBlock: latestBlockNumber,
                transactionBlock: el.blockNumber,
            };
        });
        console.log("data 2", data2);
        for (let i = 0; i < data2.length; i++) {
            if (data2[i].sender == userAddress) {
                console.log("data[i] in events", data2[i]);
                console.log("true");
                return { eventOccurred: true, amount: data2[i].amount, tokenId: data2[i].tokenId };
            }
        }
        console.log("false");
        return { eventOccurred: false, amount: 0, tokenId: 0 };
    }
    catch (error) {
        console.error("Error in getLastEvent:", error);
        throw error;
    }
});
exports.getLastUnstakeEvent = getLastUnstakeEvent;
