import{K as V,r as b,k as T,o as y,f as B,L as n,M as t,aT as I,b0 as U,b2 as N,S as w,aS as F,b3 as E,c as _,g as C,O as g,a_ as D,aQ as W,a7 as H,X as A,R as q,N as P,a6 as K}from"./vendor-7e3b54f2.js";import{b as O,Q as z,a as $}from"./QTable-f6f16d1c.js";import{U as X}from"./UiBox-438aefea.js";import{U as L}from"./UiRoundBtn-7ea8dd22.js";import{u as J}from"./useTableHandlers-649b36f2.js";import{Q as M}from"./QForm-dd93a313.js";import{U as Y}from"./UiInput-ff5536f7.js";import{u as j}from"./useSubmitForm-841c2b76.js";import{j as k}from"./jwtClient-fdc6c2bb.js";import{_ as Q}from"./index-dc2fb6bd.js";import{u as G}from"./useConfirmDialog-7754cf44.js";import{t as Z}from"./tableWrapper-c43f1a46.js";import"./QMenu-7e13e904.js";import"./QTooltip-cd4ef32e.js";import"./QMarkupTable-3806441a.js";import"./QChip-dc1c4db6.js";import"./QItemLabel-0efd9dc5.js";import"./QItemSection-cefe5148.js";import"./format-3895543d.js";import"./NetworkError-b28af479.js";async function x(){const e=[];return k.get("/database/stypes").then(({success:o,result:d})=>o?d:e).catch(()=>e)}const ee=V({name:"FormTableParam",emits:["cancel","submit"],setup(e,{emit:o}){const d=b(null),u=b([]),{localProp:p,show:h,handleSubmit:a,handleCancel:r}=j(d,o);return T(async()=>u.value=await x()),{localProp:p,popup:d,stypes:u,show:h,handleSubmit:a,handleCancel:r}}});function oe(e,o,d,u,p,h){return y(),B(E,{ref:"popup","transition-show":"rotate","transition-hide":"rotate",persistent:""},{default:n(()=>[t(F,{style:{"min-width":"300px"}},{default:n(()=>[t(I,null,{default:n(()=>[t(M,{class:"q-gutter-md",onSubmit:U(e.handleSubmit,["prevent"])},{default:n(()=>[t(O,{modelValue:e.localProp.stype,"onUpdate:modelValue":o[0]||(o[0]=a=>e.localProp.stype=a),outlined:"",dense:"",options:e.stypes,"option-label":"note",hint:"Sync method *","lazy-rules":"",rules:[a=>a||"Please select something"]},null,8,["modelValue","options","rules"]),t(N,{align:"left",class:"text-primary"},{default:n(()=>[t(w,{label:"Submit",type:"submit",color:"primary"}),t(w,{label:"Cancel",color:"primary",flat:"",class:"q-ml-sm",onClick:e.handleCancel},null,8,["onClick"])]),_:1})]),_:1},8,["onSubmit"])]),_:1})]),_:1})]),_:1},512)}const te=Q(ee,[["render",oe]]);const ae=[{name:"value",align:"left",classes:"wd-50"},{name:"actions"}],le=V({name:"FormTable",components:{UiInputVue:Y,FormTableParamVue:te,UiRoundBtnVue:L},emits:["cancel","submit"],setup(e,{emit:o}){const d=b(null),u=b(null),p=b(!1),h=_(()=>p.value?"<<":">>"),{localProp:a,show:r,handleSubmit:l,handleCancel:s,isUpdate:f}=j(d,o),S=G();return{localProp:a,popup:d,form:u,columns:ae,confirm:S,labelBtnParams:h,showParams:p,disable:_(()=>f.value),showBtnAddParam:_(()=>{var m;return p.value?(m=a.value)!=null&&m.params?a.value.params.length===0:!0:!1}),show:r,handleSubmit:l,handleCancel:s,handleAdd(){u.value.show({})},handleEdit(m){u.value.show(m)},async handleDelete(m){await S.show("Are you sure to delete this item?")&&m!==-1&&a.value.params.splice(m,1)},handleParamSubmit(m){var c;m.update||((c=a.value)!=null&&c.params||(a.value.params=[]),a.value.params.push(m.data))}}}}),ne={key:0},se={class:"text-subtitle2 text-left text-primary"},re={key:1};function ie(e,o,d,u,p,h){const a=C("ui-round-btn-vue"),r=C("form-table-param-vue");return y(),g(q,null,[t(E,{ref:"popup","transition-show":"rotate","transition-hide":"rotate",persistent:""},{default:n(()=>[t(F,{style:{"min-width":"350px"}},{default:n(()=>[t(I,null,{default:n(()=>[t(M,{class:"q-gutter-md",onSubmit:U(e.handleSubmit,["prevent"])},{default:n(()=>[t(D,{modelValue:e.localProp.name,"onUpdate:modelValue":o[0]||(o[0]=l=>e.localProp.name=l),dense:"",hint:"Name *",outlined:"","lazy-rules":"",rules:[l=>l&&l.length>0||"Please type something"]},null,8,["modelValue","rules"]),t(D,{modelValue:e.localProp.note,"onUpdate:modelValue":o[1]||(o[1]=l=>e.localProp.note=l),dense:"",outlined:"",hint:"Note"},null,8,["modelValue"]),t(w,{dense:"",class:"wd-320",outline:"",color:"grey",label:e.labelBtnParams,onClick:o[2]||(o[2]=l=>e.showParams=!e.showParams)},null,8,["label"]),e.showParams?(y(),g("div",ne,[t(z,{"hide-header":"","hide-bottom":"",rows:e.localProp.params,columns:e.columns,"row-key":"name"},{"body-cell-value":n(l=>[t($,{props:l},{default:n(()=>[W("div",se,H(l.row.stype.note),1)]),_:2},1032,["props"])]),"body-cell-actions":n(l=>[t($,{props:l},{default:n(()=>[t(a,{color:"primary",icon:"add",onClick:o[3]||(o[3]=s=>e.handleAdd())}),t(a,{color:"positive",icon:"mode_edit",onClick:s=>e.handleEdit(l.row)},null,8,["onClick"]),t(a,{color:"negative",icon:"delete",onClick:s=>e.handleEdit(l.row)},null,8,["onClick"])]),_:2},1032,["props"])]),_:1},8,["rows","columns"])])):A("",!0),e.showBtnAddParam?(y(),g("div",re,[t(w,{class:"wd-320",dense:"",label:"Add",color:"primary",onClick:o[4]||(o[4]=l=>e.handleAdd())})])):A("",!0),t(N,{align:"left",class:"text-primary"},{default:n(()=>[t(w,{label:"Submit",type:"submit",color:"primary"}),t(w,{label:"Cancel",color:"primary",flat:"",class:"q-ml-sm",onClick:e.handleCancel},null,8,["onClick"])]),_:1})]),_:1},8,["onSubmit"])]),_:1})]),_:1})]),_:1},512),t(r,{ref:"form",onSubmit:e.handleParamSubmit},null,8,["onSubmit"])],64)}const de=Q(le,[["render",ie],["__scopeId","data-v-f3a9d0ea"]]),ue="/database/tables";function me(e){return Z(ue,e)}async function ce(e){return k.post("/database/tables/delete",e).then(({success:o})=>o).catch(()=>!1)}async function pe(e){return k.put("/database/import",e).then(({success:o})=>o).catch(()=>!1)}async function be(e){return k.put("/database/save",e).then(({success:o})=>o).catch(()=>!1)}async function fe(e){return k.post("/database/tables/drop",e).then(({success:o})=>o).catch(()=>!1)}const he=[{name:"state",align:"left",classes:"wd-50"},{name:"name",label:"Name",align:"left",field:"name",classes:"wd-100",sortable:!0},{name:"note",label:"Note",align:"left",field:"note",classes:"wd-100",sortable:!0},{name:"actions"}],ve=V({name:"PageTables",components:{UiBoxVue:X,UiRoundBtnVue:L,FormTableVue:de},setup(){const e=b([]),o=me(e),d=b(!0),u=b(null),p={xl:6,lg:6,md:7,sm:11,xs:10},h=_(()=>e.value.length===0),a=G(),r=b([]),l=b(!1),{handleAdd:s,handleEdit:f,handleSubmit:S,handleCancel:m}=J(l,u,e,o,{});return T(async()=>e.value=await o.Get()),{spinner:d,columns:he,rows:e,buttonShow:h,wrapper:o,form:u,boxCols:p,confirm:a,selected:r,visible:l,handleAdd:s,handleEdit:f,handleSubmit:S,handleCancel:m,async handleDelete(c){let i=await a.show("Are you sure to delete this items?");if(i){const v=o.Selected(c,r.value);i=await ce(v),i&&(e.value=await o.Get(),r.value=[])}},async handleImport(c){let i=await a.show("Are you sure to import this table from csv?");if(i){const v=o.Selected(c,r.value);i=await pe(v),i&&(r.value=[],P.create({type:"info",message:"Import completed"}))}},async handleSave(c){let i=await a.show("Are you sure to save this table to csv?");if(i){const v=o.Selected(c,r.value);if(i=await be(v),i){for(let R of v)R.import=!0;r.value=[],P.create({type:"info",message:"Import completed"})}}},async handleDdrop(c){let i=await a.show("Are you sure to drop this tables?");if(i){const v=o.Selected(c,r.value);i=await fe(v),i&&(r.value=[],P.create({type:"info",message:"Drop completed"}))}}}}});function we(e,o,d,u,p,h){const a=C("ui-round-btn-vue"),r=C("ui-box-vue"),l=C("form-table-vue");return y(),g(q,null,[t(r,{columns:e.boxCols,header:"Local hosts",buttonShow:e.buttonShow,buttonLabel:"Add",buttonClick:e.handleAdd},{default:n(()=>[t(z,{selected:e.selected,"onUpdate:selected":o[1]||(o[1]=s=>e.selected=s),"hide-header":"",rows:e.rows,columns:e.columns,"row-key":"name",selection:"multiple","rows-per-page-options":[0,10,50,100]},{"body-cell-state":n(s=>[t($,{props:s},{default:n(()=>[t(K,{name:"mdi-table",size:"2rem"})]),_:2},1032,["props"])]),"body-cell-actions":n(s=>[t($,{props:s},{default:n(()=>[t(a,{color:"primary",icon:"add",tooltip:"Add table name",onClick:o[0]||(o[0]=f=>e.handleAdd())}),t(a,{color:"positive",icon:"mode_edit",tooltip:"Edit table name",onClick:f=>e.handleEdit(s.row)},null,8,["onClick"]),t(a,{color:"negative",icon:"delete",tooltip:"Delete table name",onClick:f=>e.handleDelete(s.row)},null,8,["onClick"]),t(a,{color:"negative",icon:"mdi-table-remove",tooltip:"Drop table",onClick:f=>e.handleDdrop(s.row)},null,8,["onClick"]),t(a,{disable:!s.row.import,color:"warning",icon:"mdi-table-arrow-left",tooltip:"Import table content from csv",onClick:f=>e.handleImport(s.row)},null,8,["disable","onClick"]),t(a,{color:"secondary",icon:"mdi-content-save",tooltip:"Save table content to csv",onClick:f=>e.handleSave(s.row)},null,8,["onClick"])]),_:2},1032,["props"])]),_:1},8,["selected","rows","columns"])]),_:1},8,["columns","buttonShow","buttonClick"]),e.visible?(y(),B(l,{key:0,ref:"form",onSubmit:e.handleSubmit,onCancel:e.handleCancel},null,8,["onSubmit","onCancel"])):A("",!0)],64)}const ze=Q(ve,[["render",we],["__scopeId","data-v-d25d9744"]]);export{ze as default};
