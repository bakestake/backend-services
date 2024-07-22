import {getMessagesBySrcTxHash} from '@layerzerolabs/scan-client';
import { getChainId } from './utils/getChainId';

export const getStatus = async(chain:string, txHash : string) => {

    const srcId = await getChainId(chain)

    const msg  = await getMessagesBySrcTxHash(
        srcId,
        txHash,
    );

    const status = msg.messages[0].status

    return status;
}
