import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);
import {GetCandles, GetSymbols} from "../services/api"
export default new Vuex.Store({
  state: {
    symbols:[],
    currentSymbol:{
      "label": "Mercurial",
      "symbol": "MER",
      "image": "https://raw.githubusercontent.com/solana-labs/token-list/main/assets/mainnet/MERt85fc5boKw3BW1eYdxonEuJNvXbiMbs6hvheau5K/logo.png",
      "id": "61af8d4d063abae6d2eec0e3",
      "pairs": [
          "MER/USD"
      ]
  },
    currentUnit:"1m",
    currentPair:"MER/USD",
    candles:[]
  },
  mutations: {
    setSymbols(state, data){
      state.symbols = data
    },
    setCurrentSymbol(state,data){
      state.currentSymbol = data
      state.currentPair = data.pairs[0]
    },
    setCurrentUnit(state,data){
      state.currentUnit = data
    },
    setCandles(state,data){
      state.candles = data
    },
    setPair(state,data){
      state.currentPair = data
    }
  },
  actions: {
    async getCandles(context){
      if (context.state.currentSymbol){
        context.commit("setCandles", await GetCandles(context.state.currentSymbol.id,context.state.currentPair, context.state.currentUnit))
      }
    },
    async getSymbols(context){
      context.commit("setSymbols", await GetSymbols())
    },

  },
});
