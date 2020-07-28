(window.webpackJsonp=window.webpackJsonp||[]).push([[9],{149:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return i})),r.d(t,"metadata",(function(){return c})),r.d(t,"rightToc",(function(){return p})),r.d(t,"default",(function(){return d}));var n=r(1),a=r(9),o=(r(0),r(197)),i={title:"Dir Parameter [object]"},c={id:"reference/opspec/op-directory/op/parameter/dir",title:"Dir Parameter [object]",description:"An object defining a parameter which accepts a [dir typed value](../../../types/dir.md).",source:"@site/docs/reference/opspec/op-directory/op/parameter/dir.md",permalink:"/docs/reference/opspec/op-directory/op/parameter/dir",editUrl:"https://github.com/opctl/opctl/edit/master/website/docs/reference/opspec/op-directory/op/parameter/dir.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1588264700,sidebar:"docs",previous:{title:"Boolean Parameter [object]",permalink:"/docs/reference/opspec/op-directory/op/parameter/boolean"},next:{title:"File Parameter [object]",permalink:"/docs/reference/opspec/op-directory/op/parameter/file"}},p=[{value:"Properties:",id:"properties",children:[{value:"default",id:"default",children:[]},{value:"description",id:"description",children:[]},{value:"isSecret",id:"issecret",children:[]}]}],l={rightToc:p},u="wrapper";function d(e){var t=e.components,r=Object(a.a)(e,["components"]);return Object(o.b)(u,Object(n.a)({},l,r,{components:t,mdxType:"MDXLayout"}),Object(o.b)("p",null,"An object defining a parameter which accepts a ",Object(o.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/reference/opspec/types/dir"}),"dir typed value"),"."),Object(o.b)("h2",{id:"properties"},"Properties:"),Object(o.b)("ul",null,Object(o.b)("li",{parentName:"ul"},"must have:",Object(o.b)("ul",{parentName:"li"},Object(o.b)("li",{parentName:"ul"},Object(o.b)("a",Object(n.a)({parentName:"li"},{href:"#description"}),"description")))),Object(o.b)("li",{parentName:"ul"},"may have:",Object(o.b)("ul",{parentName:"li"},Object(o.b)("li",{parentName:"ul"},Object(o.b)("a",Object(n.a)({parentName:"li"},{href:"#default"}),"default")),Object(o.b)("li",{parentName:"ul"},Object(o.b)("a",Object(n.a)({parentName:"li"},{href:"#issecret"}),"isSecret"))))),Object(o.b)("h3",{id:"default"},"default"),Object(o.b)("p",null,"A relative or absolute path to use as the default value of the parameter when no argument is provided."),Object(o.b)("p",null,"If the value is..."),Object(o.b)("ul",null,Object(o.b)("li",{parentName:"ul"},"an absolute path, the value is interpreted from the root of the op."),Object(o.b)("li",{parentName:"ul"},"a relative path, the value is interpreted from the current working directory at the time the op is called.",Object(o.b)("blockquote",{parentName:"li"},Object(o.b)("p",{parentName:"blockquote"},"relative path defaults are ignored when an op is called from an op as there is no current working directory.")))),Object(o.b)("h3",{id:"description"},"description"),Object(o.b)("p",null,"A ",Object(o.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/reference/opspec/op-directory/op/markdown"}),"markdown [string]")," defining a human friendly description of the parameter."),Object(o.b)("h3",{id:"issecret"},"isSecret"),Object(o.b)("p",null,"A boolean indicating if the value of the parameter is secret. This will cause it to be hidden in UI's for example. "))}d.isMDXComponent=!0},197:function(e,t,r){"use strict";r.d(t,"a",(function(){return d})),r.d(t,"b",(function(){return m}));var n=r(0),a=r.n(n);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function c(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function p(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},o=Object.keys(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var l=a.a.createContext({}),u=function(e){var t=a.a.useContext(l),r=t;return e&&(r="function"==typeof e?e(t):c({},t,{},e)),r},d=function(e){var t=u(e.components);return(a.a.createElement(l.Provider,{value:t},e.children))},b="mdxType",s={inlineCode:"code",wrapper:function(e){var t=e.children;return a.a.createElement(a.a.Fragment,{},t)}},f=Object(n.forwardRef)((function(e,t){var r=e.components,n=e.mdxType,o=e.originalType,i=e.parentName,l=p(e,["components","mdxType","originalType","parentName"]),d=u(r),b=n,f=d["".concat(i,".").concat(b)]||d[b]||s[b]||o;return r?a.a.createElement(f,c({ref:t},l,{components:r})):a.a.createElement(f,c({ref:t},l))}));function m(e,t){var r=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var o=r.length,i=new Array(o);i[0]=f;var c={};for(var p in t)hasOwnProperty.call(t,p)&&(c[p]=t[p]);c.originalType=e,c[b]="string"==typeof e?e:n,i[1]=c;for(var l=2;l<o;l++)i[l]=r[l];return a.a.createElement.apply(null,i)}return a.a.createElement.apply(null,r)}f.displayName="MDXCreateElement"}}]);