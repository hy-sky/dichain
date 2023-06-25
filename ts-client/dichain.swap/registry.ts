import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgDistributeCoins } from "./types/dichain/swap/tx";
import { MsgSwapCoin } from "./types/dichain/swap/tx";
import { MsgAddLiquidity } from "./types/dichain/swap/tx";
import { MsgMintCoins } from "./types/dichain/swap/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/dichain.swap.MsgDistributeCoins", MsgDistributeCoins],
    ["/dichain.swap.MsgSwapCoin", MsgSwapCoin],
    ["/dichain.swap.MsgAddLiquidity", MsgAddLiquidity],
    ["/dichain.swap.MsgMintCoins", MsgMintCoins],
    
];

export { msgTypes }