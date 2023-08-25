import{a as u,c as l,b as d,N as g}from"./NetworkError-b28af479.js";import{b as f,a as m}from"./index-dc2fb6bd.js";const c=u.create({baseURL:"/api/v1",withCredentials:!1,headers:{Accept:"application/json","Content-Type":"application/json"},validateStatus(t){return![408,413,429].includes(t)&&t<500}});let n=!1;const p=6e4,s=f(),a=m();c.interceptors.request.use(t=>{const e=a.user;return a.loggedIn&&(t.headers.Authorization=`Bearer ${e.token}`,t.method==="get"&&s.start()),t});c.interceptors.response.use(t=>{if(s.stop(),t.status>=400){const e=t.data.message??t.data??t.statusText;return a.loggedIn&&a.logout(),s.fault(e),l({statusCode:t.status,message:e},t)}else return n||(n=!0,setTimeout(async()=>{a.refresh(),n=!1},p)),d(t.data.data,t)},async t=>{var i;s.stop();const{response:e}=t,o=t.config,r=`[${o.method} ${o.url}] error: ${(i=e.data)==null?void 0:i.message}`||e.statusText;return console.log(r),s.fault(r),!t.response||t.code==="ECONNABORTED"?Promise.reject(new g(t.request)):Promise.reject(r)});export{c as j};
