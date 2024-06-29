import express from "express"
import cors from 'cors'
import CCQ from "./handlers/CcqHandler";

const app = express();

var corsOptions = {
  origin: "*",
  optionsSuccessStatus: 200,
};

app.use(cors(corsOptions));

app.get("/", cors(), async (req, res) => {
  res.send("wormhole service");
});

app.post("/updateGlobalLiquidity/:network", cors(), async (req, res) => {
  try{
    console.log(`updating liquidity on ${req.params.network} ....`);
    await CCQ(req.params.network);
    console.log(`updated liquidity on ${req.params.network}.`);

    res.status(200);
    res.json(`updated liquidity on ${req.params.network}.`);
  }catch(error){
    console.log("Error updating global liquidity", error)
   res.status(500).json({error: `failed to update liquiidty on ${req.params.network}`})
  }
});

app.post("/updatePvpGlobalState/:network", cors(), async (req, res) => {
  /// call pvp handler 
  try{

  }catch(error){
    console.log("Error updating pvp global status", error)
   res.status(500).json({error: `failed to update pvp global status ${req.params.network}`})
  }
});


export default app;
