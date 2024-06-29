"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.getProviderURLs = void 0;
const dotenv_1 = __importDefault(require("dotenv"));
dotenv_1.default.config({ path: "../.env" });
const getProviderURLs = (networkName) => {
    switch (networkName) {
        case "polygon":
            return process.env.POLYGON;
        case "amoy":
            return process.env.RPC_URL_AMOY;
        case "bsc":
            return process.env.BSC;
        case "bscTestnet":
            return process.env.RPC_URL_BSCTESTNET;
        case "avalanche":
            return process.env.AVALANCH;
        case "fuji":
            return process.env.RPC_URL_FUJI;
        case "arbitrum":
            return process.env.ARBITRUM;
        case "arbSepolia":
            return process.env.RPC_URL_ARBSEPOLIA;
        case "beraTestnet":
            return process.env.RPC_URL_BERA;
        case "baseSepolia":
            return process.env.RPC_URL_BASE_SEPOLIA;
        case "coreTestnet":
            return process.env.RPC_URL_CORETESTNET;
    }
};
exports.getProviderURLs = getProviderURLs;
