(window.webpackJsonp=window.webpackJsonp||[]).push([[52],{192:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return o})),r.d(t,"metadata",(function(){return i})),r.d(t,"rightToc",(function(){return l})),r.d(t,"default",(function(){return u}));var n=r(1),a=r(9),c=(r(0),r(196)),o={title:"Image [object]"},i={id:"reference/opspec/op-directory/op/call/container/image",title:"Image [object]",description:"An object which defines the image of a container call.",source:"@site/docs/reference/opspec/op-directory/op/call/container/image.md",permalink:"/docs/reference/opspec/op-directory/op/call/container/image",editUrl:"https://github.com/opctl/opctl/edit/master/docs/docs/reference/opspec/op-directory/op/call/container/image.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1587672399,sidebar:"docs",previous:{title:"Container Call [object]",permalink:"/docs/reference/opspec/op-directory/op/call/container/index"},next:{title:"Loop Vars [object]",permalink:"/docs/reference/opspec/op-directory/op/call/loop-vars"}},l=[{value:"Properties",id:"properties",children:[{value:"ref",id:"ref",children:[]},{value:"Example ref (docker.io/ubuntu:19.10)",id:"example-ref-dockerioubuntu1910",children:[]},{value:"Example ref (variable)",id:"example-ref-variable",children:[]},{value:"pullCreds",id:"pullcreds",children:[]}]}],p={rightToc:l},b="wrapper";function u(e){var t=e.components,r=Object(a.a)(e,["components"]);return Object(c.b)(b,Object(n.a)({},p,r,{components:t,mdxType:"MDXLayout"}),Object(c.b)("p",null,"An object which defines the image of a container call."),Object(c.b)("h2",{id:"properties"},"Properties"),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},"must have",Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"#ref"}),"ref")))),Object(c.b)("li",{parentName:"ul"},"may have",Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"#pullcreds"}),"pullCreds"))))),Object(c.b)("h3",{id:"ref"},"ref"),Object(c.b)("p",null,"A string referencing a local or remote image."),Object(c.b)("p",null,"Must be one of:"),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},"a ",Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/reference/opspec/op-directory/op/variable-reference"}),"variable-reference [string]")," evaluating to a ",Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"https://github.com/opencontainers/image-spec/blob/v1.0.1/image-layout.md"}),"v1.0.1 OCI (Open Container Initiative) ",Object(c.b)("inlineCode",{parentName:"a"},"image-layout")),"."),Object(c.b)("li",{parentName:"ul"},"a ",Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/reference/opspec/types/string#initialization"}),"string initializer")," evaluating to a docker image name i.e. ",Object(c.b)("inlineCode",{parentName:"li"},"[host][repository]image[tag]")," where by default host is ",Object(c.b)("inlineCode",{parentName:"li"},"docker.io")," and tag is ",Object(c.b)("inlineCode",{parentName:"li"},"latest"))),Object(c.b)("h3",{id:"example-ref-dockerioubuntu1910"},"Example ref (",Object(c.b)("a",Object(n.a)({parentName:"h3"},{href:"https://hub.docker.com/_/ubuntu"}),"docker.io/ubuntu:19.10"),")"),Object(c.b)("p",null,Object(c.b)("inlineCode",{parentName:"p"},"ref: 'ubuntu:19.10'")," or ",Object(c.b)("inlineCode",{parentName:"p"},"ref: 'docker.io/ubuntu:19.10'")),Object(c.b)("h3",{id:"example-ref-variable"},"Example ref (variable)"),Object(c.b)("p",null,Object(c.b)("inlineCode",{parentName:"p"},"ref: $(myOCIImageLayoutDir)")),Object(c.b)("h3",{id:"pullcreds"},"pullCreds"),Object(c.b)("p",null,"A ",Object(c.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/reference/opspec/op-directory/op/call/pull-creds"}),"pull-creds [object]")," defining creds used to pull the image from a private source."))}u.isMDXComponent=!0},196:function(e,t,r){"use strict";r.d(t,"a",(function(){return u})),r.d(t,"b",(function(){return m}));var n=r(0),a=r.n(n);function c(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function o(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function i(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?o(Object(r),!0).forEach((function(t){c(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):o(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function l(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},c=Object.keys(e);for(n=0;n<c.length;n++)r=c[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(e);for(n=0;n<c.length;n++)r=c[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var p=a.a.createContext({}),b=function(e){var t=a.a.useContext(p),r=t;return e&&(r="function"==typeof e?e(t):i({},t,{},e)),r},u=function(e){var t=b(e.components);return(a.a.createElement(p.Provider,{value:t},e.children))},d="mdxType",s={inlineCode:"code",wrapper:function(e){var t=e.children;return a.a.createElement(a.a.Fragment,{},t)}},f=Object(n.forwardRef)((function(e,t){var r=e.components,n=e.mdxType,c=e.originalType,o=e.parentName,p=l(e,["components","mdxType","originalType","parentName"]),u=b(r),d=n,f=u["".concat(o,".").concat(d)]||u[d]||s[d]||c;return r?a.a.createElement(f,i({ref:t},p,{components:r})):a.a.createElement(f,i({ref:t},p))}));function m(e,t){var r=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var c=r.length,o=new Array(c);o[0]=f;var i={};for(var l in t)hasOwnProperty.call(t,l)&&(i[l]=t[l]);i.originalType=e,i[d]="string"==typeof e?e:n,o[1]=i;for(var p=2;p<c;p++)o[p]=r[p];return a.a.createElement.apply(null,o)}return a.a.createElement.apply(null,r)}f.displayName="MDXCreateElement"}}]);