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
const app = (0, express_1.default)();
var corsOptions = {
    origin: "*",
    optionsSuccessStatus: 200,
};
app.use((0, cors_1.default)(corsOptions));
app.get("/", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    res.send("wormhole service");
}));
app.post("/updateGlobalLiquidity/:network", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        console.log(`updating liquidity on ${req.params.network} ....`);
        yield (0, CcqHandler_1.default)(req.params.network);
        console.log(`updated liquidity on ${req.params.network}.`);
        res.status(200);
        res.json(`updated liquidity on ${req.params.network}.`);
    }
    catch (error) {
        console.log("Error updating global liquidity", error);
        res.status(500).json({ error: `failed to update liquiidty on ${req.params.network}` });
    }
}));
app.post("/updatePvpGlobalState/:network", (0, cors_1.default)(), (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    /// call pvp handler 
    try {
    }
    catch (error) {
        console.log("Error updating pvp global status", error);
        res.status(500).json({ error: `failed to update pvp global status ${req.params.network}` });
    }
}));
exports.default = app;
