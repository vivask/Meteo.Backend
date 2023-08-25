import{C as l,u as c}from"./useChartWrapper-df43fe4d.js";import{w as u}from"./webClient-a78b5567.js";import{_ as i}from"./index-dc2fb6bd.js";import{K as d,r as h,c as f,o as b,f as v,g as C}from"./vendor-7e3b54f2.js";import"./UiSquareBtn-3d095194.js";import"./QTooltip-cd4ef32e.js";import"./useUtils-85085a27.js";import"./format-3895543d.js";import"./NetworkError-b28af479.js";async function _(e,r,o,s){const a=[],n=`/esp32/bmx280/${e}/${r}`;return u.post(n,{begin:o,end:s}).then(({success:t,result:p})=>t?p:a).catch(()=>a)}const $=d({name:"PageBme280",components:{ChartBoxVue:l},props:{parameter:{type:String}},setup(e){const r=h([]),{chartLabel:o,chartPeriod:s,chartLabels:a}=c(e.parameter,r,_),n=f(()=>{const t=e.parameter==="pressure"?"press":e.parameter==="temperature"?"tempr":"hum";return r.value.map(m=>t==="press"?m[t]/133:m[t])});return{chartLabel:o,chartPeriod:s,chartData:n,chartLabels:a,data:r}}});function g(e,r,o,s,a,n){const t=C("chart-box-vue");return b(),v(t,{modelValue:e.chartPeriod,"onUpdate:modelValue":r[0]||(r[0]=p=>e.chartPeriod=p),label:e.chartLabel,values:e.chartData,labels:e.chartLabels},null,8,["modelValue","label","values","labels"])}const K=i($,[["render",g]]);export{K as default};