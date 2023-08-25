import{K as _,r as u,N as g,g as f,o as b,f as w,L as l,M as o,aT as y,b0 as V,b2 as k,S as h,aS as P,b3 as U,c as $,k as A,O as Q,X as B,R,a6 as T}from"./vendor-7e3b54f2.js";import{Q as E,a as v}from"./QTable-f6f16d1c.js";import{U as F}from"./UiBox-438aefea.js";import{Q as I}from"./QForm-dd93a313.js";import{U as N}from"./UiInput-ff5536f7.js";import{U as D}from"./UiPasswordInput-8bf8e294.js";import{u as M}from"./useSubmitForm-841c2b76.js";import{_ as C}from"./index-dc2fb6bd.js";import{t as O}from"./tableWrapper-c43f1a46.js";import{u as q}from"./useTableHandlers-649b36f2.js";import{U as z}from"./UiRoundBtn-7ea8dd22.js";import"./QMenu-7e13e904.js";import"./QTooltip-cd4ef32e.js";import"./QMarkupTable-3806441a.js";import"./QChip-dc1c4db6.js";import"./QItemLabel-0efd9dc5.js";import"./QItemSection-cefe5148.js";import"./format-3895543d.js";import"./jwtClient-fdc6c2bb.js";import"./NetworkError-b28af479.js";import"./useConfirmDialog-7754cf44.js";const L=_({name:"FormAuth",components:{UiInputVue:N,UiPasswordInputVue:D},emits:["cancel","submit"],setup(e,{emit:t}){const i=u(null),m=u(null),p=u(null),c=u(!0),{localProp:a,show:r,handleSubmit:n,handleCancel:s}=M(i,t);return{localProp:a,popup:i,confirm:m,input:p,isPwd:c,show(d){m.value=d.value,r(d)},handleSubmit:n,handleCancel:s,onSubmit(){a.value.value!=m.value?(g.create({timeout:{}.ERROR_TIMEOUT,type:"negative",message:"Passwords do not match!"}),p.value.focus()):(a.value.attribute="Cleartext-Password",a.value.op=":=",n())}}}});function W(e,t,i,m,p,c){const a=f("ui-input-vue"),r=f("ui-password-input-vue");return b(),w(U,{ref:"popup","transition-show":"rotate","transition-hide":"rotate",persistent:""},{default:l(()=>[o(P,{style:{"min-width":"350px"}},{default:l(()=>[o(y,null,{default:l(()=>[o(I,{class:"q-gutter-md",autocomplete:"off",onSubmit:V(e.onSubmit,["prevent"])},{default:l(()=>[o(a,{modelValue:e.localProp.username,"onUpdate:modelValue":t[0]||(t[0]=n=>e.localProp.username=n),hint:"User Name *",autocomplete:"off"},null,8,["modelValue"]),o(r,{modelValue:e.localProp.value,"onUpdate:modelValue":t[1]||(t[1]=n=>e.localProp.value=n),hint:"Password *",autocomplete:"off"},null,8,["modelValue"]),o(r,{ref:"input",modelValue:e.confirm,"onUpdate:modelValue":t[2]||(t[2]=n=>e.confirm=n),hint:"Confirm password *"},null,8,["modelValue"]),o(k,{align:"left",class:"text-primary"},{default:l(()=>[o(h,{label:"Submit",type:"submit",color:"primary"}),o(h,{label:"Cancel",color:"primary",flat:"",class:"q-ml-sm",onClick:e.handleCancel},null,8,["onClick"])]),_:1})]),_:1},8,["onSubmit"])]),_:1})]),_:1})]),_:1},512)}const G=C(L,[["render",W]]),H="/radius/auth";function K(e){return O(H,e)}const X=[{name:"state"},{name:"username",align:"left",field:"username",sortable:!0},{name:"attribute",align:"left",field:"attribute"},{name:"op",align:"left",field:"op"},{name:"actions"}],j=_({name:"PageAuth",components:{UiBoxVue:F,FormAuthVue:G,UiRoundBtnVue:z},setup(){const e=u([]),t=K(e),i=u(null),m={xl:6,lg:6,md:7,sm:11,xs:10},p=$(()=>e.value.length===0),c=u(!1),a=u({}),{handleAdd:r,handleEdit:n,handleSubmit:s,handleDelete:d,handleCancel:S}=q(c,i,e,t,{});return A(async()=>e.value=await t.Get(!0)),{columns:X,rows:e,buttonShow:p,wrapper:t,form:i,boxCols:m,visible:c,user:a,handleAdd:r,handleEdit:n,handleSubmit:s,handleDelete:d,handleCancel:S}}});function J(e,t,i,m,p,c){const a=f("ui-round-btn-vue"),r=f("ui-box-vue"),n=f("form-auth-vue");return b(),Q(R,null,[o(r,{columns:e.boxCols,header:"Authentication Radius",buttonShow:e.buttonShow,buttonLabel:"Add",buttonClick:e.handleAdd},{default:l(()=>[o(E,{"hide-header":"",rows:e.rows,columns:e.columns,"row-key":"name","rows-per-page-options":[10,50,100,0]},{"body-cell-state":l(s=>[o(v,{props:s,class:"wd-20"},{default:l(()=>[o(T,{name:"mdi-account-eye-outline",size:"1.2rem"})]),_:2},1032,["props"])]),"body-cell-actions":l(s=>[o(v,{props:s},{default:l(()=>[o(h,{dense:"",round:"",color:"primary",size:"md",icon:"add",onClick:t[0]||(t[0]=d=>e.handleAdd())}),o(a,{color:"positive",icon:"mode_edit",tooltip:"Create user",onClick:d=>e.handleEdit(s.row)},null,8,["onClick"]),o(a,{color:"negative",icon:"delete",tooltip:"Delete user",onClick:d=>e.handleDelete(s.row)},null,8,["onClick"])]),_:2},1032,["props"])]),_:1},8,["rows","columns"])]),_:1},8,["columns","buttonShow","buttonClick"]),e.visible?(b(),w(n,{key:0,ref:"form",onSubmit:e.handleSubmit,onCancel:e.handleCancel},null,8,["onSubmit","onCancel"])):B("",!0)],64)}const _e=C(j,[["render",J],["__scopeId","data-v-d5d3074f"]]);export{_e as default};
