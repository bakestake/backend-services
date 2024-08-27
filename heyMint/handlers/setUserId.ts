import { Client } from 'pg';
import dotenv from "dotenv"
dotenv.config();

// Function to initialize and return a PostgreSQL client
export function createDbClient() {
    return new Client({
        user: process.env.POSTGRES_USER,
        host: process.env.POSTGRES_HOST,
        database: process.env.POSTGRES_DB,
        password: process.env.POSTGRES_PASSWORD,
        port: 5000,
    });
}

// Function to write data to the user_data table
export async function writeToDb(userAddress: string, heymintId: string): Promise<void> {
    const client = createDbClient();
    try {
        await client.connect();
        const query = `
            INSERT INTO user_data (user_address, heymint_id)
            VALUES ($1, $2)
            ON CONFLICT (user_address) DO UPDATE SET heymint_id = EXCLUDED.heymint_id;
        `;
        await client.query(query, [userAddress, heymintId]);
        console.log(`Successfully wrote data for address: ${userAddress}`);
    } catch (error) {
        console.error('Error writing to the database:', error);
    } finally {
        await client.end();
    }
}

