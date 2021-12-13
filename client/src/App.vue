<template lang="pug">
  article
    b-navbar(toggleable="lg" type="dark" variant="info")
      b-navbar-brand Solamon
    router-view
</template>

<script>

import {mapActions} from "vuex"
export default {
  name: "App",
  data(){
    return {
      candleTimer:false,
      symbolTimer:false
    }
  },
  methods: {
    ...mapActions(["getCandles","getSymbols"])

  },
  created(){
      this.getSymbols()
      this.candleTimer = setInterval(this.getSymbols, 60000)
      this.symbolTimer = setInterval(this.getCandles,  1000)
    
  },
  beforeDestroy() {
    clearInterval(this.candleTimer)
    clearInterval(this.symbolTimer)
  }
};
</script>


<style>

</style>
