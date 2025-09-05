var e=class extends HTMLElement{static get observedAttributes(){return["name"]}attributeChangedCallback(o,l,n){let r=n.length,t=Array(r),s=r-1;for(let i of n)t[s--]=i.codePointAt(0);t=String.fromCodePoint(...t),this.dispatchEvent(new CustomEvent("reverse",{detail:{value:t}}))}};customElements.define("reverse-component",e);
//# sourceMappingURL=reverse-component.js.map
