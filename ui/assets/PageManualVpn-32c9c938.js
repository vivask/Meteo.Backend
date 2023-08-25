import{b as S,Q as g,a as y}from"./QTable-f6f16d1c.js";import{U as k}from"./UiBox-438aefea.js";import{K as C,r as u,k as _,g as d,o as f,f as v,L as s,M as e,aT as P,b0 as Q,a_ as $,b2 as U,S as h,aS as A,b3 as B,c as F,O as H,X as L,R as M}from"./vendor-7e3b54f2.js";import{Q as N}from"./QForm-dd93a313.js";import{U as D}from"./UiInput-ff5536f7.js";import{u as E}from"./useSubmitForm-841c2b76.js";import{j as I}from"./jwtClient-fdc6c2bb.js";import{_ as V}from"./index-dc2fb6bd.js";import{t as T}from"./tableWrapper-c43f1a46.js";import{u as R}from"./useTableHandlers-649b36f2.js";import{U as j}from"./UiRoundBtn-7ea8dd22.js";import"./QMenu-7e13e904.js";import"./QTooltip-cd4ef32e.js";import"./QMarkupTable-3806441a.js";import"./QChip-dc1c4db6.js";import"./QItemLabel-0efd9dc5.js";import"./QItemSection-cefe5148.js";import"./format-3895543d.js";import"./NetworkError-b28af479.js";import"./useConfirmDialog-7754cf44.js";async function q(){const o=[];return I.get("/proxy/vpnlists").then(({success:t,result:a})=>t?a:o).catch(()=>o)}const W=C({name:"FormVpnHost",components:{UiInputVue:D},emits:["cancel","submit"],setup(o,{emit:t}){const a=u([]),r=u(null),{localProp:m,show:i,handleSubmit:l,handleCancel:n}=E(r,t);return _(async()=>{a.value=await q()}),{localProp:m,localList:a,popup:r,show:i,handleSubmit:l,handleCancel:n}}});function z(o,t,a,r,m,i){const l=d("ui-input-vue");return f(),v(B,{ref:"popup","transition-show":"rotate","transition-hide":"rotate",persistent:""},{default:s(()=>[e(A,{class:"min-width"},{default:s(()=>[e(P,null,{default:s(()=>[e(N,{class:"q-gutter-md",onSubmit:Q(o.handleSubmit,["prevent"])},{default:s(()=>[e(l,{modelValue:o.localProp.name,"onUpdate:modelValue":t[0]||(t[0]=n=>o.localProp.name=n),hint:"Name/IP Address *"},null,8,["modelValue"]),e(S,{modelValue:o.localProp.list,"onUpdate:modelValue":t[1]||(t[1]=n=>o.localProp.list=n),outlined:"",dense:"",options:o.localList,"option-label":"id",hint:"Acess list *","lazy-rules":"",rules:[n=>n||"Please select something"]},null,8,["modelValue","options","rules"]),e($,{modelValue:o.localProp.note,"onUpdate:modelValue":t[2]||(t[2]=n=>o.localProp.note=n),dense:"",outlined:"",hint:"Note"},null,8,["modelValue"]),e(U,{align:"left",class:"text-primary"},{default:s(()=>[e(h,{label:"Submit",type:"submit",color:"primary"}),e(h,{label:"Cancel",color:"primary",flat:"",class:"q-ml-sm",onClick:o.handleCancel},null,8,["onClick"])]),_:1})]),_:1},8,["onSubmit"])]),_:1})]),_:1})]),_:1},512)}const G=V(W,[["render",z],["__scopeId","data-v-955eb2b7"]]),K="/proxy/manualvpn";function O(o){return T(K,o)}const X=[{name:"name",align:"left",field:"name",sortable:!0},{name:"vpnlist"},{name:"note",align:"left",field:"note",sortable:!0},{name:"actions"}],J=C({name:"PageManualVpn",components:{UiBoxVue:k,UiRoundBtnVue:j,FormVpnHostVue:G},setup(){const o=u([]),t=O(o),a=u({}),r={xl:6,lg:6,md:7,sm:11,xs:10},m=F(()=>o.value.length===0),i=u(null),l=u(!1),{handleAdd:n,handleEdit:c,handleSubmit:p,handleDelete:b,handleCancel:w}=R(l,i,o,t,{});return _(async()=>{o.value=await t.Get()}),{columns:X,rows:o,host:a,buttonShow:m,form:i,boxCols:r,visible:l,handleAdd:n,handleEdit:c,handleSubmit:p,handleDelete:b,handleCancel:w}},methods:{}});function Y(o,t,a,r,m,i){const l=d("ui-round-btn-vue"),n=d("ui-box-vue"),c=d("form-vpn-host-vue");return f(),H(M,null,[e(n,{columns:o.boxCols,header:"Hosts Redirected to VPN",buttonShow:o.buttonShow,buttonLabel:"Add",buttonClick:o.handleAdd},{default:s(()=>[e(g,{"hide-header":"",rows:o.rows,columns:o.columns,"row-key":"name","rows-per-page-options":[0,10,50,100]},{"body-cell-actions":s(p=>[e(y,{props:p},{default:s(()=>[e(l,{color:"primary",icon:"add",tooltip:"Create host",onClick:o.handleAdd},null,8,["onClick"]),e(l,{color:"positive",icon:"mode_edit",tooltip:"Edit host",onClick:b=>o.handleEdit(p.row)},null,8,["onClick"]),e(l,{color:"negative",icon:"delete",tooltip:"Delete host",onClick:b=>o.handleDelete(p.row)},null,8,["onClick"])]),_:2},1032,["props"])]),_:1},8,["rows","columns"])]),_:1},8,["columns","buttonShow","buttonClick"]),o.visible?(f(),v(c,{key:0,ref:"form",onSubmit:o.handleSubmit,onCancel:o.handleCancel},null,8,["onSubmit","onCancel"])):L("",!0)],64)}const vo=V(J,[["render",Y]]);export{vo as default};