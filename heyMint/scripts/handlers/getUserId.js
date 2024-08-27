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
exports.readHeymintId = readHeymintId;
const setUserId_1 = require("./setUserId");
// Function to read heymint_id for a particular address
function readHeymintId(userAddress) {
    return __awaiter(this, void 0, void 0, function* () {
        const client = (0, setUserId_1.createDbClient)();
        try {
            yield client.connect();
            const query = `
            SELECT heymint_id FROM user_data WHERE user_address = $1;
        `;
            const res = yield client.query(query, [userAddress]);
            if (res.rows.length > 0) {
                return res.rows[0].heymint_id;
            }
            else {
                console.log(`No heymint_id found for address: ${userAddress}`);
                return null;
            }
        }
        catch (error) {
            console.error('Error reading from the database:', error);
            return null;
        }
        finally {
            yield client.end();
        }
    });
}
