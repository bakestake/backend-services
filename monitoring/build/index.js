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
const app = (0, express_1.default)();
var corsOptions = {
    origin: "*",
    optionsSuccessStatus: 200,
};
app.use((0, cors_1.default)(corsOptions));
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
exports.default = app;
