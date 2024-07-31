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
Object.defineProperty(exports, "__esModule", { value: true });
exports.getLatestStakeTS = void 0;
const ethers_1 = require("ethers");
const getProviderUrl_1 = require("./getProviderUrl");
const getterSetter_1 = require("../artifacts/getterSetter");
const getLatestStakeTS = (chain, userAddress) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const providerUrl = yield (0, getProviderUrl_1.getProviderURLs)(chain);
        if (!providerUrl) {
            throw new Error(`Unable to get provider URL for chain: ${chain}`);
        }
        const provider = new ethers_1.ethers.JsonRpcProvider(providerUrl);
        const wallet = new ethers_1.ethers.Wallet(process.env.PRIVATE_KEY || "", provider);
        const contractInst = new ethers_1.ethers.Contract("0x26705ad938791e61aa64a2a9d808378805aec819", getterSetter_1.getterSetter, wallet);
        const stakes = yield contractInst.getUserStakes(userAddress);
        if (stakes.length == 0) {
            return 0;
        }
        return Number(stakes[stakes.length - 1].timeStamp); // Convert BigInt to Number
    }
    catch (error) {
        console.error('Error in getUserStake:', error);
        return null; // or throw error if you want to handle it higher up
    }
});
exports.getLatestStakeTS = getLatestStakeTS;
