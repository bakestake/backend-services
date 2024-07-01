export const ABI = [
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_budsToken",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "_budsVault",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "_wormhole",
          "type": "address"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "inputs": [],
      "name": "ChainIdMismatch",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "EmptyWormholeAddress",
      "type": "error"
    },
    {
      "inputs": [
        {
          "internalType": "uint8",
          "name": "val",
          "type": "uint8"
        }
      ],
      "name": "InvalidBoolVal",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "InvalidChainId",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "InvalidContractAddress",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "InvalidFunctionSignature",
      "type": "error"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "received",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "expected",
          "type": "uint256"
        }
      ],
      "name": "InvalidPayloadLength",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "InvalidResponseVersion",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "NumberOfResponsesMismatch",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "PreviousGameNotFinished",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "RequestTypeMismatch",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "StaleBlockNum",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "StaleBlockTime",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "UnAuthorized",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "UnexpectedNumberOfResults",
      "type": "error"
    },
    {
      "inputs": [
        {
          "internalType": "uint8",
          "name": "received",
          "type": "uint8"
        }
      ],
      "name": "UnsupportedQueryType",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "VersionMismatch",
      "type": "error"
    },
    {
      "inputs": [
        {
          "internalType": "uint8",
          "name": "received",
          "type": "uint8"
        },
        {
          "internalType": "uint8",
          "name": "expected",
          "type": "uint8"
        }
      ],
      "name": "WrongQueryType",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "ZeroAddress",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "ZeroQueries",
      "type": "error"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "playerOne",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "address",
          "name": "playerTwo",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "GameStarted",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "owner",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "Locked",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "owner",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "Winner",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "owner",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "depositeSlashed",
      "type": "event"
    },
    {
      "inputs": [],
      "name": "QT_ETH_CALL",
      "outputs": [
        {
          "internalType": "uint8",
          "name": "",
          "type": "uint8"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "QT_ETH_CALL_BY_TIMESTAMP",
      "outputs": [
        {
          "internalType": "uint8",
          "name": "",
          "type": "uint8"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "QT_ETH_CALL_WITH_FINALITY",
      "outputs": [
        {
          "internalType": "uint8",
          "name": "",
          "type": "uint8"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "QT_MAX",
      "outputs": [
        {
          "internalType": "uint8",
          "name": "",
          "type": "uint8"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "QT_SOL_ACCOUNT",
      "outputs": [
        {
          "internalType": "uint8",
          "name": "",
          "type": "uint8"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "QT_SOL_PDA",
      "outputs": [
        {
          "internalType": "uint8",
          "name": "",
          "type": "uint8"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "VERSION",
      "outputs": [
        {
          "internalType": "uint8",
          "name": "",
          "type": "uint8"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "budsToken",
      "outputs": [
        {
          "internalType": "contract IBudsToken",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "budsVault",
      "outputs": [
        {
          "internalType": "contract IBudsVault",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "user",
          "type": "address"
        }
      ],
      "name": "gameByUser",
      "outputs": [
        {
          "internalType": "address",
          "name": "opponent",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        },
        {
          "internalType": "bool",
          "name": "playerOneTurnedIn",
          "type": "bool"
        },
        {
          "internalType": "bool",
          "name": "playerTwoTurnedIn",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "gameServer",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "player",
          "type": "address"
        }
      ],
      "name": "getGameInfo",
      "outputs": [
        {
          "components": [
            {
              "internalType": "address",
              "name": "opponent",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "amount",
              "type": "uint256"
            },
            {
              "internalType": "bool",
              "name": "playerOneTurnedIn",
              "type": "bool"
            },
            {
              "internalType": "bool",
              "name": "playerTwoTurnedIn",
              "type": "bool"
            }
          ],
          "internalType": "struct Escrow.Game",
          "name": "game",
          "type": "tuple"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes",
          "name": "response",
          "type": "bytes"
        }
      ],
      "name": "getResponseDigest",
      "outputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes",
          "name": "response",
          "type": "bytes"
        }
      ],
      "name": "getResponseHash",
      "outputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [
        {
          "components": [
            {
              "internalType": "address",
              "name": "opponent",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "amount",
              "type": "uint256"
            },
            {
              "internalType": "bool",
              "name": "playerOneTurnedIn",
              "type": "bool"
            },
            {
              "internalType": "bool",
              "name": "playerTwoTurnedIn",
              "type": "bool"
            }
          ],
          "internalType": "struct Escrow.Game",
          "name": "game",
          "type": "tuple"
        },
        {
          "internalType": "bytes",
          "name": "response",
          "type": "bytes"
        },
        {
          "components": [
            {
              "internalType": "bytes32",
              "name": "r",
              "type": "bytes32"
            },
            {
              "internalType": "bytes32",
              "name": "s",
              "type": "bytes32"
            },
            {
              "internalType": "uint8",
              "name": "v",
              "type": "uint8"
            },
            {
              "internalType": "uint8",
              "name": "guardianIndex",
              "type": "uint8"
            }
          ],
          "internalType": "struct IWormhole.Signature[]",
          "name": "signatures",
          "type": "tuple[]"
        }
      ],
      "name": "lockFunds",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes",
          "name": "response",
          "type": "bytes"
        },
        {
          "components": [
            {
              "internalType": "bytes32",
              "name": "r",
              "type": "bytes32"
            },
            {
              "internalType": "bytes32",
              "name": "s",
              "type": "bytes32"
            },
            {
              "internalType": "uint8",
              "name": "v",
              "type": "uint8"
            },
            {
              "internalType": "uint8",
              "name": "guardianIndex",
              "type": "uint8"
            }
          ],
          "internalType": "struct IWormhole.Signature[]",
          "name": "signatures",
          "type": "tuple[]"
        }
      ],
      "name": "parseAndVerifyQueryResponse",
      "outputs": [
        {
          "components": [
            {
              "internalType": "uint8",
              "name": "version",
              "type": "uint8"
            },
            {
              "internalType": "uint16",
              "name": "senderChainId",
              "type": "uint16"
            },
            {
              "internalType": "uint32",
              "name": "nonce",
              "type": "uint32"
            },
            {
              "internalType": "bytes",
              "name": "requestId",
              "type": "bytes"
            },
            {
              "components": [
                {
                  "internalType": "uint16",
                  "name": "chainId",
                  "type": "uint16"
                },
                {
                  "internalType": "uint8",
                  "name": "queryType",
                  "type": "uint8"
                },
                {
                  "internalType": "bytes",
                  "name": "request",
                  "type": "bytes"
                },
                {
                  "internalType": "bytes",
                  "name": "response",
                  "type": "bytes"
                }
              ],
              "internalType": "struct ParsedPerChainQueryResponse[]",
              "name": "responses",
              "type": "tuple[]"
            }
          ],
          "internalType": "struct ParsedQueryResponse",
          "name": "r",
          "type": "tuple"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "components": [
            {
              "internalType": "uint16",
              "name": "chainId",
              "type": "uint16"
            },
            {
              "internalType": "uint8",
              "name": "queryType",
              "type": "uint8"
            },
            {
              "internalType": "bytes",
              "name": "request",
              "type": "bytes"
            },
            {
              "internalType": "bytes",
              "name": "response",
              "type": "bytes"
            }
          ],
          "internalType": "struct ParsedPerChainQueryResponse",
          "name": "pcr",
          "type": "tuple"
        }
      ],
      "name": "parseEthCallByTimestampQueryResponse",
      "outputs": [
        {
          "components": [
            {
              "internalType": "bytes",
              "name": "requestTargetBlockIdHint",
              "type": "bytes"
            },
            {
              "internalType": "bytes",
              "name": "requestFollowingBlockIdHint",
              "type": "bytes"
            },
            {
              "internalType": "uint64",
              "name": "requestTargetTimestamp",
              "type": "uint64"
            },
            {
              "internalType": "uint64",
              "name": "targetBlockNum",
              "type": "uint64"
            },
            {
              "internalType": "uint64",
              "name": "targetBlockTime",
              "type": "uint64"
            },
            {
              "internalType": "uint64",
              "name": "followingBlockNum",
              "type": "uint64"
            },
            {
              "internalType": "bytes32",
              "name": "targetBlockHash",
              "type": "bytes32"
            },
            {
              "internalType": "bytes32",
              "name": "followingBlockHash",
              "type": "bytes32"
            },
            {
              "internalType": "uint64",
              "name": "followingBlockTime",
              "type": "uint64"
            },
            {
              "components": [
                {
                  "internalType": "address",
                  "name": "contractAddress",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "callData",
                  "type": "bytes"
                },
                {
                  "internalType": "bytes",
                  "name": "result",
                  "type": "bytes"
                }
              ],
              "internalType": "struct EthCallData[]",
              "name": "result",
              "type": "tuple[]"
            }
          ],
          "internalType": "struct EthCallByTimestampQueryResponse",
          "name": "r",
          "type": "tuple"
        }
      ],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [
        {
          "components": [
            {
              "internalType": "uint16",
              "name": "chainId",
              "type": "uint16"
            },
            {
              "internalType": "uint8",
              "name": "queryType",
              "type": "uint8"
            },
            {
              "internalType": "bytes",
              "name": "request",
              "type": "bytes"
            },
            {
              "internalType": "bytes",
              "name": "response",
              "type": "bytes"
            }
          ],
          "internalType": "struct ParsedPerChainQueryResponse",
          "name": "pcr",
          "type": "tuple"
        }
      ],
      "name": "parseEthCallQueryResponse",
      "outputs": [
        {
          "components": [
            {
              "internalType": "bytes",
              "name": "requestBlockId",
              "type": "bytes"
            },
            {
              "internalType": "uint64",
              "name": "blockNum",
              "type": "uint64"
            },
            {
              "internalType": "uint64",
              "name": "blockTime",
              "type": "uint64"
            },
            {
              "internalType": "bytes32",
              "name": "blockHash",
              "type": "bytes32"
            },
            {
              "components": [
                {
                  "internalType": "address",
                  "name": "contractAddress",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "callData",
                  "type": "bytes"
                },
                {
                  "internalType": "bytes",
                  "name": "result",
                  "type": "bytes"
                }
              ],
              "internalType": "struct EthCallData[]",
              "name": "result",
              "type": "tuple[]"
            }
          ],
          "internalType": "struct EthCallQueryResponse",
          "name": "r",
          "type": "tuple"
        }
      ],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [
        {
          "components": [
            {
              "internalType": "uint16",
              "name": "chainId",
              "type": "uint16"
            },
            {
              "internalType": "uint8",
              "name": "queryType",
              "type": "uint8"
            },
            {
              "internalType": "bytes",
              "name": "request",
              "type": "bytes"
            },
            {
              "internalType": "bytes",
              "name": "response",
              "type": "bytes"
            }
          ],
          "internalType": "struct ParsedPerChainQueryResponse",
          "name": "pcr",
          "type": "tuple"
        }
      ],
      "name": "parseEthCallWithFinalityQueryResponse",
      "outputs": [
        {
          "components": [
            {
              "internalType": "bytes",
              "name": "requestBlockId",
              "type": "bytes"
            },
            {
              "internalType": "bytes",
              "name": "requestFinality",
              "type": "bytes"
            },
            {
              "internalType": "uint64",
              "name": "blockNum",
              "type": "uint64"
            },
            {
              "internalType": "uint64",
              "name": "blockTime",
              "type": "uint64"
            },
            {
              "internalType": "bytes32",
              "name": "blockHash",
              "type": "bytes32"
            },
            {
              "components": [
                {
                  "internalType": "address",
                  "name": "contractAddress",
                  "type": "address"
                },
                {
                  "internalType": "bytes",
                  "name": "callData",
                  "type": "bytes"
                },
                {
                  "internalType": "bytes",
                  "name": "result",
                  "type": "bytes"
                }
              ],
              "internalType": "struct EthCallData[]",
              "name": "result",
              "type": "tuple[]"
            }
          ],
          "internalType": "struct EthCallWithFinalityQueryResponse",
          "name": "r",
          "type": "tuple"
        }
      ],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [
        {
          "components": [
            {
              "internalType": "uint16",
              "name": "chainId",
              "type": "uint16"
            },
            {
              "internalType": "uint8",
              "name": "queryType",
              "type": "uint8"
            },
            {
              "internalType": "bytes",
              "name": "request",
              "type": "bytes"
            },
            {
              "internalType": "bytes",
              "name": "response",
              "type": "bytes"
            }
          ],
          "internalType": "struct ParsedPerChainQueryResponse",
          "name": "pcr",
          "type": "tuple"
        }
      ],
      "name": "parseSolanaAccountQueryResponse",
      "outputs": [
        {
          "components": [
            {
              "internalType": "bytes",
              "name": "requestCommitment",
              "type": "bytes"
            },
            {
              "internalType": "uint64",
              "name": "requestMinContextSlot",
              "type": "uint64"
            },
            {
              "internalType": "uint64",
              "name": "requestDataSliceOffset",
              "type": "uint64"
            },
            {
              "internalType": "uint64",
              "name": "requestDataSliceLength",
              "type": "uint64"
            },
            {
              "internalType": "uint64",
              "name": "slotNumber",
              "type": "uint64"
            },
            {
              "internalType": "uint64",
              "name": "blockTime",
              "type": "uint64"
            },
            {
              "internalType": "bytes32",
              "name": "blockHash",
              "type": "bytes32"
            },
            {
              "components": [
                {
                  "internalType": "bytes32",
                  "name": "account",
                  "type": "bytes32"
                },
                {
                  "internalType": "uint64",
                  "name": "lamports",
                  "type": "uint64"
                },
                {
                  "internalType": "uint64",
                  "name": "rentEpoch",
                  "type": "uint64"
                },
                {
                  "internalType": "bool",
                  "name": "executable",
                  "type": "bool"
                },
                {
                  "internalType": "bytes32",
                  "name": "owner",
                  "type": "bytes32"
                },
                {
                  "internalType": "bytes",
                  "name": "data",
                  "type": "bytes"
                }
              ],
              "internalType": "struct SolanaAccountResult[]",
              "name": "results",
              "type": "tuple[]"
            }
          ],
          "internalType": "struct SolanaAccountQueryResponse",
          "name": "r",
          "type": "tuple"
        }
      ],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [
        {
          "components": [
            {
              "internalType": "uint16",
              "name": "chainId",
              "type": "uint16"
            },
            {
              "internalType": "uint8",
              "name": "queryType",
              "type": "uint8"
            },
            {
              "internalType": "bytes",
              "name": "request",
              "type": "bytes"
            },
            {
              "internalType": "bytes",
              "name": "response",
              "type": "bytes"
            }
          ],
          "internalType": "struct ParsedPerChainQueryResponse",
          "name": "pcr",
          "type": "tuple"
        }
      ],
      "name": "parseSolanaPdaQueryResponse",
      "outputs": [
        {
          "components": [
            {
              "internalType": "bytes",
              "name": "requestCommitment",
              "type": "bytes"
            },
            {
              "internalType": "uint64",
              "name": "requestMinContextSlot",
              "type": "uint64"
            },
            {
              "internalType": "uint64",
              "name": "requestDataSliceOffset",
              "type": "uint64"
            },
            {
              "internalType": "uint64",
              "name": "requestDataSliceLength",
              "type": "uint64"
            },
            {
              "internalType": "uint64",
              "name": "slotNumber",
              "type": "uint64"
            },
            {
              "internalType": "uint64",
              "name": "blockTime",
              "type": "uint64"
            },
            {
              "internalType": "bytes32",
              "name": "blockHash",
              "type": "bytes32"
            },
            {
              "components": [
                {
                  "internalType": "bytes32",
                  "name": "programId",
                  "type": "bytes32"
                },
                {
                  "internalType": "bytes[]",
                  "name": "seeds",
                  "type": "bytes[]"
                },
                {
                  "internalType": "bytes32",
                  "name": "account",
                  "type": "bytes32"
                },
                {
                  "internalType": "uint64",
                  "name": "lamports",
                  "type": "uint64"
                },
                {
                  "internalType": "uint64",
                  "name": "rentEpoch",
                  "type": "uint64"
                },
                {
                  "internalType": "bool",
                  "name": "executable",
                  "type": "bool"
                },
                {
                  "internalType": "bytes32",
                  "name": "owner",
                  "type": "bytes32"
                },
                {
                  "internalType": "bytes",
                  "name": "data",
                  "type": "bytes"
                },
                {
                  "internalType": "uint8",
                  "name": "bump",
                  "type": "uint8"
                }
              ],
              "internalType": "struct SolanaPdaResult[]",
              "name": "results",
              "type": "tuple[]"
            }
          ],
          "internalType": "struct SolanaPdaQueryResponse",
          "name": "r",
          "type": "tuple"
        }
      ],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "responsePrefix",
      "outputs": [
        {
          "internalType": "bytes",
          "name": "",
          "type": "bytes"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "selector",
      "outputs": [
        {
          "internalType": "bytes4",
          "name": "",
          "type": "bytes4"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "player",
          "type": "address"
        },
        {
          "internalType": "bool",
          "name": "result",
          "type": "bool"
        }
      ],
      "name": "settleGame",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint64",
          "name": "_blockNum",
          "type": "uint64"
        },
        {
          "internalType": "uint256",
          "name": "_minBlockNum",
          "type": "uint256"
        }
      ],
      "name": "validateBlockNum",
      "outputs": [],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint64",
          "name": "_blockTime",
          "type": "uint64"
        },
        {
          "internalType": "uint256",
          "name": "_minBlockTime",
          "type": "uint256"
        }
      ],
      "name": "validateBlockTime",
      "outputs": [],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint16",
          "name": "chainId",
          "type": "uint16"
        },
        {
          "internalType": "uint16[]",
          "name": "_validChainIds",
          "type": "uint16[]"
        }
      ],
      "name": "validateChainId",
      "outputs": [],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [
        {
          "components": [
            {
              "internalType": "address",
              "name": "contractAddress",
              "type": "address"
            },
            {
              "internalType": "bytes",
              "name": "callData",
              "type": "bytes"
            },
            {
              "internalType": "bytes",
              "name": "result",
              "type": "bytes"
            }
          ],
          "internalType": "struct EthCallData",
          "name": "r",
          "type": "tuple"
        },
        {
          "internalType": "address[]",
          "name": "_expectedContractAddresses",
          "type": "address[]"
        },
        {
          "internalType": "bytes4[]",
          "name": "_expectedFunctionSignatures",
          "type": "bytes4[]"
        }
      ],
      "name": "validateEthCallData",
      "outputs": [],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [
        {
          "components": [
            {
              "internalType": "address",
              "name": "contractAddress",
              "type": "address"
            },
            {
              "internalType": "bytes",
              "name": "callData",
              "type": "bytes"
            },
            {
              "internalType": "bytes",
              "name": "result",
              "type": "bytes"
            }
          ],
          "internalType": "struct EthCallData[]",
          "name": "r",
          "type": "tuple[]"
        },
        {
          "internalType": "address[]",
          "name": "_expectedContractAddresses",
          "type": "address[]"
        },
        {
          "internalType": "bytes4[]",
          "name": "_expectedFunctionSignatures",
          "type": "bytes4[]"
        }
      ],
      "name": "validateMultipleEthCallData",
      "outputs": [],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes",
          "name": "response",
          "type": "bytes"
        },
        {
          "components": [
            {
              "internalType": "bytes32",
              "name": "r",
              "type": "bytes32"
            },
            {
              "internalType": "bytes32",
              "name": "s",
              "type": "bytes32"
            },
            {
              "internalType": "uint8",
              "name": "v",
              "type": "uint8"
            },
            {
              "internalType": "uint8",
              "name": "guardianIndex",
              "type": "uint8"
            }
          ],
          "internalType": "struct IWormhole.Signature[]",
          "name": "signatures",
          "type": "tuple[]"
        }
      ],
      "name": "verifyQueryResponseSignatures",
      "outputs": [],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "wormhole",
      "outputs": [
        {
          "internalType": "contract IWormhole",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    }
  ]