import{Q as n}from"./QTooltip-cd4ef32e.js";import{K as f,r as x,k as b,o as V,f as M,L as t,g as y,M as l,aT as C,b0 as g,aQ as m,a_ as s,Y as u,b2 as S,S as U,aS as h}from"./vendor-7e3b54f2.js";import{Q}from"./QMarkupTable-3806441a.js";import{Q as F}from"./QForm-dd93a313.js";import{U as w}from"./UiBox-438aefea.js";import{u as B}from"./useConfirmDialog-7754cf44.js";import{j as i}from"./jwtClient-fdc6c2bb.js";import{_ as k}from"./index-dc2fb6bd.js";import"./QItemLabel-0efd9dc5.js";import"./NetworkError-b28af479.js";async function N(){const e={};return i.get("/esp32/settings").then(({success:a,result:d})=>a?d:e).catch(()=>e)}function c(e){e.max_6814_nh3=parseFloat(e.max_6814_nh3),e.max_6814_no2=parseFloat(e.max_6814_no2),e.max_6814_co=parseFloat(e.max_6814_co),e.min_temp=parseFloat(e.min_temp),e.max_temp=parseFloat(e.max_temp),e.min_bmx280_tempr=parseFloat(e.min_bmx280_tempr),e.max_bmx280_tempr=parseFloat(e.max_bmx280_tempr),e.max_rad_stat=parseFloat(e.max_rad_stat),e.max_rad_dyn=parseFloat(e.max_rad_dyn),e.max_ch2o=parseFloat(e.max_ch2o),i.put("/esp32/settings",e)}const v=f({name:"PageEsp32Alarm",components:{UiBoxVue:w},setup(){const e=x({}),a=B(),d={xl:6,lg:6,md:7,sm:11,xs:10},r=async()=>e.value=await N();return b(()=>r()),{settings:e,boxCols:d,refresh:r,async handleSubmit(){await a.show("Update settings?")&&c(e.value)}}}}),R={class:"text-left"},A={class:"text-left"},H={class:"text-left"},O={class:"text-left"},E={class:"text-left"},$={class:"text-left"},D={class:"text-left"},I={class:"text-left"},T={class:"text-left"},j={class:"text-left"};function L(e,a,d,r,p,P){const _=y("ui-box-vue");return V(),M(_,{columns:e.boxCols,header:"Alarm Setting",buttonShow:!0,buttonLabel:"Refresh",buttonClick:e.refresh},{default:t(()=>[l(h,null,{default:t(()=>[l(C,null,{default:t(()=>[l(F,{class:"q-gutter-md",onSubmit:a[10]||(a[10]=g(o=>e.handleSubmit(),["prevent"]))},{default:t(()=>[l(Q,{dense:"","wrap-cells":""},{default:t(()=>[m("tbody",null,[m("tr",null,[m("td",R,[l(s,{modelValue:e.settings.min_bmx280_tempr,"onUpdate:modelValue":a[0]||(a[0]=o=>e.settings.min_bmx280_tempr=o),dense:"",label:"BME280/Min temperatuer (°C)",type:"number"},{default:t(()=>[l(n,null,{default:t(()=>[u("Minimum alarm temperature")]),_:1})]),_:1},8,["modelValue"])])]),m("tr",null,[m("td",A,[l(s,{modelValue:e.settings.max_bmx280_tempr,"onUpdate:modelValue":a[1]||(a[1]=o=>e.settings.max_bmx280_tempr=o),dense:"",label:"BME280/Max temperatuer (°C)",type:"number"},{default:t(()=>[l(n,null,{default:t(()=>[u("Maximum alarm temperature")]),_:1})]),_:1},8,["modelValue"])])]),m("tr",null,[m("td",H,[l(s,{modelValue:e.settings.max_6814_no2,"onUpdate:modelValue":a[2]||(a[2]=o=>e.settings.max_6814_no2=o),style:{"min-width":"100px"},dense:"",label:"MICS6814/Max NO2 (mg/m3)",type:"number"},{default:t(()=>[l(n,null,{default:t(()=>[u("Maximum alarm NO2")]),_:1})]),_:1},8,["modelValue"])])]),m("tr",null,[m("td",O,[l(s,{modelValue:e.settings.max_6814_nh3,"onUpdate:modelValue":a[3]||(a[3]=o=>e.settings.max_6814_nh3=o),dense:"",label:"MICS6814/Max NH3 (mg/m3)",type:"number"},{default:t(()=>[l(n,null,{default:t(()=>[u("Maximum alarm NH3")]),_:1})]),_:1},8,["modelValue"])])]),m("tr",null,[m("td",E,[l(s,{modelValue:e.settings.max_6814_co,"onUpdate:modelValue":a[4]||(a[4]=o=>e.settings.max_6814_co=o),dense:"",label:"MICS6814/Max CO (mg/m3)",type:"number"},{default:t(()=>[l(n,null,{default:t(()=>[u("Maximum alarm NH3")]),_:1})]),_:1},8,["modelValue"])])]),m("tr",null,[m("td",$,[l(s,{modelValue:e.settings.max_rad_stat,"onUpdate:modelValue":a[5]||(a[5]=o=>e.settings.max_rad_stat=o),dense:"",label:"RadSens/Max Static Radiation (µR/h)",type:"number"},{default:t(()=>[l(n,null,{default:t(()=>[u("Maximum static radiation")]),_:1})]),_:1},8,["modelValue"])])]),m("tr",null,[m("td",D,[l(s,{modelValue:e.settings.max_rad_dyn,"onUpdate:modelValue":a[6]||(a[6]=o=>e.settings.max_rad_dyn=o),dense:"",label:"RadSens/Max Dynamic Radiation (µR/h)",type:"number"},{default:t(()=>[l(n,null,{default:t(()=>[u("Maximum dynamic radiation")]),_:1})]),_:1},8,["modelValue"])])]),m("tr",null,[m("td",I,[l(s,{modelValue:e.settings.max_ch2o,"onUpdate:modelValue":a[7]||(a[7]=o=>e.settings.max_ch2o=o),dense:"",label:"ZE08CH2O/Max CH2O (ppm)",type:"number"},{default:t(()=>[l(n,null,{default:t(()=>[u("Maximum alarm CH2O")]),_:1})]),_:1},8,["modelValue"])])]),m("tr",null,[m("td",T,[l(s,{modelValue:e.settings.min_ds18b20,"onUpdate:modelValue":a[8]||(a[8]=o=>e.settings.min_ds18b20=o),dense:"",label:"DS18B20/Min temperatuer (°C)",type:"number"},{default:t(()=>[l(n,null,{default:t(()=>[u("Minimum alarm temperature")]),_:1})]),_:1},8,["modelValue"])])]),m("tr",null,[m("td",j,[l(s,{modelValue:e.settings.max_ds18b20,"onUpdate:modelValue":a[9]||(a[9]=o=>e.settings.max_ds18b20=o),dense:"",label:"DS18B20/Min temperatuer (°C)",type:"number"},{default:t(()=>[l(n,null,{default:t(()=>[u("Minimum alarm temperature")]),_:1})]),_:1},8,["modelValue"])])])])]),_:1}),l(S,{align:"left",class:"text-primary"},{default:t(()=>[l(U,{label:"Submit",type:"submit",color:"primary "})]),_:1})]),_:1})]),_:1})]),_:1})]),_:1},8,["columns","buttonClick"])}const ae=k(v,[["render",L]]);export{ae as default};