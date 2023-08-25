import{K as h,r as d,g as p,o as b,f as w,L as t,M as o,aT as y,b0 as P,a_ as _,b2 as Q,S as C,aS as $,b3 as U,c as A,k as I,O as z,X as B,R as F,a6 as N}from"./vendor-7e3b54f2.js";import{Q as Z,a as v}from"./QTable-f6f16d1c.js";import{U as M}from"./UiBox-438aefea.js";import{Q as D}from"./QForm-dd93a313.js";import{U as E}from"./UiInput-ff5536f7.js";import{u as T}from"./useSubmitForm-841c2b76.js";import{_ as g}from"./index-dc2fb6bd.js";import{t as L}from"./tableWrapper-c43f1a46.js";import{u as q}from"./useTableHandlers-649b36f2.js";import{u as H}from"./useUtils-85085a27.js";import{U as R}from"./UiRoundBtn-7ea8dd22.js";import"./QMenu-7e13e904.js";import"./QTooltip-cd4ef32e.js";import"./QMarkupTable-3806441a.js";import"./QChip-dc1c4db6.js";import"./QItemLabel-0efd9dc5.js";import"./QItemSection-cefe5148.js";import"./format-3895543d.js";import"./jwtClient-fdc6c2bb.js";import"./NetworkError-b28af479.js";import"./useConfirmDialog-7754cf44.js";const W=h({name:"FormZone",components:{UiInputVue:E},emits:["cancel","submit"],setup(e,{emit:n}){const r=d(null),{localProp:i,show:m,handleSubmit:u,handleCancel:l}=T(r,n);return{localProp:i,popup:r,show:m,handleSubmit:u,handleCancel:l}}});function G(e,n,r,i,m,u){const l=p("ui-input-vue");return b(),w(U,{ref:"popup","transition-show":"rotate","transition-hide":"rotate",persistent:""},{default:t(()=>[o($,{class:"min-width"},{default:t(()=>[o(y,null,{default:t(()=>[o(D,{class:"q-gutter-md",onSubmit:P(e.handleSubmit,["prevent"])},{default:t(()=>[o(l,{modelValue:e.localProp.name,"onUpdate:modelValue":n[0]||(n[0]=a=>e.localProp.name=a),hint:"Host Name *"},null,8,["modelValue"]),o(l,{modelValue:e.localProp.address,"onUpdate:modelValue":n[1]||(n[1]=a=>e.localProp.address=a),hint:"IP Address *",rule:"ip"},null,8,["modelValue"]),o(_,{modelValue:e.localProp.mac,"onUpdate:modelValue":n[2]||(n[2]=a=>e.localProp.mac=a),dense:"",outlined:"",hint:"MAC Address"},null,8,["modelValue"]),o(_,{modelValue:e.localProp.note,"onUpdate:modelValue":n[3]||(n[3]=a=>e.localProp.note=a),dense:"",outlined:"",hint:"Note"},null,8,["modelValue"]),o(Q,{align:"left",class:"text-primary"},{default:t(()=>[o(C,{label:"Submit",type:"submit",color:"primary"}),o(C,{label:"Cancel",color:"primary",flat:"",class:"q-ml-sm",onClick:e.handleCancel},null,8,["onClick"])]),_:1})]),_:1},8,["onSubmit"])]),_:1})]),_:1})]),_:1},512)}const K=g(W,[["render",G],["__scopeId","data-v-f0008912"]]),O="/proxy/zones";function X(e){return L(O,e)}const j=[{name:"state"},{name:"address",label:"IP address",align:"left",field:"address",sortable:!0},{name:"name",label:"Name",align:"left",field:"name"},{name:"mac",label:"MAC address",align:"left",field:"mac"},{name:"note",label:"Note",align:"left",field:"note"},{name:"actions"}],J=h({name:"PageZones",components:{UiBoxVue:M,FormZoneVue:K,UiRoundBtnVue:R},setup(){const e=d([]),n=X(e),r=d(null),i=d({}),m={xl:6,lg:7,md:7,sm:11,xs:10},u=A(()=>e.value.length===0),l=d(!1),{handleAdd:a,handleEdit:c,handleSubmit:s,handleDelete:f,handleCancel:V}=q(l,r,e,n,{}),{activeIcon:S,activeColor:k}=H();return I(async()=>{e.value=await n.Get(!0)}),{columns:j,rows:e,zone:i,buttonShow:u,form:r,boxCols:m,visible:l,activeIcon:S,activeColor:k,handleAdd:a,handleEdit:c,handleSubmit:s,handleDelete:f,handleCancel:V}}});function Y(e,n,r,i,m,u){const l=p("ui-round-btn-vue"),a=p("ui-box-vue"),c=p("form-zone-vue");return b(),z(F,null,[o(a,{columns:e.boxCols,header:"Local hosts",buttonShow:e.buttonShow,buttonLabel:"Add",buttonClick:e.handleAdd},{default:t(()=>[o(Z,{rows:e.rows,columns:e.columns,"row-key":"name","rows-per-page-options":[0,10,50,100]},{"body-cell-state":t(s=>[o(v,{props:s,class:"wd-30"},{default:t(()=>[o(N,{name:e.activeIcon(s.row.active),size:"1.2rem",color:e.activeColor(s.row.active)},null,8,["name","color"])]),_:2},1032,["props"])]),"body-cell-actions":t(s=>[o(v,{props:s},{default:t(()=>[o(l,{color:"primary",icon:"add",tooltip:"Create zone",onClick:e.handleAdd},null,8,["onClick"]),o(l,{color:"positive",icon:"mode_edit",tooltip:"Edit zone",onClick:f=>e.handleEdit(s.row)},null,8,["onClick"]),o(l,{color:"negative",icon:"delete",tooltip:"Delete zone",onClick:f=>e.handleDelete(s.row)},null,8,["onClick"])]),_:2},1032,["props"])]),_:1},8,["rows","columns"])]),_:1},8,["columns","buttonShow","buttonClick"]),e.visible?(b(),w(c,{key:0,ref:"form",onSubmit:e.handleSubmit,onCancel:e.handleCancel},null,8,["onSubmit","onCancel"])):B("",!0)],64)}const we=g(J,[["render",Y],["__scopeId","data-v-16f072c4"]]);export{we as default};