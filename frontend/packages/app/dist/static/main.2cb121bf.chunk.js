(window.webpackJsonp=window.webpackJsonp||[]).push([[5],{465:function(e,t,a){a(466),e.exports=a(917)},917:function(e,t,a){"use strict";a.r(t);var n={};a.r(n),a.d(n,"WelcomePlugin",(function(){return O}));a(469);var r=a(0),l=a.n(r),o=a(17),i=a.n(o),c=a(38),m=a(91),u=a(951),s=a(58),g=a(925),h=a(919),E=a(399),p=a(960);const d={hour:"2-digit",minute:"2-digit"},b={timeZone:"UTC",...d},f={timeZone:"America/New_York",...d},v={timeZone:"Asia/Tokyo",...d},T={timeZone:"Europe/Stockholm",...d},k={timeNY:"",timeUTC:"",timeTYO:"",timeSTO:""};function y(){const e=new Date,t=window.navigator.language;return{timeNY:e.toLocaleTimeString(t,f),timeUTC:e.toLocaleTimeString(t,b),timeTYO:e.toLocaleTimeString(t,v),timeSTO:e.toLocaleTimeString(t,T)}}var w=()=>{const[{timeNY:e,timeUTC:t,timeTYO:a,timeSTO:n},r]=l.a.useState(k);return l.a.useEffect(()=>{r(y());const e=setInterval(()=>{r(y())},1e3);return()=>{clearInterval(e)}},[]),l.a.createElement(l.a.Fragment,null,l.a.createElement(c.d,{label:"NYC",value:e}),l.a.createElement(c.d,{label:"UTC",value:t}),l.a.createElement(c.d,{label:"STO",value:n}),l.a.createElement(c.d,{label:"TYO",value:a}))};var S=()=>l.a.createElement(c.f,{theme:c.j.home},l.a.createElement(c.c,{title:"Welcome to Backstage",subtitle:"Some quick intro and links."},l.a.createElement(w,null)),l.a.createElement(c.a,null,l.a.createElement(c.b,{title:"Getting Started"},l.a.createElement(c.g,null)),l.a.createElement(u.a,{container:!0},l.a.createElement(u.a,{item:!0,xs:12,md:6},l.a.createElement(c.e,null,l.a.createElement(s.a,{variant:"body1",gutterBottom:!0},"You now have a running instance of Backstage!",l.a.createElement("span",{role:"img","aria-label":"confetti"},"🎉"),"Let's make sure you get the most out of this platform by walking you through the basics."),l.a.createElement(s.a,{variant:"h6",gutterBottom:!0},"The Setup"),l.a.createElement(s.a,{variant:"body1",paragraph:!0},"Backstage is put together from three base concepts: the core, the app and the plugins."),l.a.createElement(g.a,null,l.a.createElement(h.a,null,l.a.createElement(E.a,{primary:"The core is responsible for base functionality."})),l.a.createElement(h.a,null,l.a.createElement(E.a,{primary:"The app provides the base UI and connects the plugins."})),l.a.createElement(h.a,null,l.a.createElement(E.a,{primary:"The plugins make Backstage useful for the end users with specific views and functionality."}))),l.a.createElement(s.a,{variant:"h6",gutterBottom:!0},"Try It Out"),l.a.createElement(s.a,{variant:"body1",paragraph:!0},"We suggest you either check out the documentation for"," ",l.a.createElement(p.a,{href:"https://github.com/spotify/backstage/blob/master/docs/getting-started/create-a-plugin.md"},"creating a plugin")," ","or have a look in the code for the"," ",l.a.createElement(p.a,{component:m.b,to:"/home"},"Home Page")," ",'in the directory "plugins/home-page/src".'))),l.a.createElement(u.a,{item:!0},l.a.createElement(c.e,null,l.a.createElement(s.a,{variant:"h5"},"Quick Links"),l.a.createElement(g.a,null,l.a.createElement(h.a,null,l.a.createElement(p.a,{href:"https://backstage.io"},"backstage.io")),l.a.createElement(h.a,null,l.a.createElement(p.a,{href:"https://github.com/spotify/backstage/blob/master/docs/getting-started/create-a-plugin.md"},"Create a plugin"))))))));const O=Object(c.i)({id:"welcome",register({router:e}){e.registerRoute("/",S)}}),Y=Object(c.h)({plugins:Object.values(n)}),B=Y.getProvider(),C=Y.getRouter(),L=Y.getRoutes();var U=()=>l.a.createElement(B,null,l.a.createElement(C,null,l.a.createElement(L,null)));i.a.render(l.a.createElement(U,null),document.getElementById("root"))}},[[465,10,7,4,3,9,6,0,1,2,8,11]]]);
//# sourceMappingURL=main.2cb121bf.chunk.js.map