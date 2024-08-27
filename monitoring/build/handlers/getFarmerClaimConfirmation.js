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
exports.getFarmerClaimConfirmation = void 0;
const ethers_1 = require("ethers");
const dotenv_1 = __importDefault(require("dotenv"));
dotenv_1.default.config({ path: "../.env" });
const getProviderUrl_1 = require("./getProviderUrl");
const nftFaucet_1 = require("../artifacts/nftFaucet");
const getFarmerClaimConfirmation = (chain, userAddress) => __awaiter(void 0, void 0, void 0, function* () {
    const curUrl = yield (0, getProviderUrl_1.getProviderURLs)(chain);
    console.log(curUrl);
    try {
        const provider = new ethers_1.ethers.JsonRpcProvider(curUrl);
        console.log("provider", provider);
        const contract = new ethers_1.ethers.Contract("0x16c2c9B38F59Ca3749D29a428d426558619c85E8", nftFaucet_1.nftFaucetAbi, provider);
        const latestBlockNumber = yield provider.getBlockNumber();
        console.log("latest block", latestBlockNumber);
        const claimed = yield contract.farmerClaimedBy(userAddress);
        return claimed;
    }
    catch (error) {
        console.error("Error in getLastEvent:", error);
        throw error;
    }
});
exports.getFarmerClaimConfirmation = getFarmerClaimConfirmation;
