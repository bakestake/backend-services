import { createDbClient } from "./setUserId";

// Function to read heymint_id for a particular address
export async function readHeymintId(userAddress: string): Promise<string | null> {
    const client = createDbClient();
    try {
        await client.connect();
        const query = `
            SELECT heymint_id FROM user_data WHERE user_address = $1;
        `;
        const res = await client.query(query, [userAddress]);
        if (res.rows.length > 0) {
            return res.rows[0].heymint_id;
        } else {
            console.log(`No heymint_id found for address: ${userAddress}`);
            return null;
        }
    } catch (error) {
        console.error('Error reading from the database:', error);
        return null;
    }// } finally {
    //     await client.end();
    // }
}
