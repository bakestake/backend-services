import { Client, Pool } from 'pg';
import dotenv from "dotenv"
dotenv.config();

// Function to write data to the user_data table
export async function writeToDb(address: string, id: string, pool: Pool) {
  const query = `INSERT INTO your_table (address, heymint_id) VALUES ($1, $2)`;
  try {
    const client = await pool.connect();
    await client.query(query, [address, id]);
    client.release();
    return { message: "Data inserted successfully" };
  } catch (error) {
    console.error("Error inserting data into the database:", error);
    throw new Error("Failed to insert data");
  }
}

