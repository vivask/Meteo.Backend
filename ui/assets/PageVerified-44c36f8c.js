import{a as p,Q as d}from"./QTable-f6f16d1c.js";import{U as f}from"./UiBox-438aefea.js";import{u as _}from"./useUtils-85085a27.js";import{U as b}from"./UiRoundBtn-7ea8dd22.js";import{u as w}from"./useConfirmDialog-7754cf44.js";import{t as g}from"./tableWrapper-c43f1a46.js";import{_ as v}from"./index-dc2fb6bd.js";import{K as h,r as k,k as x,o as C,f as y,L as i,g as c,M as l}from"./vendor-7e3b54f2.js";import"./QMenu-7e13e904.js";import"./QTooltip-cd4ef32e.js";import"./QMarkupTable-3806441a.js";import"./QChip-dc1c4db6.js";import"./QItemLabel-0efd9dc5.js";import"./QItemSection-cefe5148.js";import"./format-3895543d.js";import"./jwtClient-fdc6c2bb.js";import"./NetworkError-b28af479.js";const U="/radius/verified";function V(e){return g(U,e)}const $=h({name:"PageVerified",components:{UiBoxVue:f,UiRoundBtnVue:b},setup(){const e=[{name:"username",label:"User",align:"left",field:"username",sortable:!0},{name:"callingstationid",label:"User Id",align:"left",field:"callingstationid",sortable:!0},{name:"acctupdatetime",label:"Last used",align:"left",field:"acctupdatetime",sortable:!0,format:(t,a)=>s(t)},{name:"actions"}],o=k([]),n=V(o),u=w(),m={xl:6,lg:6,md:7,sm:11,xs:10},{formatLongDate:s}=_(),r=async()=>o.value=await n.Get();return x(()=>r()),{columns:e,rows:o,boxCols:m,formatLongDate:s,refresh:r,async handleExclude(t){await u.show("Are you sure to exclude this item?")&&(o.value=await n.Delete(t))}}}});function B(e,o,n,u,m,s){const r=c("ui-round-btn-vue"),t=c("ui-box-vue");return C(),y(t,{columns:e.boxCols,header:"Verified users",buttonShow:!0,buttonLabel:"Refresh",buttonClick:e.refresh},{default:i(()=>[l(d,{rows:e.rows,columns:e.columns,"row-key":"name","rows-per-page-options":[0,10,50,100]},{"body-cell-actions":i(a=>[l(p,{props:a},{default:i(()=>[l(r,{color:"warning",icon:"mdi-minus-thick",tooltip:"Exclude user from verified",onClick:L=>e.handleExclude(a.row)},null,8,["onClick"])]),_:2},1032,["props"])]),_:1},8,["rows","columns"])]),_:1},8,["columns","buttonClick"])}const F=v($,[["render",B],["__scopeId","data-v-b1552a4e"]]);export{F as default};
