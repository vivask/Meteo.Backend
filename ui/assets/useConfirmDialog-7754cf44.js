import{D as o}from"./vendor-7e3b54f2.js";function a(){return{show:n=>new Promise((e,s)=>{o.create({title:"Confirm",message:n,cancel:!0}).onOk(()=>e(!0)).onCancel(()=>e(!1)).onDismiss(()=>e(!1))})}}export{a as u};
