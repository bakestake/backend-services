import express from "express"
import cors from 'cors'
import CCQ from "./handlers/CcqHandler";
import PVP from "./handlers/PvpHandler";
import dotenv from "dotenv";

const app = express();

dotenv.config()

var corsOptions = {
  origin: "*",
  optionsSuccessStatus: 200,
};

app.use(cors(corsOptions));

app.use(function (req, res, next) {
  res.header('Access-Control-Allow-Origin', '*')
  res.header('Access-Control-Allow-Headers', 'Origin, Content-Type, Accept, API-KEY')
  next()
})

app.get("/", cors(), async (req, res) => {
  res.send("wormhole service");
});

app.use((req, res, next) => {
  const apiKey = req.get('API-KEY')
  if (!apiKey || apiKey != process.env.API_KEY) {
    res.status(401).json({error: 'unauthorised api key'})
  } else {
    next()
  }
})

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

app.get("/updatePvpGlobalState/:network", cors(), async (req, res) => {
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
