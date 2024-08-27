import express from "express";
import cors from 'cors';
import dotenv from "dotenv";
import { getUserStake } from "./handlers/getUserStake";
import { getClaimTS } from "./handlers/getClaimTs";
import { getLatestStakeTS } from "./handlers/getLatestStakeTS";
import { writeToDb } from "./handlers/setUserId";
import { readHeymintId } from "./handlers/getUserId";

const app = express();

dotenv.config();

var corsOptions = {
  origin: "*",
  optionsSuccessStatus: 200,
};

app.use(cors(corsOptions));

app.get("/", async (req, res) => {
  res.send("hey mint service");
});

app.get("/getUserStake/:network/:address", async (req, res) => {
  try {
    console.log(`Querying getUserStake for network: ${req.params.network}, address: ${req.params.address}`);
    console.log(req.params.network, req.params.address)
    const response = await getUserStake(req.params.network, req.params.address);
    console.log(`Queried getUserStake successfully.`);

    res.status(200).json(response); // Assuming response contains the required data
  } catch (error) {
    console.error("Error querying getUserStake:", error);
    res.status(500).json({ error: `Failed to perform CCQ: ${error}` });
  }
});

app.get("/getLatestStakeTs/:network/:address", async (req, res) => {
  try {
    console.log(`Querying getUserStake for network: ${req.params.network}, address: ${req.params.address}`);
    console.log(req.params.network, req.params.address)
    const response = await getLatestStakeTS(req.params.network, req.params.address);
    console.log(`Queried getUserStake successfully.`);

    res.status(200).json(response); // Assuming response contains the required data
  } catch (error) {
    console.error("Error querying getUserStake:", error);
    res.status(500).json({ error: `Failed to perform CCQ: ${error}` });
  }
});

app.get("/nextClaim/:network/:address", async (req, res) => {
  try {
    console.log(`Querying getClaimTS for network: ${req.params.network}, address: ${req.params.address}`);
    const response = await getClaimTS(req.params.network, req.params.address);
    console.log(`Queried getClaimTS successfully.`);

    res.status(200).json(response);
  } catch (error) {
    console.error("Error querying getClaimTS:", error);
    res.status(500).json({ error: `Failed to query ${req.params.network}: ${error}` });
  }
});

app.get("/getHeymMintID/:address", async (req, res) => {
  try {
    console.log(`Querying heymint id`);
    const response = await readHeymintId(req.params.address);
    console.log(`Queried heymint id successfully.`);

    res.status(200).json(response);
  } catch (error) {
    console.error("Error querying heymint id:", error);
    res.status(500).json({ error: `Failed to query heymint id : ${error}` });
  }
});

app.post("/setHeyMintID/:address/:id", async (req, res) => {
  try {
    console.log(`setting heymint id`);
    const response = await writeToDb(req.params.address, req.params.id);
    console.log(`successfully set heymint id for ${req.params.address}`);

    res.status(200).json(response);
  } catch (error) {
    console.error("Error setting heymint id:", error);
    res.status(500).json({ error: `Failed to set heymint id : ${error}` });
  }
});



export default app;
