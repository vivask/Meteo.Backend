import{K as S,r as p,g as f,o as h,f as g,L as a,M as o,aT as M,b0 as U,b2 as P,S as w,aS as Y,b3 as B,c as K,k as A,O as v,X as C,R as H,a6 as k,aQ as y,a7 as d}from"./vendor-7e3b54f2.js";import{Q as L,a as c}from"./QTable-f6f16d1c.js";import{U as T}from"./UiBox-438aefea.js";import{U as x}from"./UiRoundBtn-7ea8dd22.js";import{u as F}from"./useTableHandlers-649b36f2.js";import{Q as I}from"./QForm-dd93a313.js";import{U as N}from"./UiInput-ff5536f7.js";import{u as E}from"./useSubmitForm-841c2b76.js";import{_ as V}from"./index-dc2fb6bd.js";import{u as q}from"./useUtils-85085a27.js";import{t as z}from"./tableWrapper-c43f1a46.js";import"./QMenu-7e13e904.js";import"./QTooltip-cd4ef32e.js";import"./QMarkupTable-3806441a.js";import"./QChip-dc1c4db6.js";import"./QItemLabel-0efd9dc5.js";import"./QItemSection-cefe5148.js";import"./format-3895543d.js";import"./useConfirmDialog-7754cf44.js";import"./jwtClient-fdc6c2bb.js";import"./NetworkError-b28af479.js";const R=S({name:"FormSshKey",components:{UiInputVue:N},emits:["cancel","submit"],setup(e,{emit:n}){const l=p(null),{localProp:m,show:u,handleSubmit:i,handleCancel:s}=E(l,n);return{localProp:m,popup:l,show:u,handleSubmit:i,handleCancel:s}}});function W(e,n,l,m,u,i){const s=f("ui-input-vue");return h(),g(B,{ref:"popup","transition-show":"rotate","transition-hide":"rotate",persistent:""},{default:a(()=>[o(Y,{style:{"min-width":"350px"}},{default:a(()=>[o(M,null,{default:a(()=>[o(I,{class:"q-gutter-md",onSubmit:U(e.handleSubmit,["prevent"])},{default:a(()=>[o(s,{modelValue:e.localProp.owner,"onUpdate:modelValue":n[0]||(n[0]=r=>e.localProp.owner=r),hint:"Key Name *"},null,8,["modelValue"]),o(s,{modelValue:e.localProp.finger,"onUpdate:modelValue":n[1]||(n[1]=r=>e.localProp.finger=r),hint:"Content *",type:"textarea"},null,8,["modelValue"]),o(P,{align:"left",class:"text-primary"},{default:a(()=>[o(w,{label:"Submit",type:"submit",color:"primary"}),o(w,{label:"Cancel",color:"primary",flat:"",class:"q-ml-sm",onClick:e.handleCancel},null,8,["onClick"])]),_:1})]),_:1},8,["onSubmit"])]),_:1})]),_:1})]),_:1},512)}const G=V(R,[["render",W]]),O="/sshclient/sshkeys";function X(e){return z(O,e)}const j=[{name:"state"},{name:"icon"},{name:"key"},{name:"actions"}],J=S({name:"PageSshKeys",components:{UiBoxVue:T,UiRoundBtnVue:x,FormSshKeyVue:G},setup(){const e=p([]),n=X(e),l=p(null),m={xl:5,lg:5,md:7,sm:11,xs:10},u=K(()=>e.value.length===0),i=p(!1),{handleAdd:s,handleSubmit:r,handleDelete:b,handleCancel:t}=F(i,l,e,n,{}),{formatLongDate:_,isEmptyTime:D,activeIcon:Q,activeColor:$}=q();return A(async()=>e.value=await n.Get(!0)),{columns:j,rows:e,buttonShow:u,wrapper:n,form:l,boxCols:m,visible:i,formatLongDate:_,isEmptyTime:D,activeIcon:Q,activeColor:$,handleAdd:s,handleSubmit:r,handleDelete:b,handleCancel:t}}}),Z={class:"text-subtitle1 text-bold text-left text-primary"},ee={class:"text-subtitle2","text-left":""},oe={class:"text-meta text-left"},te={key:0,class:"text-meta text-left"};function ae(e,n,l,m,u,i){const s=f("ui-round-btn-vue"),r=f("ui-box-vue"),b=f("form-ssh-key-vue");return h(),v(H,null,[o(r,{columns:e.boxCols,header:"SSH key management",buttonShow:e.buttonShow,buttonLabel:"Add",buttonClick:e.handleAdd,tooltip:"Add new ssh key"},{default:a(()=>[o(L,{"hide-header":"",rows:e.rows,columns:e.columns,"row-key":"id","rows-per-page-options":[10,50,100,0]},{"body-cell-state":a(t=>[o(c,{props:t,class:"wd-20"},{default:a(()=>[o(k,{name:e.activeIcon(t.row.activity),size:"1.2rem",color:e.activeColor(t.row.activity)},null,8,["name","color"])]),_:2},1032,["props"])]),"body-cell-icon":a(t=>[o(c,{props:t,class:"wd-80"},{default:a(()=>[o(k,{name:"mdi-key-variant",size:"md"})]),_:2},1032,["props"])]),"body-cell-key":a(t=>[o(c,{props:t,class:"wd-100"},{default:a(()=>[y("div",Z,d(t.row.owner),1),y("div",ee,d(t.row.short_finger),1),y("div",oe," Created "+d(e.formatLongDate(t.row.created,"MMM DD, YYYY HH:mm:ss")),1),e.isEmptyTime(t.row.used)?C("",!0):(h(),v("div",te," Last used "+d(e.formatLongDate(t.row.used,"MMM DD, YYYY HH:mm:ss")),1))]),_:2},1032,["props"])]),"body-cell-actions":a(t=>[o(c,{props:t},{default:a(()=>[o(s,{color:"primary",icon:"add",tooltip:"Create ssh key",onClick:e.handleAdd},null,8,["onClick"]),o(s,{color:"negative",icon:"delete",tooltip:"Delete ssh key",onClick:_=>e.handleDelete(t.row)},null,8,["onClick"])]),_:2},1032,["props"])]),_:1},8,["rows","columns"])]),_:1},8,["columns","buttonShow","buttonClick"]),e.visible?(h(),g(b,{key:0,ref:"form",onSubmit:e.handleSubmit,onCancel:e.handleCancel},null,8,["onSubmit","onCancel"])):C("",!0)],64)}const Ve=V(J,[["render",ae],["__scopeId","data-v-52efafa8"]]);export{Ve as default};