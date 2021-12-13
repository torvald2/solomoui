export async function GetSymbols(){
    const resp = await fetch("/api/symbols",{
        method:"GET",
    })
    return await resp.json()
}


export async function GetCandles(symbol,pair, unit){
    const resp = await fetch(`/api/candles?symbol=${symbol}&unit=${unit}&pair=${pair}`,{
        method:"GET",
    })
    return await resp.json()
}
