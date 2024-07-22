import express from "express"
import cors from 'cors'
import { getStatus } from "./handlers/getStatus";

const app = express();

var corsOptions = {
  origin: "*",
  optionsSuccessStatus: 200,
};

app.use(cors(corsOptions));

app.get("/", cors(), async (req, res) => {
  res.send("lz status service");
});

app.get("/getStatus/:chain/:txId", cors(), async (req, res) => {
  try{
    console.log(`Querying status`);
    const response = await getStatus(req.params.chain, req.params.txId);
    console.log(`status received.`);

    res.status(200);
    res.json({"status" :response});
  }catch(error){
    console.log("Error getting status", error)
   res.status(500).json({error: `failed to get status`})
  }
});

export default app;
