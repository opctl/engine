(window.webpackJsonp=window.webpackJsonp||[]).push([[45],{185:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return o})),r.d(t,"metadata",(function(){return i})),r.d(t,"rightToc",(function(){return p})),r.d(t,"default",(function(){return b}));var n=r(1),a=r(9),c=(r(0),r(197)),o={title:"Array Parameter [object]"},i={id:"reference/opspec/op-directory/op/parameter/array",title:"Array Parameter [object]",description:"An object defining a parameter which accepts an [array typed value](../../../types/array.md).",source:"@site/docs/reference/opspec/op-directory/op/parameter/array.md",permalink:"/docs/reference/opspec/op-directory/op/parameter/array",editUrl:"https://github.com/opctl/opctl/edit/master/website/docs/reference/opspec/op-directory/op/parameter/array.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1588264700,sidebar:"docs",previous:{title:"Parameter [object]",permalink:"/docs/reference/opspec/op-directory/op/parameter/index"},next:{title:"Boolean Parameter [object]",permalink:"/docs/reference/opspec/op-directory/op/parameter/boolean"}},p=[{value:"Properties:",id:"properties",children:[{value:"constraints",id:"constraints",children:[]},{value:"default",id:"default",children:[]},{value:"description",id:"description",children:[]},{value:"isSecret",id:"issecret",children:[]}]}],l={rightToc:p},s="wrapper";function b(e){var t=e.components,r=Object(a.a)(e,["components"]);return Object(c.b)(s,Object(n.a)({},l,r,{components:t,mdxType:"MDXLayout"}),Object(c.b)("p",null,"An object defining a parameter which accepts an ",Object(c.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/reference/opspec/types/array"}),"array typed value"),"."),Object(c.b)("h2",{id:"properties"},"Properties:"),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},"must have:",Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"#description"}),"description")))),Object(c.b)("li",{parentName:"ul"},"may have:",Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"#constraints"}),"constraints")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"#default"}),"default")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"#issecret"}),"isSecret"))))),Object(c.b)("h3",{id:"constraints"},"constraints"),Object(c.b)("p",null,"A ",Object(c.b)("a",Object(n.a)({parentName:"p"},{href:"https://tools.ietf.org/html/draft-wright-json-schema-00"}),"JSON Schema v4 [object]")," defining constraints on the parameters value."),Object(c.b)("h3",{id:"default"},"default"),Object(c.b)("p",null,"An array to use as the value of the parameter when no argument is provided."),Object(c.b)("h3",{id:"description"},"description"),Object(c.b)("p",null,"A ",Object(c.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/reference/opspec/op-directory/op/markdown"}),"markdown [string]")," defining a human friendly description of the parameter."),Object(c.b)("h3",{id:"issecret"},"isSecret"),Object(c.b)("p",null,"An boolean indicating if the value of the parameter is secret. This will cause it to be hidden in UI's for example. "))}b.isMDXComponent=!0},197:function(e,t,r){"use strict";r.d(t,"a",(function(){return b})),r.d(t,"b",(function(){return m}));var n=r(0),a=r.n(n);function c(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function o(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function i(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?o(Object(r),!0).forEach((function(t){c(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):o(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function p(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},c=Object.keys(e);for(n=0;n<c.length;n++)r=c[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(e);for(n=0;n<c.length;n++)r=c[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var l=a.a.createContext({}),s=function(e){var t=a.a.useContext(l),r=t;return e&&(r="function"==typeof e?e(t):i({},t,{},e)),r},b=function(e){var t=s(e.components);return(a.a.createElement(l.Provider,{value:t},e.children))},u="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return a.a.createElement(a.a.Fragment,{},t)}},f=Object(n.forwardRef)((function(e,t){var r=e.components,n=e.mdxType,c=e.originalType,o=e.parentName,l=p(e,["components","mdxType","originalType","parentName"]),b=s(r),u=n,f=b["".concat(o,".").concat(u)]||b[u]||d[u]||c;return r?a.a.createElement(f,i({ref:t},l,{components:r})):a.a.createElement(f,i({ref:t},l))}));function m(e,t){var r=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var c=r.length,o=new Array(c);o[0]=f;var i={};for(var p in t)hasOwnProperty.call(t,p)&&(i[p]=t[p]);i.originalType=e,i[u]="string"==typeof e?e:n,o[1]=i;for(var l=2;l<c;l++)o[l]=r[l];return a.a.createElement.apply(null,o)}return a.a.createElement.apply(null,r)}f.displayName="MDXCreateElement"}}]);