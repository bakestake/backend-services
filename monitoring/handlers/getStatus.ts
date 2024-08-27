import {createClient} from '@layerzerolabs/scan-client';


export const getStatus = async(chain:string, txHash : string) => {

    // Initialize a client with the desired environment
    const client = createClient('testnet');

    // Get a list of messages by transaction hash
    const {messages} = await client.getMessagesBySrcTxHash(
        txHash
    );

    console.log(messages)

    return messages;
}
