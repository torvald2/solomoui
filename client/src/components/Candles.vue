<template lang="pug">
b-container.mt-2
  b-row( align-v="center")
    b-col(cols="1" align-self="start")
      | Валюта
    b-col(cols="4" align-self="start")
      v-select(
        :options="paginated"
        @open="onOpen"
        @close="onClose"
        @search="(query) => (search = query)"
        :value="currentSymbol"
        @input="onSymbolSelect"
      )
        template(#list-footer)
          li(v-show="hasNextPage" ref="load" )
            | Загрузка ...
        template(#selected-option="opt")
          b-img-lazy(rounded :src="opt.image" width="20" height="20").mr-3
          |{{opt.symbol}}({{opt.label}})
        template(#option="opt")
          b-img-lazy(rounded :src="opt.image" width="20" height="20").mr-3
          |{{opt.symbol}}({{opt.label}})
    b-col(cols="1")
      | Период
    b-col(cols="1")
      b-form-select(:options="periods" :value="currentUnit" size="lg" @change="onPeriodChange")
  b-row.mt-2
    b-col
      TradingVue(:data="chart" :width="width" :height="500" :titleTxt="currentPair")

    </template>

</template>


<script>
import TradingVue from 'trading-vue-js'
import {mapState,mapMutations} from "vuex"

export default {
  name: "Candles",
  components:{TradingVue},
  data(){
    return {
      limit:10,
      search:"",
      observer: null,
      width:window.innerWidth- window.innerWidth/4,
      periods:[
        {value:"1m",text:" 1M  "},
        {value:"3m",text:" 3M  "},
        {value:"5m",text:" 5M  "},
        {value:"15m",text:" 15M  "},
        {value:"30m",text:" 30M  "},
        {value:"1h",text:" 1H  "},
        {value:"2h",text:" 2M  "},
        {value:"4h",text:" 4H  "},
        {value:"1h",text:" 12H  "},
        {value:"1D",text:" 1D  "},
        {value:"3D",text:" 3D  "},
        {value:"1w",text:" 1W  "}
      ]
    }
  },
  computed:{
    ...mapState(["currentSymbol","symbols","currentUnit","candles","currentPair"]),
    paginated() {
      return this.filtered.slice(0, this.limit)
    },
    filtered() {
      return this.symbols.filter((symbol) => symbol.symbol.includes(this.search))
    },
    hasNextPage() {
      return this.paginated.length < this.filtered.length
    },
    pairs(){
      return this.currentSymbol.pairs
    },
    chart(){
      let data =[]
      for (let c of this.candles){
        data.push([
          new Date(c.id.time).getTime(),
          c.open,
          c.high,
          c.low,
          c.close,
          0
        ])
      }
      return {
        
        chart:{
          type:"Candles",
          data: data,
          grid:{
            logScale:true,
          }
        }
      }
    }
  },
  methods:{
    ...mapMutations(["setCurrentSymbol","setCurrentUnit","setPair"]),
    async onOpen() {
      if (this.hasNextPage) {
        await this.$nextTick()
        this.observer.observe(this.$refs.load)
      }
    },
    onResize() {
            this.width = window.innerWidth- window.innerWidth/4
    },
    onPeriodChange(val){
      this.setCurrentUnit(val)
    },
    onPairChange(val){
      this.setPair(val)

    },
    onSymbolSelect(value){
      this.setCurrentSymbol(value)
    },
    onClose() {
      this.observer.disconnect()
    },
    async infiniteScroll([{ isIntersecting, target }]) {
      if (isIntersecting) {
        const ul = target.offsetParent
        const scrollTop = target.offsetParent.scrollTop
        this.limit += 10
        await this.$nextTick()
        ul.scrollTop = scrollTop
      }
    },
  },
  mounted(){
    this.observer = new IntersectionObserver(this.infiniteScroll)
    this.$nextTick(() => {
            window.addEventListener('resize', this.onResize);
        })
  },
  beforeDestroy(){
    window.removeEventListener('resize', this.onResize); 
  }
  
  
};
</script>

<style scoped>

</style>
