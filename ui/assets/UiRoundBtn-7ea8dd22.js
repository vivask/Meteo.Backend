import{Q as i}from"./QTooltip-cd4ef32e.js";import{K as n,c as a,o as t,f as e,L as r,Y as p,a7 as l,X as c,aZ as m,S as d}from"./vendor-7e3b54f2.js";import{_ as u}from"./index-dc2fb6bd.js";const f=n({name:"UiRoundBtn",inheritAttrs:!1,props:{color:{type:String,default:"primary"},icon:String,tooltip:String},setup(o){return{showTooltiip:a(()=>!!o.tooltip&&o.tooltip.length>0)}}});function h(o,s,g,_,B,S){return t(),e(d,m({class:"q-ml-xs",dense:"",round:"",size:"md",color:o.color,icon:o.icon},o.$attrs),{default:r(()=>[o.showTooltiip?(t(),e(i,{key:0},{default:r(()=>[p(l(o.tooltip),1)]),_:1})):c("",!0)]),_:1},16,["color","icon"])}const y=u(f,[["render",h]]);export{y as U};