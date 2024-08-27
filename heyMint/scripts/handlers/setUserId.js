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
exports.createDbClient = createDbClient;
exports.writeToDb = writeToDb;
const pg_1 = require("pg");
const dotenv_1 = __importDefault(require("dotenv"));
dotenv_1.default.config();
// Function to initialize and return a PostgreSQL client
function createDbClient() {
    return new pg_1.Client({
        user: process.env.POSTGRES_USER,
        host: process.env.POSTGRES_HOST,
        database: process.env.POSTGRES_DB,
        password: process.env.POSTGRES_PASSWORD,
        port: 5000,
    });
}
// Function to write data to the user_data table
function writeToDb(userAddress, heymintId) {
    return __awaiter(this, void 0, void 0, function* () {
        const client = createDbClient();
        try {
            yield client.connect();
            const query = `
            INSERT INTO user_data (user_address, heymint_id)
            VALUES ($1, $2)
            ON CONFLICT (user_address) DO UPDATE SET heymint_id = EXCLUDED.heymint_id;
        `;
            yield client.query(query, [userAddress, heymintId]);
            console.log(`Successfully wrote data for address: ${userAddress}`);
        }
        catch (error) {
            console.error('Error writing to the database:', error);
        }
        finally {
            yield client.end();
        }
    });
}
