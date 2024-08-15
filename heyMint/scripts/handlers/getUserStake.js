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
exports.getUserStake = void 0;
const ethers_1 = require("ethers");
const getProviderUrl_1 = require("./getProviderUrl");
const getterSetter_1 = require("../artifacts/getterSetter");
const getUserStake = (chain, userAddress) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const providerUrl = yield (0, getProviderUrl_1.getProviderURLs)(chain);
        if (!providerUrl) {
            throw new Error(`Unable to get provider URL for chain: ${chain}`);
        }
        const provider = new ethers_1.ethers.JsonRpcProvider(providerUrl);
        const wallet = new ethers_1.ethers.Wallet(process.env.PRIVATE_KEY || "", provider);
        const contractInst = new ethers_1.ethers.Contract("0xB2A338Fb022365Aa40a2c7ADA3Bbf1Ae001D6dbe", getterSetter_1.getterSetter, wallet);
        const stakes = yield contractInst.getUserStakes(userAddress);
        let stakedAmount = BigInt(0);
        for (let i = 0; i < stakes.length; i++) {
            stakedAmount += BigInt(stakes[i].budsAmount);
        }
        return Number(stakedAmount) / 1e18; // Convert BigInt to Number
    }
    catch (error) {
        console.error('Error in getUserStake:', error);
        return null; // or throw error if you want to handle it higher up
    }
});
exports.getUserStake = getUserStake;
