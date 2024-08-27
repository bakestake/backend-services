import express from "express"
import cors from 'cors'
import { getStatus } from "./handlers/getStatus";
import { getLastStakeEvent } from "./handlers/getStakeConfirmation";
import { getLastRaidEvent } from "./handlers/getRaidConfirmation";
import { getLastUnstakeEvent } from "./handlers/getUnstakeConfirmation";
import { getLastRewardClaimEvent } from "./handlers/getRewardClaimConfirmation";
import { getBudsClaimConfirmation } from "./handlers/getBudsClaimConfirmation";
import { getNarcClaimConfirmation } from "./handlers/getNarcClaimConfirmation";
import { getFarmerClaimConfirmation } from "./handlers/getFarmerClaimConfirmation";
import { getBurnConfirmation } from "./handlers/getBurnConfirmation";

const app = express();

var corsOptions = {
  origin: "*",
  optionsSuccessStatus: 200,
};

app.use(cors(corsOptions));

app.use(function (req, res, next) {
  res.header('Access-Control-Allow-Origin', '*')
  next()
})

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

app.get("/getStatus/stake/:chain/:startblock/:user", cors(), async (req, res) => {
  try{
    console.log(`Querying stake confirmation`);
    const response = await getLastStakeEvent(req.params.chain, req.params.startblock, req.params.user);

    res.status(200);
    res.json({"status" :response});
  }catch(error){
    console.log("Error getting stake confirmation", error)
   res.status(500).json({error: `failed to get stake confirmation`})
  }
});


app.get("/getStatus/raid/:chain/:startblock/:user", cors(), async (req, res) => {
  try{
    console.log(`Querying raid confirmation`);
    const response = await getLastRaidEvent(req.params.chain, req.params.startblock, req.params.user);

    res.status(200);
    res.json({"status" :response});
  }catch(error){
    console.log("Error getting raid confirmation", error)
   res.status(500).json({error: `failed to get raid confirmation`})
  }
});

app.get("/getStatus/unstake/:chain/:startblock/:user", cors(), async (req, res) => {
  try{
    console.log(`Querying unstake confirmation`);
    const response = await getLastUnstakeEvent(req.params.chain, req.params.startblock, req.params.user);

    res.status(200);
    res.json({"status" :response});
  }catch(error){
    console.log("Error getting unstake confirmation", error)
   res.status(500).json({error: `failed to get unstake confirmation`})
  }
});

app.get("/getStatus/rewardClaim/:chain/:startblock/:user", cors(), async (req, res) => {
  try{
    console.log(`Querying reward claim confirmation`);
    const response = await getLastRewardClaimEvent(req.params.chain, req.params.startblock, req.params.user);

    res.status(200);
    res.json({"status" :response});
  }catch(error){
    console.log("Error getting reward claim confirmation", error)
   res.status(500).json({error: `failed to get reward claim confirmation`})
  }
});

app.get("/getStatus/budsClaim/:chain/:user", cors(), async (req, res) => {
  try{
    console.log(`Querying buds claim confirmation`);
    const response = await getBudsClaimConfirmation(req.params.chain, req.params.user);

    res.status(200);
    res.json({"status" :response});
  }catch(error){
    console.log("Error getting buds claim confirmation", error)
   res.status(500).json({error: `failed to get buds claim confirmation`})
  }
});

app.get("/getStatus/narcClaim/:chain/:user", cors(), async (req, res) => {
  try{
    console.log(`Querying narc claim confirmation`);
    const response = await getNarcClaimConfirmation(req.params.chain, req.params.user);

    res.status(200);
    res.json({"status" :response});
  }catch(error){
    console.log("Error getting narc claim confirmation", error)
   res.status(500).json({error: `failed to get narc claim confirmation`})
  }
});

app.get("/getStatus/farmerClaim/:chain/:user", cors(), async (req, res) => {
  try{
    console.log(`Querying farmer claim confirmation`);
    const response = await getFarmerClaimConfirmation(req.params.chain, req.params.user);

    res.status(200);
    res.json({"status" :response});
  }catch(error){
    console.log("Error getting farmer claim confirmation", error)
   res.status(500).json({error: `failed to get farmer claim confirmation`})
  }
});

app.get("/getStatus/burn/:chain/:startblock/:user", cors(), async (req, res) => {
  try{
    console.log(`Querying farmer claim confirmation`);
    const response = await getBurnConfirmation(req.params.chain, req.params.startblock, req.params.user);

    res.status(200);
    res.json({"status" :response});
  }catch(error){
    console.log("Error getting farmer claim confirmation", error)
   res.status(500).json({error: `failed to get farmer claim confirmation`})
  }
});

export default app;
