"use strict";(globalThis["webpackChunkquasar"]=globalThis["webpackChunkquasar"]||[]).push([[191],{7191:(e,t,s)=>{s.r(t),s.d(t,{default:()=>C});var o=s(9835),a=s(6970);const l={class:"q-pa-md"},n={class:"row justify-center items-start crisper"},r=(0,o.Uk)("SSH hosts management"),c={class:"text-subtitle1 text-bold host-name"},i={class:"text-subtitle2"},d={class:"text-subtitle2"},m={class:"text-subtitle2"};function u(e,t,s,u,p,w){const b=(0,o.up)("q-item-label"),v=(0,o.up)("q-item"),_=(0,o.up)("q-icon"),h=(0,o.up)("q-td"),g=(0,o.up)("q-btn"),f=(0,o.up)("q-table");return(0,o.wg)(),(0,o.iD)("div",l,[(0,o._)("div",n,[(0,o._)("div",{class:(0,a.C_)(["square rounded-borders",u.cols])},[(0,o.Wm)(v,{class:"bot-line"},{default:(0,o.w5)((()=>[(0,o.Wm)(b,{class:"text-bold text-h6"},{default:(0,o.w5)((()=>[r])),_:1})])),_:1}),(0,o.Wm)(f,{"hide-header":"",rows:u.rows,columns:u.columns,"row-key":"name","rows-per-page-options":[10,50,100,0]},{"body-cell-state":(0,o.w5)((e=>[(0,o.Wm)(h,{props:e,class:"wd-20"},{default:(0,o.w5)((()=>[(0,o.Wm)(_,{name:u.activeIcon(e.row),size:"1.2rem",color:u.activeColor(e.row)},null,8,["name","color"])])),_:2},1032,["props"])])),"body-cell-icon":(0,o.w5)((e=>[(0,o.Wm)(h,{props:e,class:"wd-80"},{default:(0,o.w5)((()=>[(0,o.Wm)(_,{name:"computer",size:"md"})])),_:2},1032,["props"])])),"body-cell-key":(0,o.w5)((e=>[(0,o.Wm)(h,{props:e,class:"wd-100"},{default:(0,o.w5)((()=>[(0,o._)("div",c,(0,a.zw)(e.row.name),1),(0,o._)("div",i,(0,a.zw)(e.row.finger),1),(0,o._)("div",d,(0,a.zw)(e.row.createdat),1),(0,o._)("div",m,(0,a.zw)(e.row.usedat),1)])),_:2},1032,["props"])])),"body-cell-actions":(0,o.w5)((e=>[(0,o.Wm)(h,{props:e},{default:(0,o.w5)((()=>[(0,o.Wm)(g,{class:"q-ml-xs",dense:"",round:"",color:"negative",size:"md",icon:"delete",onClick:t=>u.onDelete(e.row)},null,8,["onClick"])])),_:2},1032,["props"])])),_:1},8,["rows","columns"])],2)])])}var p=s(9302),w=s(499);const b=[{name:"state"},{name:"icon"},{name:"key"},{name:"name"},{name:"finger"},{name:"createdat"},{name:"usedat"},{name:"actions"}],v=[{active:!0,name:"192.168.1.6",finger:"b3BlbnNzaC1rZXkt",createdat:"Добавлено Aug 12, 2022 10:51:36",used:" Последний раз использовался Sep 20, 2022 02:00:05"}],_={setup(){const e=(0,p.Z)(),t=(0,w.iH)(null),s=(0,w.iH)(null),a=(0,w.iH)(null);return{create:(0,w.iH)(!1),columns:b,rows:v,name:t,vpnlist:s,note:a,isShowHeaderButton:(0,o.Fl)((()=>0===v.length)),cols:(0,o.Fl)((()=>"col-"+("sm"==e.screen.name?8:"xs"==e.screen.name?11:4))),onEdit(e){console.log(e)},onDelete(e){console.log(e)},onSubmit(e){this.create=!1},activeIcon(e){return e.active?"task_alt":"highlight_off"},activeColor(e){return e.active?"positive":"grey"}}},methods:{}};var h=s(1639),g=s(490),f=s(3115),k=s(4356),q=s(7220),W=s(2857),x=s(4455),y=s(9984),Z=s.n(y);const z=(0,h.Z)(_,[["render",u],["__scopeId","data-v-20295d11"]]),C=z;Z()(_,"components",{QItem:g.Z,QItemLabel:f.Z,QTable:k.Z,QTd:q.Z,QIcon:W.Z,QBtn:x.Z})}}]);