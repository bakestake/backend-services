import { ADDRESSES } from "./constants";

export const getChainId = async (chain: string) => {
    switch (chain) {
        case "polygon":
            return ADDRESSES.polygon.endpointId;
        case "amoy":
            return ADDRESSES.amoy.endpointId;
        case "bsc":
            return ADDRESSES.bsc.endpointId;
        case "bscTestnet":
            return ADDRESSES.bscTestnet.endpointId;
        case "avalanche":
            return ADDRESSES.avalanche.endpointId;
        case "fuji":
            return ADDRESSES.fuji.endpointId;
        case "arbitrum":
            return ADDRESSES.arbitrum.endpointId;
        case "arbSepolia":
            return ADDRESSES.arbSepolia.endpointId;
        case "optimism":
            return ADDRESSES.optimism.endpointId;
        case "opGoerli":
            return ADDRESSES.opGoerli.endpointId;
        case "base":
            return ADDRESSES.base.endpointId;
        case "baseSepolia":
            return ADDRESSES.baseSpolia.endpointId;
        case "bera_atrio":
            return ADDRESSES.bera_atrio.endpointId;
        case "coreTestnet":
            return ADDRESSES.coreTestnet.endpointId;
        default:
            throw new Error(`Chain ${chain} not supported`);
    }
};
