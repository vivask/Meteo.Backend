import{Q as _}from"./QTooltip-cd4ef32e.js";import{K as v,r as h,c as y,k as C,w,o as k,f as U,L as r,g as V,M as u,S as x,Y as S}from"./vendor-7e3b54f2.js";import{a as A,Q}from"./QTable-f6f16d1c.js";import{U as $}from"./UiBox-438aefea.js";import{u as B}from"./useUtils-85085a27.js";import{u as L}from"./useConfirmDialog-7754cf44.js";import{_ as P,u as T}from"./index-dc2fb6bd.js";import{j as d}from"./jwtClient-fdc6c2bb.js";import"./QMenu-7e13e904.js";import"./QMarkupTable-3806441a.js";import"./QChip-dc1c4db6.js";import"./QItemLabel-0efd9dc5.js";import"./QItemSection-cefe5148.js";import"./format-3895543d.js";import"./NetworkError-b28af479.js";async function I(t){const a=[],o="/radius/accounting/"+t;return d.post(o,{offset:0,limit:500}).then(({success:c,result:i})=>c?i:a).catch(()=>a)}async function j(t){const a="/radius/accounting/verified/"+t;return d.put(a).then(({success:o})=>o).catch(()=>!1)}async function D(t){const a="/radius/accounting/clear/"+t;return d.put(a).then(({success:o})=>o).catch(()=>!1)}const M=v({name:"PageAccounting",components:{UiBoxVue:$},setup(){const t=[{name:"state"},{name:"username",label:"User",align:"left",field:"username",sortable:!0},{name:"callingstationid",label:"User Id",align:"left",field:"callingstationid",sortable:!0},{name:"acctstarttime",label:"Start time",align:"left",field:"acctstarttime",sortable:!0,format:(e,m)=>i(e)},{name:"nasportid",label:"AP",align:"left",field:"nasportid",sortable:!0},{name:"acctupdatetime",label:"Update time",align:"left",field:"acctupdatetime",sortable:!0,format:(e,m)=>i(e)},{name:"acctstoptime",label:"Stop time",align:"left",field:"acctstoptime",sortable:!0,format:(e,m)=>i(e)},{name:"calledstationid",label:"Station Id",align:"left",field:"calledstationid",sortable:!0}],a=h([]),o=L(),c={xl:9,lg:9,md:7,sm:11,xs:10},{formatLongDate:i}=B(),f=T(),s=y(()=>f.usersFilter.value),n=e=>e.verified&&e.verified.length>0&&e.valid&&e.valid.length>0,p=e=>n(e)&&e.valid===e.username,g=e=>p(e)?"positive":n(e)?"warning":"negative",b=e=>e.verified.length>0?"verified_user":"mdi-shield-account",l=async()=>a.value=await I(s.value);return C(()=>l()),w(s,()=>l()),{columns:t,rows:a,confirm:o,boxCols:c,formatLongDate:i,refresh:l,color:g,icon:b,async handleVerify(e){n(e)||await j(e.id)&&l()},async handleClear(){await D(s.value)&&l()}}}});function N(t,a,o,c,i,f){const s=V("ui-box-vue");return k(),U(s,{columns:t.boxCols,header:"Accounting Radius",buttonShow:!0,buttonLabel:"Clear",buttonClick:t.handleClear},{default:r(()=>[u(Q,{rows:t.rows,columns:t.columns,"row-key":"name","rows-per-page-options":[10,50,100,0]},{"body-cell-state":r(n=>[u(A,{props:n,class:"wd-60"},{default:r(()=>[u(x,{class:"q-ml-xs",dense:"",round:"",color:t.color(n.row),size:"md",icon:t.icon(n.row),onClick:p=>t.handleVerify(n.row)},{default:r(()=>[u(_,null,{default:r(()=>[S("Confirm verify user")]),_:1})]),_:2},1032,["color","icon","onClick"])]),_:2},1032,["props"])]),_:1},8,["rows","columns"])]),_:1},8,["columns","buttonClick"])}const te=P(M,[["render",N],["__scopeId","data-v-bd2ba66c"]]);export{te as default};
