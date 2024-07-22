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
exports.getChainId = void 0;
const constants_1 = require("./constants");
const getChainId = (chain) => __awaiter(void 0, void 0, void 0, function* () {
    switch (chain) {
        case "polygon":
            return constants_1.ADDRESSES.polygon.endpointId;
        case "amoy":
            return constants_1.ADDRESSES.amoy.endpointId;
        case "bsc":
            return constants_1.ADDRESSES.bsc.endpointId;
        case "bscTestnet":
            return constants_1.ADDRESSES.bscTestnet.endpointId;
        case "avalanche":
            return constants_1.ADDRESSES.avalanche.endpointId;
        case "fuji":
            return constants_1.ADDRESSES.fuji.endpointId;
        case "arbitrum":
            return constants_1.ADDRESSES.arbitrum.endpointId;
        case "arbSepolia":
            return constants_1.ADDRESSES.arbSepolia.endpointId;
        case "optimism":
            return constants_1.ADDRESSES.optimism.endpointId;
        case "opGoerli":
            return constants_1.ADDRESSES.opGoerli.endpointId;
        case "base":
            return constants_1.ADDRESSES.base.endpointId;
        case "baseSepolia":
            return constants_1.ADDRESSES.baseSpolia.endpointId;
        case "bera_atrio":
            return constants_1.ADDRESSES.bera_atrio.endpointId;
        case "coreTestnet":
            return constants_1.ADDRESSES.coreTestnet.endpointId;
        default:
            throw new Error(`Chain ${chain} not supported`);
    }
});
exports.getChainId = getChainId;
