"use strict";(globalThis["webpackChunkquasar"]=globalThis["webpackChunkquasar"]||[]).push([[977],{4977:(e,l,o)=>{o.r(l),o.d(l,{default:()=>x});var n=o(9835),t=o(6970);const c={class:"q-pa-md"},s={class:"row justify-center items-start crisper"},a=(0,n.Uk)("Table Synchronization"),d={class:"q-mt-md"};function r(e,l,o,r,i,m){const u=(0,n.up)("q-item-label"),p=(0,n.up)("q-item"),b=(0,n.up)("q-checkbox"),w=(0,n.up)("q-space"),g=(0,n.up)("q-btn"),S=(0,n.up)("q-td"),h=(0,n.up)("q-table");return(0,n.wg)(),(0,n.iD)("div",c,[(0,n._)("div",s,[(0,n._)("div",{class:(0,t.C_)(["square rounded-borders",r.cols])},[(0,n.Wm)(p,{class:"bot-line"},{default:(0,n.w5)((()=>[(0,n.Wm)(u,{class:"text-bold text-h6"},{default:(0,n.w5)((()=>[a])),_:1})])),_:1}),(0,n.Wm)(h,{"hide-header":"",rows:r.rows,columns:r.columns,"row-key":"table","rows-per-page-options":[10,50,100,0]},{top:(0,n.w5)((()=>[(0,n.Wm)(b,{modelValue:r.isSelected,"onUpdate:modelValue":[l[0]||(l[0]=e=>r.isSelected=e),r.onSelectedAll],color:"grey"},null,8,["modelValue","onUpdate:modelValue"]),(0,n.Wm)(w),(0,n.Wm)(g,{class:"q-ml-xs",dense:"",color:"warning",size:"md",icon:"sync",disabled:r.isDisabled,onClick:l[1]||(l[1]=e=>r.onAllSync())},null,8,["disabled"])])),"body-cell-selected":(0,n.w5)((e=>[(0,n.Wm)(S,{props:e,class:"wd-40"},{default:(0,n.w5)((()=>[(0,n.Wm)(b,{modelValue:e.row.model,"onUpdate:modelValue":[l=>e.row.model=l,r.onSelected],color:"grey",val:e.row.selected},null,8,["modelValue","onUpdate:modelValue","val"])])),_:2},1032,["props"])])),"body-cell-direction":(0,n.w5)((e=>[(0,n.Wm)(S,{props:e},{default:(0,n.w5)((()=>[(0,n.Wm)(g,{class:"q-ml-xs",dense:"",color:"warning",size:"md",val:e.rowIndex,label:e.row.direction,onClick:l=>r.onDirection(e.rowIndex)},null,8,["val","label","onClick"])])),_:2},1032,["props"])])),"body-cell-action":(0,n.w5)((e=>[(0,n.Wm)(S,{props:e},{default:(0,n.w5)((()=>[(0,n.Wm)(g,{class:"q-ml-xs",dense:"",color:"warning",size:"md",icon:"sync",onClick:l=>r.onSync(e.row)},null,8,["onClick"])])),_:2},1032,["props"])])),_:1},8,["rows","columns"]),(0,n._)("div",d," Selected: "+(0,t.zw)(JSON.stringify(r.selected)),1)],2)])])}var i=o(9302),m=o(499);const u=[{name:"selected"},{name:"model"},{name:"table",align:"left",field:"table",sortable:!0},{name:"method"},{name:"direction"},{name:"action"}],p=[{selected:!1,model:(0,m.iH)([]),table:"bmx280",method:"replace",direction:"M=>S"},{selected:!1,model:(0,m.iH)([]),table:"ds18b20",method:"replace",direction:"S=>M"}],b={setup(){const e=(0,i.Z)(),l=(0,m.iH)([]),o=(0,m.iH)(!1);return{columns:u,rows:p,selected:l,cols:(0,n.Fl)((()=>"col-"+("sm"==e.screen.name?8:"xs"==e.screen.name?11:5))),isSelected:o,isDisabled:(0,n.Fl)((()=>0===l.value.length)),getSelectedString(){return 0===l.value.length?"":`${l.value.length} record${l.value.length>1?"s":""} selected of ${p.length}`},onSync(e){console.log("onSync not implemented")},onDirection(e){p[e].direction="M=>S"===p[e].direction?"S=>M":"M=>S"},onAllSync(){console.log("onAllSync not implemented")},onSelected(){o.value=l.value.length===p.length||0!==l.value.length&&null},onSelectedAll(){console.log(o.value),!0===o.value?console.log("unselect all"):console.log("select all")}}},computed:{directionLabel:function(e){return"M=>S"}}};var w=o(1639),g=o(490),S=o(3115),h=o(4356),v=o(6937),y=o(136),f=o(4455),q=o(7220),k=o(9984),W=o.n(k);const _=(0,w.Z)(b,[["render",r],["__scopeId","data-v-f01bf1ce"]]),x=_;W()(b,"components",{QItem:g.Z,QItemLabel:S.Z,QTable:h.Z,QCheckbox:v.Z,QSpace:y.Z,QBtn:f.Z,QTd:q.Z})}}]);