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
exports.getClaimTS = void 0;
const ethers_1 = __importDefault(require("ethers"));
const getProviderUrl_1 = require("./getProviderUrl");
const faucet_1 = require("../artifacts/faucet");
const getClaimTS = (chain, userAddress) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const contractInst = new ethers_1.default.Contract("0xB2A338Fb022365Aa40a2c7ADA3Bbf1Ae001D6dbe", faucet_1.faucetABI, new ethers_1.default.Wallet(process.env.PRIVATE_KEY || "", new ethers_1.default.JsonRpcProvider((yield (0, getProviderUrl_1.getProviderURLs)(chain)) || "")));
        const ts = yield contractInst.nextClaimTimeInSeconds(userAddress);
        return ts;
    }
    catch (error) {
        console.log(error);
    }
});
exports.getClaimTS = getClaimTS;
