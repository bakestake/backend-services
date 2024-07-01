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
const CcqHandler_1 = __importDefault(require("./handlers/CcqHandler"));
const PvpHandler_1 = __importDefault(require("./handlers/PvpHandler"));
const app = (0, express_1.default)();
var corsOptions = {
    origin: "*",
    optionsSuccessStatus: 200,
};
app.use((0, cors_1.default)(corsOptions));
app.use(function (req, res, next) {
    res.header('Access-Control-Allow-Origin', '*');
    res.header('Access-Control-Allow-Headers', 'Origin, Content-Type, Accept, API-KEY');
    next();
});
app.get("/", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    res.send("wormhole service");
}));
app.use((req, res, next) => {
    const apiKey = req.get('API-KEY');
    if (!apiKey || apiKey !== process.env.API_KEY) {
        res.status(401).json({ error: 'unauthorised' });
    }
    else {
        next();
    }
});
app.get("/updateGlobalLiquidity", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        console.log(`Querying across the network`);
        const response = yield (0, CcqHandler_1.default)();
        console.log(`Queried all networks.`);
        res.status(200);
        res.json(response);
    }
    catch (error) {
        console.log("Error updating global liquidity", error);
        res.status(500).json({ error: `failed to perform CCQ` });
    }
}));
app.get("/updatePvpGlobalState/:network", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    /// call pvp handler 
    try {
        console.log(`Querying on ${req.params.network} ....`);
        const response = yield (0, PvpHandler_1.default)(req.params.network);
        console.log(`Querying on ${req.params.network}.`);
        res.status(200);
        res.json(response);
    }
    catch (error) {
        console.log("Error querying pvp global status", error);
        res.status(500).json({ error: `failed to query pvp global status on ${req.params.network}` });
    }
}));
exports.default = app;
