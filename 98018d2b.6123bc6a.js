(window.webpackJsonp=window.webpackJsonp||[]).push([[32],{172:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return c})),r.d(t,"metadata",(function(){return o})),r.d(t,"rightToc",(function(){return l})),r.d(t,"default",(function(){return b}));var n=r(1),i=r(9),a=(r(0),r(197)),c={title:"File"},o={id:"reference/opspec/types/file",title:"File",description:"File typed values are a filesystem file entry.",source:"@site/docs/reference/opspec/types/file.md",permalink:"/docs/reference/opspec/types/file",editUrl:"https://github.com/opctl/opctl/edit/master/website/docs/reference/opspec/types/file.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1587672399,sidebar:"docs",previous:{title:"Dir",permalink:"/docs/reference/opspec/types/dir"},next:{title:"Number",permalink:"/docs/reference/opspec/types/number"}},l=[{value:"Initialization",id:"initialization",children:[]},{value:"Coercion",id:"coercion",children:[]}],p={rightToc:l},s="wrapper";function b(e){var t=e.components,r=Object(i.a)(e,["components"]);return Object(a.b)(s,Object(n.a)({},p,r,{components:t,mdxType:"MDXLayout"}),Object(a.b)("p",null,"File typed values are a filesystem file entry."),Object(a.b)("p",null,"Files..."),Object(a.b)("ul",null,Object(a.b)("li",{parentName:"ul"},"are mutable, i.e. making changes to a file results in the file being changed everywhere it's referenced"),Object(a.b)("li",{parentName:"ul"},"can be passed in/out of ops via ",Object(a.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/reference/opspec/op-directory/op/parameter/file"}),"file parameters")),Object(a.b)("li",{parentName:"ul"},"can be initialized via ",Object(a.b)("a",Object(n.a)({parentName:"li"},{href:"#initialization"}),"file initialization")),Object(a.b)("li",{parentName:"ul"},"are coerced according to ",Object(a.b)("a",Object(n.a)({parentName:"li"},{href:"#coercion"}),"file coercion"))),Object(a.b)("h3",{id:"initialization"},"Initialization"),Object(a.b)("p",null,"File typed values can be constructed from a literal string or templated string; see ",Object(a.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/reference/opspec/types/string#initialization"}),"string initialization"),"."),Object(a.b)("h3",{id:"coercion"},"Coercion"),Object(a.b)("p",null,"File typed values are coercible to:"),Object(a.b)("ul",null,Object(a.b)("li",{parentName:"ul"},Object(a.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/reference/opspec/types/boolean"}),"boolean")," (files which are empty, all ",Object(a.b)("inlineCode",{parentName:"li"},'"0"'),", or (case insensitive) ",Object(a.b)("inlineCode",{parentName:"li"},'"f"')," or ",Object(a.b)("inlineCode",{parentName:"li"},'"false"')," coerce to false; all else coerce to true)"),Object(a.b)("li",{parentName:"ul"},Object(a.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/reference/opspec/types/array"}),"array")," (if value of file is an array in JSON format)"),Object(a.b)("li",{parentName:"ul"},Object(a.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/reference/opspec/types/number"}),"number")," (if value of file is numeric)"),Object(a.b)("li",{parentName:"ul"},Object(a.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/reference/opspec/types/object"}),"object")," (if value of file is an object in JSON format)"),Object(a.b)("li",{parentName:"ul"},Object(a.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/reference/opspec/types/string"}),"string"))))}b.isMDXComponent=!0},197:function(e,t,r){"use strict";r.d(t,"a",(function(){return b})),r.d(t,"b",(function(){return d}));var n=r(0),i=r.n(n);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function c(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?c(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):c(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function l(e,t){if(null==e)return{};var r,n,i=function(e,t){if(null==e)return{};var r,n,i={},a=Object.keys(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||(i[r]=e[r]);return i}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(i[r]=e[r])}return i}var p=i.a.createContext({}),s=function(e){var t=i.a.useContext(p),r=t;return e&&(r="function"==typeof e?e(t):o({},t,{},e)),r},b=function(e){var t=s(e.components);return(i.a.createElement(p.Provider,{value:t},e.children))},f="mdxType",u={inlineCode:"code",wrapper:function(e){var t=e.children;return i.a.createElement(i.a.Fragment,{},t)}},m=Object(n.forwardRef)((function(e,t){var r=e.components,n=e.mdxType,a=e.originalType,c=e.parentName,p=l(e,["components","mdxType","originalType","parentName"]),b=s(r),f=n,m=b["".concat(c,".").concat(f)]||b[f]||u[f]||a;return r?i.a.createElement(m,o({ref:t},p,{components:r})):i.a.createElement(m,o({ref:t},p))}));function d(e,t){var r=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var a=r.length,c=new Array(a);c[0]=m;var o={};for(var l in t)hasOwnProperty.call(t,l)&&(o[l]=t[l]);o.originalType=e,o[f]="string"==typeof e?e:n,c[1]=o;for(var p=2;p<a;p++)c[p]=r[p];return i.a.createElement.apply(null,c)}return i.a.createElement.apply(null,r)}m.displayName="MDXCreateElement"}}]);