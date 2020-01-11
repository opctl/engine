(window.webpackJsonp=window.webpackJsonp||[]).push([[20],{144:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return i})),r.d(t,"metadata",(function(){return o})),r.d(t,"rightToc",(function(){return l})),r.d(t,"default",(function(){return u}));var n=r(1),a=r(9),c=(r(0),r(177)),i={title:"Number"},o={id:"opspec/reference/types/number",title:"Number",description:"Number typed values are a number.",source:"@site/docs/opspec/reference/types/number.md",permalink:"/docs/opspec/reference/types/number",editUrl:"https://github.com/opctl/opctl/edit/master/docs/docs/opspec/reference/types/number.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1578700982,sidebar:"docs",previous:{title:"File",permalink:"/docs/opspec/reference/types/file"},next:{title:"Object",permalink:"/docs/opspec/reference/types/object"}},l=[{value:"Initialization",id:"initialization",children:[]},{value:"Coercion",id:"coercion",children:[]}],b={rightToc:l},p="wrapper";function u(e){var t=e.components,r=Object(a.a)(e,["components"]);return Object(c.b)(p,Object(n.a)({},b,r,{components:t,mdxType:"MDXLayout"}),Object(c.b)("p",null,"Number typed values are a number."),Object(c.b)("p",null,"Numbers..."),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},"are immutable, i.e. assigning to an number results in a copy of the original number"),Object(c.b)("li",{parentName:"ul"},"can be passed in/out of ops via ",Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"../structure/op-directory/op/parameter/number"}),"number parameters")),Object(c.b)("li",{parentName:"ul"},"can be initialized via ",Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"#initialization"}),"number initialization")),Object(c.b)("li",{parentName:"ul"},"are coerced according to ",Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"#coercion"}),"number coercion"))),Object(c.b)("h3",{id:"initialization"},"Initialization"),Object(c.b)("p",null,"Number typed values can be constructed from a literal or templated number."),Object(c.b)("p",null,"A templated number is a number which includes one or more value reference.\nAt runtime, each reference gets evaluated and replaced with it's corresponding value."),Object(c.b)("h4",{id:"initialization-example-literal"},"Initialization Example (literal)"),Object(c.b)("pre",null,Object(c.b)("code",Object(n.a)({parentName:"pre"},{className:"language-yaml"}),"2\n")),Object(c.b)("h4",{id:"initialization-example-templated"},"Initialization Example (templated)"),Object(c.b)("p",null,"given:"),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},Object(c.b)("inlineCode",{parentName:"li"},"someNumber"),Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},"is in scope"),Object(c.b)("li",{parentName:"ul"},"is type coercible to number")))),Object(c.b)("pre",null,Object(c.b)("code",Object(n.a)({parentName:"pre"},{className:"language-yaml"}),"# $(someNumber) replaced w/ someNumber\n222$(someNumber)3e10\n")),Object(c.b)("h3",{id:"coercion"},"Coercion"),Object(c.b)("p",null,"Number typed values are coercible to:"),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/opspec/reference/types/boolean"}),"boolean")," (numbers which are 0 coerce to false; all else coerce to true)"),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/opspec/reference/types/file"}),"file")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/opspec/reference/types/string"}),"string"))),Object(c.b)("h4",{id:"coercion-example-number-to-file"},"Coercion Example (number to file)"),Object(c.b)("pre",null,Object(c.b)("code",Object(n.a)({parentName:"pre"},{className:"language-yaml"}),"name: numAsFile\nrun:\n  container:\n    image: { ref: alpine }\n    cmd:\n    - sh\n    - -ce\n    - cat /numCoercedToFile\n    files:\n      /numCoercedToFile: 2.2\n")))}u.isMDXComponent=!0},177:function(e,t,r){"use strict";r.d(t,"a",(function(){return u})),r.d(t,"b",(function(){return f}));var n=r(0),a=r.n(n);function c(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){c(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function l(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},c=Object.keys(e);for(n=0;n<c.length;n++)r=c[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(e);for(n=0;n<c.length;n++)r=c[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var b=a.a.createContext({}),p=function(e){var t=a.a.useContext(b),r=t;return e&&(r="function"==typeof e?e(t):o({},t,{},e)),r},u=function(e){var t=p(e.components);return a.a.createElement(b.Provider,{value:t},e.children)},m="mdxType",s={inlineCode:"code",wrapper:function(e){var t=e.children;return a.a.createElement(a.a.Fragment,{},t)}},d=Object(n.forwardRef)((function(e,t){var r=e.components,n=e.mdxType,c=e.originalType,i=e.parentName,b=l(e,["components","mdxType","originalType","parentName"]),u=p(r),m=n,d=u["".concat(i,".").concat(m)]||u[m]||s[m]||c;return r?a.a.createElement(d,o({ref:t},b,{components:r})):a.a.createElement(d,o({ref:t},b))}));function f(e,t){var r=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var c=r.length,i=new Array(c);i[0]=d;var o={};for(var l in t)hasOwnProperty.call(t,l)&&(o[l]=t[l]);o.originalType=e,o[m]="string"==typeof e?e:n,i[1]=o;for(var b=2;b<c;b++)i[b]=r[b];return a.a.createElement.apply(null,i)}return a.a.createElement.apply(null,r)}d.displayName="MDXCreateElement"}}]);