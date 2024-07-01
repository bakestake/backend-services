import express from "express"
import cors from 'cors'
import CCQ from "./handlers/CcqHandler";
import PVP from "./handlers/PvpHandler";

const app = express();

var corsOptions = {
  origin: "*",
  optionsSuccessStatus: 200,
};

app.use(cors(corsOptions));

app.get("/", cors(), async (req, res) => {
  res.send("wormhole service");
});

app.get("/updateGlobalLiquidity", cors(), async (req, res) => {
  try{
    console.log(`Querying across the network`);
    const response = await CCQ();
    console.log(`Queried all networks.`);

    res.status(200);
    res.json(response);
  }catch(error){
    console.log("Error updating global liquidity", error)
   res.status(500).json({error: `failed to perform CCQ`})
  }
});

app.post("/updatePvpGlobalState/:network", cors(), async (req, res) => {
  /// call pvp handler 
  try{
    console.log(`Querying on ${req.params.network} ....`);
    const response = await PVP(req.params.network);
    console.log(`Querying on ${req.params.network}.`);

    res.status(200);
    res.json(response);
  }catch(error){
    console.log("Error querying pvp global status", error)
   res.status(500).json({error: `failed to query pvp global status on ${req.params.network}`})
  }
});


export default app;
