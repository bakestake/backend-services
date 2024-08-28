
## Read APIs

  url - http://35.180.118.106:8000/
### 1. Home

- **Endpoint:** `GET /`
- **Description:** Returns a simple message indicating the home page.
- **Response:**
  ```json
  "Home! Put path variables for data"
  ```

### 2. Total Staked Buds Across All Chains

- **Endpoint:** `GET /totalStakedBudsAcrossAllChains`
- **Description:** Retrieves the total count of staked "Buds" across all chains.
- **Response:**
  ```json
  {
    "totalStakedBudsAcrossAllChains": <count>
  }
  ```

### 3. Get Events

- **Endpoint:** `GET /getEvents/:networkName/:function/:startBlock/:userAddress`
- **Description:** Retrieves events based on specified parameters.
- **Parameters:**
  - `networkName`: Name of the network
  - `function`: Event function
  - `startBlock`: Starting block
  - `userAddress`: User's address
- **Request body**
  - {
      user common.Address `json:user`
      startBlock big.Int `json:eid`
    }
- **Response:**
  ```json
  {
    "stake_found": <eventData>
  }
  ```

### 4. Get Current Block Number

- **Endpoint:** `GET /getCurrentBlockNumber/:networkName`
- **Description:** Retrieves the current block number of a specific network.
- **Parameters:**
  - `networkName`: Name of the network
- **Response:**
  ```json
  {
    "currentBlock": <blockNumber>
  }
  ```

### 5. Get APR (Annual Percentage Rate)

- **Endpoint:** `GET /getApr/`
- **Description:** Retrieves the current Annual Percentage Rate for all networks.
- **Response:**
  ```json
  array with aprs for corresponding chains ["amoy":1000, "arbSepolia":6999, .....]
  ```

### 6. Get Local Buds Count

- **Endpoint:** `GET /getLocalBudsCount/`
- **Description:** Retrieves the local count of "Buds" for all networks.
- **Response:**
  ```json
  
  array with local staked buds for corresponding chains ["amoy":1000, "arbSepolia":6999, .....]
  ```

### 7. Get Rewards

- **Endpoint:** `GET /getRewards/:address`
- **Description:** Retrieves current rewards for a specific address on all networks.
- **Parameters:**
  - `address`: User's address
- **Response:**
  ```json
  {
    "staking_rewards": ["amoy":73909,"arbSepolia":289,...]
  }
  ```

### 8. Get Stakers Count

- **Endpoint:** `GET /getStakersCount`
- **Description:** Retrieves the count of stakers on all network.
- **Response:**
  ```json
  {
    "stakers": ["amoy":200,"arbSepolia":289,...]
  }
  ```

### 9. Get Buds Balance

- **Endpoint:** `GET /getBudsBalance/:networkName`
- **Description:** Retrieves the balance of "Buds" for a specific network.
- **Parameters:**
  - `networkName`: Name of the network
- **Response:**
  ```json
  {
    "budsBalance": <balance>
  }
  ```

### 10. Get Fees

- **Endpoint:** `GET /getFees/:networkName`
- **Description:** Retrieves fees for a specific operation involving "Buds."
- **Parameters:**
  - `networkName`: Chain name
- **Request body**
  - {
      BudsAmount string       `json:"budsAmount"`
      TokenID    string       `json:"tokenId"`
      Address    string       `json:"userAddress"`
      DestEid    uint32       `json:"destEid"`
    }
- **Response:**
  ```json
  {
    "cctx_fees": <feesData>
  }
  ```

### 11. Most staked
- **Endpoint:** `GET /mostBaked`
- **Description:** Returns most staked network"
- **Response:**
  ```json
  {
    "most_staked": maxChain, "stakedAmount":maxDif
  }
  ```

### 12. Most Rekt
- **Endpoint:** `GET /mostRekt`
- **Description:** Returns most raided network"
- **Response:**
  ```json
  {
    "most_staked": maxChain, "lostAmount":maxDif
  }
  ```

## Wormhole
url : http://35.180.118.106:3000/
### 1. update liqduity

- **Endpoint:** `GET /updateGlobalLiquidity`
- **Description:** returns ccq data acquired from ccq to all chain chains. this data should be submitted to updateState function in contract
- header - API-KEY : B1q2a3z4w5s-A0p9lh-K5t6f7b-E4r5t6y
- **Response:**
  ```json
    {
      "bytes" : bytes,
      "sigs" : signatures
    }
  ```

## transaction confirmation
url - http://35.180.118.106:4000/

1. Cross chain tx monitoring
- **url** - getStatus/:param1/param2
- **param1** :- chain name
- **param2** :- source tx id
- **response** -
  ```json
    {
      "tx_completed" : true
    }
  ```

2. Stake tx confirmation
- **url** - /getStatus/stake/:chain/:startblock/:user
- **param1** :- chain name
- **param2** :- latest block number
- **param3** :- user address
- **response** -
  ```json
    { 
      "eventOccurred": boolean, 
      "amount": string
    }
  ```

3. Raid tx confirmation
- **url** - /getStatus/raid/:chain/:startblock/:user
- **param1** :- chain name
- **param2** :- latest block number
- **param3** :- user address
- **response** -
  ```json
    { 
      "eventOccurred": boolean, 
      "isSuccess": boolean
    }
  ```

4. Claim tx confirmation
- **url** - getStatus/rewardClaim/:chain/:startblock/:user
- **param1** :- chain name
- **param2** :- latest block number
- **param3** :- user address
- **response** -
  ```json
    {
      eventOccurred: boolean, 
      amount: uint
    }
  ```

5. Unstake tx confirmation
- **url** - /getStatus/unstake/:chain/:startblock/:user
- **param1** :- chain name
- **param2** :- latest block number
- **param3** :- user address
- **response** -
  ```json
    { 
      eventOccurred: boolean, 
      amount: uint, 
      tokenId:uint
    }
  ```

6. Buds claim confirmation
- **url** - /getStatus/budsClaim/:chain/:user
- **param1** :- chain name
- **param2** :- user address
- **response** -
  ```json
    {
      boolean [true(complete) or false(not found)]
    }
  ```

7. Farmer claim confirmation
- **url** -/getStatus/farmerClaim/:chain/:user
- **param1** :- chain name
- **param2** :- user address
- **response** -
  ```json
    {
      boolean [true(complete) or false(not found)]
    }
  ```

8. Narc claim confirmation
- **url** - /getStatus/narcClaim/:chain/:user
- **param1** :- chain name
- **param2** :- source tx id
- **response** -
  ```json
    {
      boolean [true(complete) or false(not found)]
    }
  ```

9. Buds burn confirmation
- **url** - /getStatus/burn/:chain/:startblock/:user
- **param1** :- chain name
- **param2** :- latest block number
- **param3** :- user address
- **response** -
  ```json
    {
      boosterMinted: [string] name of NFT minted [stoner or informat],
      tokenId: token id minted [uint]
    }
  ```

## Heymint 
url - http://35.180.118.106:6000/

### 1. Get User staked buds
- **Endpoint:** `GET /getUserStake/:network/:address`
- **Description:** Returns buds staked by user on particular chain"
- **Parameters:**
  - `network`: Chain name
  - `address`: address of user
- **Response:**
  ```json
  {
    "userStake" : budsAmount
  }
  ```
### 2. Get next claim timestamp
- **Endpoint:** `GET /nextClaim/:network/:address`
- **Description:** Returns timestamp for next buds claim for particular user on particular chain"
- **Parameters:**
  - `network`: Chain name
  - `address`: address of user
- **Response:**
  ```json
  {
    "timestamp" : timestamp
  }
  ```
### 3. Get next claim timestamp
- **Endpoint:** `GET /getLatestStakeTs/:network/:address`
- **Description:** Returns timestamp for latest stake"
- **Parameters:**
  - `network`: Chain name
  - `address`: address of user
- **Response:**
  ```json
  {
    "timestamp" : timestamp
  }
  ```

## Valid chain names for passing in params and request body-
1. amoy
2. bscTestnets
3. arbSepolia
4. baseSepolia
5. fuji
6. beraTestnet
