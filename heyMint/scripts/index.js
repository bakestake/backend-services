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
const dotenv_1 = __importDefault(require("dotenv"));
const getUserStake_1 = require("./handlers/getUserStake");
const getClaimTs_1 = require("./handlers/getClaimTs");
const getLatestStakeTS_1 = require("./handlers/getLatestStakeTS");
const setUserId_1 = require("./handlers/setUserId");
const getUserId_1 = require("./handlers/getUserId");
const app = (0, express_1.default)();
dotenv_1.default.config();
var corsOptions = {
    origin: "*",
    optionsSuccessStatus: 200,
};
app.use((0, cors_1.default)(corsOptions));
app.use(function (req, res, next) {
    res.header('Access-Control-Allow-Origin', '*');
    next();
});
app.get("/", (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    res.send("hey mint service");
}));
app.get("/getUserStake/:network/:address", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        console.log(`Querying getUserStake for network: ${req.params.network}, address: ${req.params.address}`);
        console.log(req.params.network, req.params.address);
        const response = yield (0, getUserStake_1.getUserStake)(req.params.network, req.params.address);
        console.log(`Queried getUserStake successfully.`);
        res.status(200).json(response); // Assuming response contains the required data
    }
    catch (error) {
        console.error("Error querying getUserStake:", error);
        res.status(500).json({ error: `Failed to perform CCQ: ${error}` });
    }
}));
app.get("/getLatestStakeTs/:network/:address", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        console.log(`Querying getUserStake for network: ${req.params.network}, address: ${req.params.address}`);
        console.log(req.params.network, req.params.address);
        const response = yield (0, getLatestStakeTS_1.getLatestStakeTS)(req.params.network, req.params.address);
        console.log(`Queried getUserStake successfully.`);
        res.status(200).json(response); // Assuming response contains the required data
    }
    catch (error) {
        console.error("Error querying getUserStake:", error);
        res.status(500).json({ error: `Failed to perform CCQ: ${error}` });
    }
}));
app.get("/nextClaim/:network/:address", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        console.log(`Querying getClaimTS for network: ${req.params.network}, address: ${req.params.address}`);
        const response = yield (0, getClaimTs_1.getClaimTS)(req.params.network, req.params.address);
        console.log(`Queried getClaimTS successfully.`);
        res.status(200).json(response);
    }
    catch (error) {
        console.error("Error querying getClaimTS:", error);
        res.status(500).json({ error: `Failed to query ${req.params.network}: ${error}` });
    }
}));
app.get("/getHeymMintID/:address", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        console.log(`Querying heymint id`);
        const response = yield (0, getUserId_1.readHeymintId)(req.params.address);
        console.log(`Queried heymint id successfully.`);
        res.status(200).json(response);
    }
    catch (error) {
        console.error("Error querying heymint id:", error);
        res.status(500).json({ error: `Failed to query heymint id : ${error}` });
    }
}));
app.post("/setHeyMintID/:address/:id", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        console.log(`setting heymint id`);
        const response = yield (0, setUserId_1.writeToDb)(req.params.address, req.params.id);
        console.log(`successfully set heymint id for ${req.params.address}`);
        res.status(200).json(response);
    }
    catch (error) {
        console.error("Error setting heymint id:", error);
        res.status(500).json({ error: `Failed to set heymint id : ${error}` });
    }
}));
exports.default = app;
