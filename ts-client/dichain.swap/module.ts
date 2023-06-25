// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgDistributeCoins } from "./types/dichain/swap/tx";
import { MsgSwapCoin } from "./types/dichain/swap/tx";
import { MsgAddLiquidity } from "./types/dichain/swap/tx";
import { MsgMintCoins } from "./types/dichain/swap/tx";

import { Params as typeParams} from "./types"

export { MsgDistributeCoins, MsgSwapCoin, MsgAddLiquidity, MsgMintCoins };

type sendMsgDistributeCoinsParams = {
  value: MsgDistributeCoins,
  fee?: StdFee,
  memo?: string
};

type sendMsgSwapCoinParams = {
  value: MsgSwapCoin,
  fee?: StdFee,
  memo?: string
};

type sendMsgAddLiquidityParams = {
  value: MsgAddLiquidity,
  fee?: StdFee,
  memo?: string
};

type sendMsgMintCoinsParams = {
  value: MsgMintCoins,
  fee?: StdFee,
  memo?: string
};


type msgDistributeCoinsParams = {
  value: MsgDistributeCoins,
};

type msgSwapCoinParams = {
  value: MsgSwapCoin,
};

type msgAddLiquidityParams = {
  value: MsgAddLiquidity,
};

type msgMintCoinsParams = {
  value: MsgMintCoins,
};


export const registry = new Registry(msgTypes);

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	const structure: {fields: Field[]} = { fields: [] }
	for (let [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgDistributeCoins({ value, fee, memo }: sendMsgDistributeCoinsParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgDistributeCoins: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgDistributeCoins({ value: MsgDistributeCoins.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgDistributeCoins: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgSwapCoin({ value, fee, memo }: sendMsgSwapCoinParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgSwapCoin: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgSwapCoin({ value: MsgSwapCoin.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgSwapCoin: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAddLiquidity({ value, fee, memo }: sendMsgAddLiquidityParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAddLiquidity: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAddLiquidity({ value: MsgAddLiquidity.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAddLiquidity: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgMintCoins({ value, fee, memo }: sendMsgMintCoinsParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgMintCoins: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgMintCoins({ value: MsgMintCoins.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgMintCoins: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgDistributeCoins({ value }: msgDistributeCoinsParams): EncodeObject {
			try {
				return { typeUrl: "/dichain.swap.MsgDistributeCoins", value: MsgDistributeCoins.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgDistributeCoins: Could not create message: ' + e.message)
			}
		},
		
		msgSwapCoin({ value }: msgSwapCoinParams): EncodeObject {
			try {
				return { typeUrl: "/dichain.swap.MsgSwapCoin", value: MsgSwapCoin.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgSwapCoin: Could not create message: ' + e.message)
			}
		},
		
		msgAddLiquidity({ value }: msgAddLiquidityParams): EncodeObject {
			try {
				return { typeUrl: "/dichain.swap.MsgAddLiquidity", value: MsgAddLiquidity.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAddLiquidity: Could not create message: ' + e.message)
			}
		},
		
		msgMintCoins({ value }: msgMintCoinsParams): EncodeObject {
			try {
				return { typeUrl: "/dichain.swap.MsgMintCoins", value: MsgMintCoins.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgMintCoins: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	public structure: Record<string,unknown>;
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		this.structure =  {
						Params: getStructure(typeParams.fromPartial({})),
						
		};
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			DichainSwap: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;