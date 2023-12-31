<template>
  <div>
    <div class="pt-8">
      <div class="text-xs text-gray-600">Send to</div>

      <div>
        <input
          v-model="state.tx.receiver"
          class="mt-1 py-2 px-4 h-12 bg-gray-100 border-xs text-base leading-tight w-full rounded-xl outline-0"
          :class="{
            'border border-red-400':
              state.tx.receiver.length > 0 && !validReceiver,
          }"
          placeholder="Recipient address"
          :disabled="!hasAnyBalance"
        />
        <div
          v-if="state.tx.receiver.length > 0 && !validReceiver"
          class="text-xs text-red-400 mt-1"
        >
          Invalid address
        </div>
      </div>
    </div>
    <div v-if="hasAnyBalance">
      <IgntAmountSelect
        class="token-selector--main"
        :selected="state.tx.amounts"
        :balances="(balances.assets as Amount[])"
        @update="handleTxAmountUpdate"
      />
    </div>

    <div
      class="flex text-xs font-semibold items-center mt-8"
      :class="[
        {
          'cursor-pointer': hasAnyBalance,
          'text-gray-400': !hasAnyBalance,
        },
      ]"
      @click="
        (evt: MouseEvent) => {
          toggleAdvanced();
					return evt;
        }
      "
    >
      Advanced
      <template v-if="hasAnyBalance">
        <IgntChevronDownIcon
          :class="{ 'rotate-180': state.advancedOpen }"
          class="ml-1"
        />
      </template>
    </div>

    <div
      v-if="state.advancedOpen && hasAnyBalance"
      style="width: 100%; height: 24px"
    />

    <div v-if="state.advancedOpen && hasAnyBalance" class="advanced">
      <div class="text-xs pb-2">Fees</div>

      <IgntAmountSelect
        class="token-selector"
        :selected="state.tx.fees"
        :balances="(balances.assets as Amount[])"
        @update="handleTxFeesUpdate"
      />

      <div class="text-xs mt-8 text-gray-600">Reference (memo)</div>

      <div class="mb-4">
        <input
          v-model="state.tx.memo"
          class="mt-1 py-2 px-4 h-12 bg-gray-100 border-xs text-base leading-tight w-full rounded-xl outline-0"
          placeholder="Enter a reference"
        />
      </div>

      <div class="text-xs text-gray-600">Channel</div>

      <div class="input-wrapper">
        <input
          v-model="state.tx.ch"
          class="mt-1 py-2 px-4 h-12 bg-gray-100 border-xs text-base leading-tight w-full rounded-xl outline-0"
          placeholder="Enter a channel"
        />
      </div>
    </div>

    <div style="width: 100%; height: 24px" />

    <div>
      <IgntButton
        style="width: 100%"
        :disabled="!ableToTx"
        @click="sendTx"
        :busy="isTxOngoing"
        >Send</IgntButton
      >
      <div
        v-if="isTxError"
        class="flex items-center justify-center text-xs text-red-500 italic mt-2"
      >
        Error submitting Tx
      </div>

      <div
        v-if="isTxSuccess"
        class="flex items-center justify-center text-xs text-green-500 italic mt-2"
      >
        Tx submitted succesfully
      </div>
    </div>

    <div>
      <IgntButton
          style="width: 100%; margin-top: 20px;"
          @click="mintCoinsTx"
          :busy="isTxOngoing"
      >Mint 100000000eth,10000000btc</IgntButton
      >
    </div>
  </div>
</template>
<script setup lang="ts">
import { fromBech32 } from "@cosmjs/encoding";
import { useAddress } from "@/def-composables/useAddress";
import { useAssets } from "@/def-composables/useAssets";
import type { Amount } from "@/utils/interfaces";
import { reactive } from "vue";
import Long from "long";
import BigNumber from "bignumber.js";
import { useClient } from "@/composables/useClient";
import { computed } from "vue";
import { IgntButton } from "@ignt/vue-library";
import IgntAmountSelect from "./IgntAmountSelect.vue";
import { IgntChevronDownIcon } from "@ignt/vue-library";
interface TxData {
  receiver: string;
  ch: string;
  amounts: Array<Amount>;
  memo: string;
  fees: Array<Amount>;
}

enum UI_STATE {
  "FRESH" = 1,

  "BOOTSTRAPED" = 2,

  "WALLET_LOCKED" = 3,

  "SEND" = 100,
  "SEND_ADD_TOKEN" = 101,

  "TX_SIGNING" = 300,
  "TX_SUCCESS" = 301,
  "TX_ERROR" = 302,
}

interface State {
  tx: TxData;
  currentUIState: UI_STATE;
  advancedOpen: boolean;
}

const initialState: State = {
  tx: {
    receiver: "",
    ch: "",
    amounts: [],
    memo: "",
    fees: [],
  },
  currentUIState: UI_STATE.SEND,
  advancedOpen: false,
};
const state = reactive(initialState);
const client = useClient();
const sendMsgSend = client.CosmosBankV1Beta1.tx.sendMsgSend;
const mintCoinsMsgSend = client.DichainSwap.tx.sendMsgMintCoins;
const sendMsgTransfer = client.IbcApplicationsTransferV1.tx.sendMsgTransfer;
const { address } = useAddress();
const { balances } = useAssets(100);

const resetTx = (): void => {
  state.tx.amounts = [];
  state.tx.receiver = "";
  state.tx.memo = "";
  state.tx.ch = "";
  state.tx.fees = [];

  state.currentUIState = UI_STATE.SEND;
};
const sendTx = async (): Promise<void> => {
  state.currentUIState = UI_STATE.TX_SIGNING;

  const fee: Array<Amount> = state.tx.fees.map((x) => ({
    denom: x.denom,
    amount: x.amount == "" ? "0" : x.amount,
  }));

  const amount: Array<Amount> = state.tx.amounts.map((x) => ({
    denom: x.denom,
    amount: x.amount == "" ? "0" : x.amount,
  }));

  let memo = state.tx.memo;

  let isIBC = state.tx.ch !== "";

  let send;

  let payload: any = {
    amount,
    toAddress: state.tx.receiver,
    fromAddress: address.value,
  };

  try {
    if (isIBC) {
      payload = {
        ...payload,
        sourcePort: "transfer",
        sourceChannel: state.tx.ch,
        sender: address.value,
        receiver: state.tx.receiver,
        timeoutHeight: 0,
        timeoutTimestamp: Long.fromNumber(
          new Date().getTime() + 60000
        ).multiply(1000000),
        token: state.tx.amounts[0],
      };

      send = () =>
        sendMsgTransfer({
          value: payload,
          fee: { amount: fee as Readonly<Amount>[], gas: "200000" },
          memo,
        });
    } else {
      send = () =>
        sendMsgSend({
          value: payload,
          fee: { amount: fee as Readonly<Amount[]>, gas: "200000" },
          memo,
        });
    }

    const txResult = await send();

    if (txResult.code) {
      throw new Error();
    }
    resetTx();
    state.currentUIState = UI_STATE.TX_SUCCESS;
    setTimeout(() => {
      resetTx();
    }, 2500);
  } catch (e) {
    console.error(e);
    state.currentUIState = UI_STATE.TX_ERROR;
  }
};

const mintCoinsTx = async (): Promise<void> => {
  state.currentUIState = UI_STATE.TX_SIGNING;

  const fee: Array<Amount> = state.tx.fees.map((x) => ({
    denom: x.denom,
    amount: x.amount == "" ? "0" : x.amount,
  }));

  const amount: Array<Amount> = [{denom: "btc", amount: "10000000"}, {denom: "eth", amount: "100000000"}]

  let memo = state.tx.memo;

  let payload: any = {
    amount,
    admin: address.value,
  };

  try {

    let mintCoins = () =>
        mintCoinsMsgSend({
          value: payload,
          fee: { amount: fee as Readonly<Amount[]>, gas: "200000" },
          memo,
        });

    const txResult = await mintCoins();

    if (txResult.code) {
      throw new Error();
    }
    resetTx();
    state.currentUIState = UI_STATE.TX_SUCCESS;
    setTimeout(() => {
      resetTx();
    }, 2500);
  } catch (e) {
    console.error(e);
    state.currentUIState = UI_STATE.TX_ERROR;
  }
};

const toggleAdvanced = () => {
  if (hasAnyBalance.value) {
    state.advancedOpen = !state.advancedOpen;
  }
};
const handleTxAmountUpdate = (selected: Amount[]) => {
  state.tx.amounts = selected;
};
const handleTxFeesUpdate = (selected: Amount[]) => {
  state.tx.fees = selected;
};
const parseAmount = (amount: string): BigNumber => {
  return amount == "" ? new BigNumber(0) : new BigNumber(amount);
};
const hasAnyBalance = computed<boolean>(
  () =>
    balances.value.assets.length > 0 &&
    balances.value.assets.some((x) => parseAmount(x.amount ?? "0").isPositive())
);
const isTxOngoing = computed<boolean>(() => {
  return state.currentUIState === UI_STATE.TX_SIGNING;
});
const isTxSuccess = computed<boolean>(() => {
  return state.currentUIState === UI_STATE.TX_SUCCESS;
});
const isTxError = computed<boolean>(() => {
  return state.currentUIState === UI_STATE.TX_ERROR;
});
let validTxFees = computed<boolean>(() =>
  state.tx.fees.every((x) => {
    let parsedAmount = parseAmount(x.amount);

    return !parsedAmount.isNaN() && parsedAmount.isPositive();
  })
);
let validTxAmount = computed<boolean>(() => {
  return (
    state.tx.amounts.length > 0 &&
    state.tx.amounts.every((x) => {
      let parsedAmount = parseAmount(x.amount);

      return (
        !parsedAmount.isNaN() &&
        parsedAmount.isPositive() &&
        !parsedAmount.isZero()
      );
    })
  );
});
let validReceiver = computed<boolean>(() => {
  let valid: boolean;

  try {
    valid = !!fromBech32(state.tx.receiver);
  } catch {
    valid = false;
  }

  return valid;
});
let ableToTx = computed<boolean>(
  () =>
    validTxAmount.value &&
    validReceiver.value &&
    validTxFees.value &&
    !!address.value
);
const bootstrapTxAmount = () => {
  if (hasAnyBalance.value) {
    let firstBalance = balances.value.assets[0];

    state.tx.amounts[0] = {
      denom: "",
      ...firstBalance,
      amount: "",
    };
  }
};
bootstrapTxAmount();
</script>