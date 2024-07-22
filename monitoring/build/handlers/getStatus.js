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
exports.getStatus = void 0;
const scan_client_1 = require("@layerzerolabs/scan-client");
const getChainId_1 = require("./utils/getChainId");
const getStatus = (chain, txHash) => __awaiter(void 0, void 0, void 0, function* () {
    const srcId = yield (0, getChainId_1.getChainId)(chain);
    const msg = yield (0, scan_client_1.getMessagesBySrcTxHash)(srcId, txHash);
    const status = msg.messages[0].status;
    return status;
});
exports.getStatus = getStatus;
