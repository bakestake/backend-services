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
const express_1 = __importDefault(require("express"));
const cors_1 = __importDefault(require("cors"));
const getStatus_1 = require("./handlers/getStatus");
const getStakeConfirmation_1 = require("./handlers/getStakeConfirmation");
const getRaidConfirmation_1 = require("./handlers/getRaidConfirmation");
const getUnstakeConfirmation_1 = require("./handlers/getUnstakeConfirmation");
const getRewardClaimConfirmation_1 = require("./handlers/getRewardClaimConfirmation");
const getBudsClaimConfirmation_1 = require("./handlers/getBudsClaimConfirmation");
const getNarcClaimConfirmation_1 = require("./handlers/getNarcClaimConfirmation");
const getFarmerClaimConfirmation_1 = require("./handlers/getFarmerClaimConfirmation");
const getBurnConfirmation_1 = require("./handlers/getBurnConfirmation");
const app = (0, express_1.default)();
var corsOptions = {
    origin: "*",
    optionsSuccessStatus: 200,
};
app.use((0, cors_1.default)(corsOptions));
app.use(function (req, res, next) {
    res.header('Access-Control-Allow-Origin', '*');
    next();
});
app.get("/", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    res.send("lz status service");
}));
app.get("/getStatus/:chain/:txId", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        console.log(`Querying status`);
        const response = yield (0, getStatus_1.getStatus)(req.params.chain, req.params.txId);
        console.log(`status received.`);
        res.status(200);
        res.json({ "status": response });
    }
    catch (error) {
        console.log("Error getting status", error);
        res.status(500).json({ error: `failed to get status` });
    }
}));
app.get("/getStatus/stake/:chain/:startblock/:user", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        console.log(`Querying stake confirmation`);
        const response = yield (0, getStakeConfirmation_1.getLastStakeEvent)(req.params.chain, req.params.startblock, req.params.user);
        res.status(200);
        res.json({ "status": response });
    }
    catch (error) {
        console.log("Error getting stake confirmation", error);
        res.status(500).json({ error: `failed to get stake confirmation` });
    }
}));
app.get("/getStatus/raid/:chain/:startblock/:user", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        console.log(`Querying raid confirmation`);
        const response = yield (0, getRaidConfirmation_1.getLastRaidEvent)(req.params.chain, req.params.startblock, req.params.user);
        res.status(200);
        res.json({ "status": response });
    }
    catch (error) {
        console.log("Error getting raid confirmation", error);
        res.status(500).json({ error: `failed to get raid confirmation` });
    }
}));
app.get("/getStatus/unstake/:chain/:startblock/:user", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        console.log(`Querying unstake confirmation`);
        const response = yield (0, getUnstakeConfirmation_1.getLastUnstakeEvent)(req.params.chain, req.params.startblock, req.params.user);
        res.status(200);
        res.json({ "status": response });
    }
    catch (error) {
        console.log("Error getting unstake confirmation", error);
        res.status(500).json({ error: `failed to get unstake confirmation` });
    }
}));
app.get("/getStatus/rewardClaim/:chain/:startblock/:user", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        console.log(`Querying reward claim confirmation`);
        const response = yield (0, getRewardClaimConfirmation_1.getLastRewardClaimEvent)(req.params.chain, req.params.startblock, req.params.user);
        res.status(200);
        res.json({ "status": response });
    }
    catch (error) {
        console.log("Error getting reward claim confirmation", error);
        res.status(500).json({ error: `failed to get reward claim confirmation` });
    }
}));
app.get("/getStatus/budsClaim/:chain/:user", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        console.log(`Querying buds claim confirmation`);
        const response = yield (0, getBudsClaimConfirmation_1.getBudsClaimConfirmation)(req.params.chain, req.params.user);
        res.status(200);
        res.json({ "status": response });
    }
    catch (error) {
        console.log("Error getting buds claim confirmation", error);
        res.status(500).json({ error: `failed to get buds claim confirmation` });
    }
}));
app.get("/getStatus/narcClaim/:chain/:user", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        console.log(`Querying narc claim confirmation`);
        const response = yield (0, getNarcClaimConfirmation_1.getNarcClaimConfirmation)(req.params.chain, req.params.user);
        res.status(200);
        res.json({ "status": response });
    }
    catch (error) {
        console.log("Error getting narc claim confirmation", error);
        res.status(500).json({ error: `failed to get narc claim confirmation` });
    }
}));
app.get("/getStatus/farmerClaim/:chain/:user", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        console.log(`Querying farmer claim confirmation`);
        const response = yield (0, getFarmerClaimConfirmation_1.getFarmerClaimConfirmation)(req.params.chain, req.params.user);
        res.status(200);
        res.json({ "status": response });
    }
    catch (error) {
        console.log("Error getting farmer claim confirmation", error);
        res.status(500).json({ error: `failed to get farmer claim confirmation` });
    }
}));
app.get("/getStatus/burn/:chain/:startblock/:user", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        console.log(`Querying farmer claim confirmation`);
        const response = yield (0, getBurnConfirmation_1.getBurnConfirmation)(req.params.chain, req.params.startblock, req.params.user);
        res.status(200);
        res.json({ "status": response });
    }
    catch (error) {
        console.log("Error getting farmer claim confirmation", error);
        res.status(500).json({ error: `failed to get farmer claim confirmation` });
    }
}));
exports.default = app;
