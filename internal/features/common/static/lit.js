/**
 * @license
 * Copyright 2019 Google LLC
 * SPDX-License-Identifier: BSD-3-Clause
 */
const Kt = globalThis, De = Kt.ShadowRoot && (Kt.ShadyCSS === void 0 || Kt.ShadyCSS.nativeShadow) && "adoptedStyleSheets" in Document.prototype && "replace" in CSSStyleSheet.prototype, on = Symbol(), Ue = /* @__PURE__ */ new WeakMap();
let An = class {
  constructor(t, e, n) {
    if (this._$cssResult$ = !0, n !== on) throw Error("CSSResult is not constructable. Use `unsafeCSS` or `css` instead.");
    this.cssText = t, this.t = e;
  }
  get styleSheet() {
    let t = this.o;
    const e = this.t;
    if (De && t === void 0) {
      const n = e !== void 0 && e.length === 1;
      n && (t = Ue.get(e)), t === void 0 && ((this.o = t = new CSSStyleSheet()).replaceSync(this.cssText), n && Ue.set(e, t));
    }
    return t;
  }
  toString() {
    return this.cssText;
  }
};
const Dn = (i) => new An(typeof i == "string" ? i : i + "", void 0, on), Cn = (i, t) => {
  if (De) i.adoptedStyleSheets = t.map((e) => e instanceof CSSStyleSheet ? e : e.styleSheet);
  else for (const e of t) {
    const n = document.createElement("style"), o = Kt.litNonce;
    o !== void 0 && n.setAttribute("nonce", o), n.textContent = e.cssText, i.appendChild(n);
  }
}, Fe = De ? (i) => i : (i) => i instanceof CSSStyleSheet ? ((t) => {
  let e = "";
  for (const n of t.cssRules) e += n.cssText;
  return Dn(e);
})(i) : i;
/**
 * @license
 * Copyright 2017 Google LLC
 * SPDX-License-Identifier: BSD-3-Clause
 */
const { is: Tn, defineProperty: On, getOwnPropertyDescriptor: Pn, getOwnPropertyNames: xn, getOwnPropertySymbols: In, getPrototypeOf: Nn } = Object, he = globalThis, ke = he.trustedTypes, Mn = ke ? ke.emptyScript : "", Rn = he.reactiveElementPolyfillSupport, xt = (i, t) => i, ne = { toAttribute(i, t) {
  switch (t) {
    case Boolean:
      i = i ? Mn : null;
      break;
    case Object:
    case Array:
      i = i == null ? i : JSON.stringify(i);
  }
  return i;
}, fromAttribute(i, t) {
  let e = i;
  switch (t) {
    case Boolean:
      e = i !== null;
      break;
    case Number:
      e = i === null ? null : Number(i);
      break;
    case Object:
    case Array:
      try {
        e = JSON.parse(i);
      } catch {
        e = null;
      }
  }
  return e;
} }, Ce = (i, t) => !Tn(i, t), Be = { attribute: !0, type: String, converter: ne, reflect: !1, useDefault: !1, hasChanged: Ce };
Symbol.metadata ??= Symbol("metadata"), he.litPropertyMetadata ??= /* @__PURE__ */ new WeakMap();
let pt = class extends HTMLElement {
  static addInitializer(t) {
    this._$Ei(), (this.l ??= []).push(t);
  }
  static get observedAttributes() {
    return this.finalize(), this._$Eh && [...this._$Eh.keys()];
  }
  static createProperty(t, e = Be) {
    if (e.state && (e.attribute = !1), this._$Ei(), this.prototype.hasOwnProperty(t) && ((e = Object.create(e)).wrapped = !0), this.elementProperties.set(t, e), !e.noAccessor) {
      const n = Symbol(), o = this.getPropertyDescriptor(t, n, e);
      o !== void 0 && On(this.prototype, t, o);
    }
  }
  static getPropertyDescriptor(t, e, n) {
    const { get: o, set: r } = Pn(this.prototype, t) ?? { get() {
      return this[e];
    }, set(a) {
      this[e] = a;
    } };
    return { get: o, set(a) {
      const l = o?.call(this);
      r?.call(this, a), this.requestUpdate(t, l, n);
    }, configurable: !0, enumerable: !0 };
  }
  static getPropertyOptions(t) {
    return this.elementProperties.get(t) ?? Be;
  }
  static _$Ei() {
    if (this.hasOwnProperty(xt("elementProperties"))) return;
    const t = Nn(this);
    t.finalize(), t.l !== void 0 && (this.l = [...t.l]), this.elementProperties = new Map(t.elementProperties);
  }
  static finalize() {
    if (this.hasOwnProperty(xt("finalized"))) return;
    if (this.finalized = !0, this._$Ei(), this.hasOwnProperty(xt("properties"))) {
      const e = this.properties, n = [...xn(e), ...In(e)];
      for (const o of n) this.createProperty(o, e[o]);
    }
    const t = this[Symbol.metadata];
    if (t !== null) {
      const e = litPropertyMetadata.get(t);
      if (e !== void 0) for (const [n, o] of e) this.elementProperties.set(n, o);
    }
    this._$Eh = /* @__PURE__ */ new Map();
    for (const [e, n] of this.elementProperties) {
      const o = this._$Eu(e, n);
      o !== void 0 && this._$Eh.set(o, e);
    }
    this.elementStyles = this.finalizeStyles(this.styles);
  }
  static finalizeStyles(t) {
    const e = [];
    if (Array.isArray(t)) {
      const n = new Set(t.flat(1 / 0).reverse());
      for (const o of n) e.unshift(Fe(o));
    } else t !== void 0 && e.push(Fe(t));
    return e;
  }
  static _$Eu(t, e) {
    const n = e.attribute;
    return n === !1 ? void 0 : typeof n == "string" ? n : typeof t == "string" ? t.toLowerCase() : void 0;
  }
  constructor() {
    super(), this._$Ep = void 0, this.isUpdatePending = !1, this.hasUpdated = !1, this._$Em = null, this._$Ev();
  }
  _$Ev() {
    this._$ES = new Promise((t) => this.enableUpdating = t), this._$AL = /* @__PURE__ */ new Map(), this._$E_(), this.requestUpdate(), this.constructor.l?.forEach((t) => t(this));
  }
  addController(t) {
    (this._$EO ??= /* @__PURE__ */ new Set()).add(t), this.renderRoot !== void 0 && this.isConnected && t.hostConnected?.();
  }
  removeController(t) {
    this._$EO?.delete(t);
  }
  _$E_() {
    const t = /* @__PURE__ */ new Map(), e = this.constructor.elementProperties;
    for (const n of e.keys()) this.hasOwnProperty(n) && (t.set(n, this[n]), delete this[n]);
    t.size > 0 && (this._$Ep = t);
  }
  createRenderRoot() {
    const t = this.shadowRoot ?? this.attachShadow(this.constructor.shadowRootOptions);
    return Cn(t, this.constructor.elementStyles), t;
  }
  connectedCallback() {
    this.renderRoot ??= this.createRenderRoot(), this.enableUpdating(!0), this._$EO?.forEach((t) => t.hostConnected?.());
  }
  enableUpdating(t) {
  }
  disconnectedCallback() {
    this._$EO?.forEach((t) => t.hostDisconnected?.());
  }
  attributeChangedCallback(t, e, n) {
    this._$AK(t, n);
  }
  _$ET(t, e) {
    const n = this.constructor.elementProperties.get(t), o = this.constructor._$Eu(t, n);
    if (o !== void 0 && n.reflect === !0) {
      const r = (n.converter?.toAttribute !== void 0 ? n.converter : ne).toAttribute(e, n.type);
      this._$Em = t, r == null ? this.removeAttribute(o) : this.setAttribute(o, r), this._$Em = null;
    }
  }
  _$AK(t, e) {
    const n = this.constructor, o = n._$Eh.get(t);
    if (o !== void 0 && this._$Em !== o) {
      const r = n.getPropertyOptions(o), a = typeof r.converter == "function" ? { fromAttribute: r.converter } : r.converter?.fromAttribute !== void 0 ? r.converter : ne;
      this._$Em = o;
      const l = a.fromAttribute(e, r.type);
      this[o] = l ?? this._$Ej?.get(o) ?? l, this._$Em = null;
    }
  }
  requestUpdate(t, e, n) {
    if (t !== void 0) {
      const o = this.constructor, r = this[t];
      if (n ??= o.getPropertyOptions(t), !((n.hasChanged ?? Ce)(r, e) || n.useDefault && n.reflect && r === this._$Ej?.get(t) && !this.hasAttribute(o._$Eu(t, n)))) return;
      this.C(t, e, n);
    }
    this.isUpdatePending === !1 && (this._$ES = this._$EP());
  }
  C(t, e, { useDefault: n, reflect: o, wrapped: r }, a) {
    n && !(this._$Ej ??= /* @__PURE__ */ new Map()).has(t) && (this._$Ej.set(t, a ?? e ?? this[t]), r !== !0 || a !== void 0) || (this._$AL.has(t) || (this.hasUpdated || n || (e = void 0), this._$AL.set(t, e)), o === !0 && this._$Em !== t && (this._$Eq ??= /* @__PURE__ */ new Set()).add(t));
  }
  async _$EP() {
    this.isUpdatePending = !0;
    try {
      await this._$ES;
    } catch (e) {
      Promise.reject(e);
    }
    const t = this.scheduleUpdate();
    return t != null && await t, !this.isUpdatePending;
  }
  scheduleUpdate() {
    return this.performUpdate();
  }
  performUpdate() {
    if (!this.isUpdatePending) return;
    if (!this.hasUpdated) {
      if (this.renderRoot ??= this.createRenderRoot(), this._$Ep) {
        for (const [o, r] of this._$Ep) this[o] = r;
        this._$Ep = void 0;
      }
      const n = this.constructor.elementProperties;
      if (n.size > 0) for (const [o, r] of n) {
        const { wrapped: a } = r, l = this[o];
        a !== !0 || this._$AL.has(o) || l === void 0 || this.C(o, void 0, r, l);
      }
    }
    let t = !1;
    const e = this._$AL;
    try {
      t = this.shouldUpdate(e), t ? (this.willUpdate(e), this._$EO?.forEach((n) => n.hostUpdate?.()), this.update(e)) : this._$EM();
    } catch (n) {
      throw t = !1, this._$EM(), n;
    }
    t && this._$AE(e);
  }
  willUpdate(t) {
  }
  _$AE(t) {
    this._$EO?.forEach((e) => e.hostUpdated?.()), this.hasUpdated || (this.hasUpdated = !0, this.firstUpdated(t)), this.updated(t);
  }
  _$EM() {
    this._$AL = /* @__PURE__ */ new Map(), this.isUpdatePending = !1;
  }
  get updateComplete() {
    return this.getUpdateComplete();
  }
  getUpdateComplete() {
    return this._$ES;
  }
  shouldUpdate(t) {
    return !0;
  }
  update(t) {
    this._$Eq &&= this._$Eq.forEach((e) => this._$ET(e, this[e])), this._$EM();
  }
  updated(t) {
  }
  firstUpdated(t) {
  }
};
pt.elementStyles = [], pt.shadowRootOptions = { mode: "open" }, pt[xt("elementProperties")] = /* @__PURE__ */ new Map(), pt[xt("finalized")] = /* @__PURE__ */ new Map(), Rn?.({ ReactiveElement: pt }), (he.reactiveElementVersions ??= []).push("2.1.1");
/**
 * @license
 * Copyright 2017 Google LLC
 * SPDX-License-Identifier: BSD-3-Clause
 */
const Te = globalThis, ie = Te.trustedTypes, je = ie ? ie.createPolicy("lit-html", { createHTML: (i) => i }) : void 0, rn = "$lit$", tt = `lit$${Math.random().toFixed(9).slice(2)}$`, an = "?" + tt, Hn = `<${an}>`, ct = document, Ft = () => ct.createComment(""), kt = (i) => i === null || typeof i != "object" && typeof i != "function", Oe = Array.isArray, Un = (i) => Oe(i) || typeof i?.[Symbol.iterator] == "function", de = `[ 	
\f\r]`, Ct = /<(?:(!--|\/[^a-zA-Z])|(\/?[a-zA-Z][^>\s]*)|(\/?$))/g, Le = /-->/g, Xe = />/g, at = RegExp(`>|${de}(?:([^\\s"'>=/]+)(${de}*=${de}*(?:[^ 	
\f\r"'\`<>=]|("|')|))|$)`, "g"), Ye = /'/g, ze = /"/g, sn = /^(?:script|style|textarea|title)$/i, Fn = (i) => (t, ...e) => ({ _$litType$: i, strings: t, values: e }), We = Fn(1), _t = Symbol.for("lit-noChange"), C = Symbol.for("lit-nothing"), Ge = /* @__PURE__ */ new WeakMap(), ut = ct.createTreeWalker(ct, 129);
function ln(i, t) {
  if (!Oe(i) || !i.hasOwnProperty("raw")) throw Error("invalid template strings array");
  return je !== void 0 ? je.createHTML(t) : t;
}
const kn = (i, t) => {
  const e = i.length - 1, n = [];
  let o, r = t === 2 ? "<svg>" : t === 3 ? "<math>" : "", a = Ct;
  for (let l = 0; l < e; l++) {
    const s = i[l];
    let h, d, u = -1, g = 0;
    for (; g < s.length && (a.lastIndex = g, d = a.exec(s), d !== null); ) g = a.lastIndex, a === Ct ? d[1] === "!--" ? a = Le : d[1] !== void 0 ? a = Xe : d[2] !== void 0 ? (sn.test(d[2]) && (o = RegExp("</" + d[2], "g")), a = at) : d[3] !== void 0 && (a = at) : a === at ? d[0] === ">" ? (a = o ?? Ct, u = -1) : d[1] === void 0 ? u = -2 : (u = a.lastIndex - d[2].length, h = d[1], a = d[3] === void 0 ? at : d[3] === '"' ? ze : Ye) : a === ze || a === Ye ? a = at : a === Le || a === Xe ? a = Ct : (a = at, o = void 0);
    const _ = a === at && i[l + 1].startsWith("/>") ? " " : "";
    r += a === Ct ? s + Hn : u >= 0 ? (n.push(h), s.slice(0, u) + rn + s.slice(u) + tt + _) : s + tt + (u === -2 ? l : _);
  }
  return [ln(i, r + (i[e] || "<?>") + (t === 2 ? "</svg>" : t === 3 ? "</math>" : "")), n];
};
class Bt {
  constructor({ strings: t, _$litType$: e }, n) {
    let o;
    this.parts = [];
    let r = 0, a = 0;
    const l = t.length - 1, s = this.parts, [h, d] = kn(t, e);
    if (this.el = Bt.createElement(h, n), ut.currentNode = this.el.content, e === 2 || e === 3) {
      const u = this.el.content.firstChild;
      u.replaceWith(...u.childNodes);
    }
    for (; (o = ut.nextNode()) !== null && s.length < l; ) {
      if (o.nodeType === 1) {
        if (o.hasAttributes()) for (const u of o.getAttributeNames()) if (u.endsWith(rn)) {
          const g = d[a++], _ = o.getAttribute(u).split(tt), v = /([.?@])?(.*)/.exec(g);
          s.push({ type: 1, index: r, name: v[2], strings: _, ctor: v[1] === "." ? jn : v[1] === "?" ? Ln : v[1] === "@" ? Xn : ue }), o.removeAttribute(u);
        } else u.startsWith(tt) && (s.push({ type: 6, index: r }), o.removeAttribute(u));
        if (sn.test(o.tagName)) {
          const u = o.textContent.split(tt), g = u.length - 1;
          if (g > 0) {
            o.textContent = ie ? ie.emptyScript : "";
            for (let _ = 0; _ < g; _++) o.append(u[_], Ft()), ut.nextNode(), s.push({ type: 2, index: ++r });
            o.append(u[g], Ft());
          }
        }
      } else if (o.nodeType === 8) if (o.data === an) s.push({ type: 2, index: r });
      else {
        let u = -1;
        for (; (u = o.data.indexOf(tt, u + 1)) !== -1; ) s.push({ type: 7, index: r }), u += tt.length - 1;
      }
      r++;
    }
  }
  static createElement(t, e) {
    const n = ct.createElement("template");
    return n.innerHTML = t, n;
  }
}
function bt(i, t, e = i, n) {
  if (t === _t) return t;
  let o = n !== void 0 ? e._$Co?.[n] : e._$Cl;
  const r = kt(t) ? void 0 : t._$litDirective$;
  return o?.constructor !== r && (o?._$AO?.(!1), r === void 0 ? o = void 0 : (o = new r(i), o._$AT(i, e, n)), n !== void 0 ? (e._$Co ??= [])[n] = o : e._$Cl = o), o !== void 0 && (t = bt(i, o._$AS(i, t.values), o, n)), t;
}
class Bn {
  constructor(t, e) {
    this._$AV = [], this._$AN = void 0, this._$AD = t, this._$AM = e;
  }
  get parentNode() {
    return this._$AM.parentNode;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  u(t) {
    const { el: { content: e }, parts: n } = this._$AD, o = (t?.creationScope ?? ct).importNode(e, !0);
    ut.currentNode = o;
    let r = ut.nextNode(), a = 0, l = 0, s = n[0];
    for (; s !== void 0; ) {
      if (a === s.index) {
        let h;
        s.type === 2 ? h = new jt(r, r.nextSibling, this, t) : s.type === 1 ? h = new s.ctor(r, s.name, s.strings, this, t) : s.type === 6 && (h = new Yn(r, this, t)), this._$AV.push(h), s = n[++l];
      }
      a !== s?.index && (r = ut.nextNode(), a++);
    }
    return ut.currentNode = ct, o;
  }
  p(t) {
    let e = 0;
    for (const n of this._$AV) n !== void 0 && (n.strings !== void 0 ? (n._$AI(t, n, e), e += n.strings.length - 2) : n._$AI(t[e])), e++;
  }
}
class jt {
  get _$AU() {
    return this._$AM?._$AU ?? this._$Cv;
  }
  constructor(t, e, n, o) {
    this.type = 2, this._$AH = C, this._$AN = void 0, this._$AA = t, this._$AB = e, this._$AM = n, this.options = o, this._$Cv = o?.isConnected ?? !0;
  }
  get parentNode() {
    let t = this._$AA.parentNode;
    const e = this._$AM;
    return e !== void 0 && t?.nodeType === 11 && (t = e.parentNode), t;
  }
  get startNode() {
    return this._$AA;
  }
  get endNode() {
    return this._$AB;
  }
  _$AI(t, e = this) {
    t = bt(this, t, e), kt(t) ? t === C || t == null || t === "" ? (this._$AH !== C && this._$AR(), this._$AH = C) : t !== this._$AH && t !== _t && this._(t) : t._$litType$ !== void 0 ? this.$(t) : t.nodeType !== void 0 ? this.T(t) : Un(t) ? this.k(t) : this._(t);
  }
  O(t) {
    return this._$AA.parentNode.insertBefore(t, this._$AB);
  }
  T(t) {
    this._$AH !== t && (this._$AR(), this._$AH = this.O(t));
  }
  _(t) {
    this._$AH !== C && kt(this._$AH) ? this._$AA.nextSibling.data = t : this.T(ct.createTextNode(t)), this._$AH = t;
  }
  $(t) {
    const { values: e, _$litType$: n } = t, o = typeof n == "number" ? this._$AC(t) : (n.el === void 0 && (n.el = Bt.createElement(ln(n.h, n.h[0]), this.options)), n);
    if (this._$AH?._$AD === o) this._$AH.p(e);
    else {
      const r = new Bn(o, this), a = r.u(this.options);
      r.p(e), this.T(a), this._$AH = r;
    }
  }
  _$AC(t) {
    let e = Ge.get(t.strings);
    return e === void 0 && Ge.set(t.strings, e = new Bt(t)), e;
  }
  k(t) {
    Oe(this._$AH) || (this._$AH = [], this._$AR());
    const e = this._$AH;
    let n, o = 0;
    for (const r of t) o === e.length ? e.push(n = new jt(this.O(Ft()), this.O(Ft()), this, this.options)) : n = e[o], n._$AI(r), o++;
    o < e.length && (this._$AR(n && n._$AB.nextSibling, o), e.length = o);
  }
  _$AR(t = this._$AA.nextSibling, e) {
    for (this._$AP?.(!1, !0, e); t !== this._$AB; ) {
      const n = t.nextSibling;
      t.remove(), t = n;
    }
  }
  setConnected(t) {
    this._$AM === void 0 && (this._$Cv = t, this._$AP?.(t));
  }
}
class ue {
  get tagName() {
    return this.element.tagName;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  constructor(t, e, n, o, r) {
    this.type = 1, this._$AH = C, this._$AN = void 0, this.element = t, this.name = e, this._$AM = o, this.options = r, n.length > 2 || n[0] !== "" || n[1] !== "" ? (this._$AH = Array(n.length - 1).fill(new String()), this.strings = n) : this._$AH = C;
  }
  _$AI(t, e = this, n, o) {
    const r = this.strings;
    let a = !1;
    if (r === void 0) t = bt(this, t, e, 0), a = !kt(t) || t !== this._$AH && t !== _t, a && (this._$AH = t);
    else {
      const l = t;
      let s, h;
      for (t = r[0], s = 0; s < r.length - 1; s++) h = bt(this, l[n + s], e, s), h === _t && (h = this._$AH[s]), a ||= !kt(h) || h !== this._$AH[s], h === C ? t = C : t !== C && (t += (h ?? "") + r[s + 1]), this._$AH[s] = h;
    }
    a && !o && this.j(t);
  }
  j(t) {
    t === C ? this.element.removeAttribute(this.name) : this.element.setAttribute(this.name, t ?? "");
  }
}
class jn extends ue {
  constructor() {
    super(...arguments), this.type = 3;
  }
  j(t) {
    this.element[this.name] = t === C ? void 0 : t;
  }
}
class Ln extends ue {
  constructor() {
    super(...arguments), this.type = 4;
  }
  j(t) {
    this.element.toggleAttribute(this.name, !!t && t !== C);
  }
}
class Xn extends ue {
  constructor(t, e, n, o, r) {
    super(t, e, n, o, r), this.type = 5;
  }
  _$AI(t, e = this) {
    if ((t = bt(this, t, e, 0) ?? C) === _t) return;
    const n = this._$AH, o = t === C && n !== C || t.capture !== n.capture || t.once !== n.once || t.passive !== n.passive, r = t !== C && (n === C || o);
    o && this.element.removeEventListener(this.name, this, n), r && this.element.addEventListener(this.name, this, t), this._$AH = t;
  }
  handleEvent(t) {
    typeof this._$AH == "function" ? this._$AH.call(this.options?.host ?? this.element, t) : this._$AH.handleEvent(t);
  }
}
class Yn {
  constructor(t, e, n) {
    this.element = t, this.type = 6, this._$AN = void 0, this._$AM = e, this.options = n;
  }
  get _$AU() {
    return this._$AM._$AU;
  }
  _$AI(t) {
    bt(this, t);
  }
}
const zn = Te.litHtmlPolyfillSupport;
zn?.(Bt, jt), (Te.litHtmlVersions ??= []).push("3.3.1");
const Wn = (i, t, e) => {
  const n = e?.renderBefore ?? t;
  let o = n._$litPart$;
  if (o === void 0) {
    const r = e?.renderBefore ?? null;
    n._$litPart$ = o = new jt(t.insertBefore(Ft(), r), r, void 0, e ?? {});
  }
  return o._$AI(i), o;
};
/**
 * @license
 * Copyright 2017 Google LLC
 * SPDX-License-Identifier: BSD-3-Clause
 */
const Pe = globalThis;
class It extends pt {
  constructor() {
    super(...arguments), this.renderOptions = { host: this }, this._$Do = void 0;
  }
  createRenderRoot() {
    const t = super.createRenderRoot();
    return this.renderOptions.renderBefore ??= t.firstChild, t;
  }
  update(t) {
    const e = this.render();
    this.hasUpdated || (this.renderOptions.isConnected = this.isConnected), super.update(t), this._$Do = Wn(e, this.renderRoot, this.renderOptions);
  }
  connectedCallback() {
    super.connectedCallback(), this._$Do?.setConnected(!0);
  }
  disconnectedCallback() {
    super.disconnectedCallback(), this._$Do?.setConnected(!1);
  }
  render() {
    return _t;
  }
}
It._$litElement$ = !0, It.finalized = !0, Pe.litElementHydrateSupport?.({ LitElement: It });
const Gn = Pe.litElementPolyfillSupport;
Gn?.({ LitElement: It });
(Pe.litElementVersions ??= []).push("4.2.1");
/**
 * @license
 * Copyright 2017 Google LLC
 * SPDX-License-Identifier: BSD-3-Clause
 */
const qn = (i) => (t, e) => {
  e !== void 0 ? e.addInitializer(() => {
    customElements.define(i, t);
  }) : customElements.define(i, t);
};
/**
 * @license
 * Copyright 2017 Google LLC
 * SPDX-License-Identifier: BSD-3-Clause
 */
const Vn = { attribute: !0, type: String, converter: ne, reflect: !1, hasChanged: Ce }, Kn = (i = Vn, t, e) => {
  const { kind: n, metadata: o } = e;
  let r = globalThis.litPropertyMetadata.get(o);
  if (r === void 0 && globalThis.litPropertyMetadata.set(o, r = /* @__PURE__ */ new Map()), n === "setter" && ((i = Object.create(i)).wrapped = !0), r.set(e.name, i), n === "accessor") {
    const { name: a } = e;
    return { set(l) {
      const s = t.get.call(this);
      t.set.call(this, l), this.requestUpdate(a, s, i);
    }, init(l) {
      return l !== void 0 && this.C(a, void 0, i, l), l;
    } };
  }
  if (n === "setter") {
    const { name: a } = e;
    return function(l) {
      const s = this[a];
      t.call(this, l), this.requestUpdate(a, s, i);
    };
  }
  throw Error("Unsupported decorator location: " + n);
};
function xe(i) {
  return (t, e) => typeof e == "object" ? Kn(i, t, e) : ((n, o, r) => {
    const a = o.hasOwnProperty(r);
    return o.constructor.createProperty(r, n), a ? Object.getOwnPropertyDescriptor(o, r) : void 0;
  })(i, t, e);
}
/**
 * @license
 * Copyright 2017 Google LLC
 * SPDX-License-Identifier: BSD-3-Clause
 */
const Zn = (i, t, e) => (e.configurable = !0, e.enumerable = !0, Reflect.decorate && typeof t != "object" && Object.defineProperty(i, t, e), e);
/**
 * @license
 * Copyright 2017 Google LLC
 * SPDX-License-Identifier: BSD-3-Clause
 */
function Jn(i, t) {
  return (e, n, o) => {
    const r = (a) => a.renderRoot?.querySelector(i) ?? null;
    return Zn(e, n, { get() {
      return r(this);
    } });
  };
}
/**!
 * Sortable 1.15.6
 * @author	RubaXa   <trash@rubaxa.org>
 * @author	owenm    <owen23355@gmail.com>
 * @license MIT
 */
function qe(i, t) {
  var e = Object.keys(i);
  if (Object.getOwnPropertySymbols) {
    var n = Object.getOwnPropertySymbols(i);
    t && (n = n.filter(function(o) {
      return Object.getOwnPropertyDescriptor(i, o).enumerable;
    })), e.push.apply(e, n);
  }
  return e;
}
function z(i) {
  for (var t = 1; t < arguments.length; t++) {
    var e = arguments[t] != null ? arguments[t] : {};
    t % 2 ? qe(Object(e), !0).forEach(function(n) {
      Qn(i, n, e[n]);
    }) : Object.getOwnPropertyDescriptors ? Object.defineProperties(i, Object.getOwnPropertyDescriptors(e)) : qe(Object(e)).forEach(function(n) {
      Object.defineProperty(i, n, Object.getOwnPropertyDescriptor(e, n));
    });
  }
  return i;
}
function Zt(i) {
  "@babel/helpers - typeof";
  return typeof Symbol == "function" && typeof Symbol.iterator == "symbol" ? Zt = function(t) {
    return typeof t;
  } : Zt = function(t) {
    return t && typeof Symbol == "function" && t.constructor === Symbol && t !== Symbol.prototype ? "symbol" : typeof t;
  }, Zt(i);
}
function Qn(i, t, e) {
  return t in i ? Object.defineProperty(i, t, {
    value: e,
    enumerable: !0,
    configurable: !0,
    writable: !0
  }) : i[t] = e, i;
}
function V() {
  return V = Object.assign || function(i) {
    for (var t = 1; t < arguments.length; t++) {
      var e = arguments[t];
      for (var n in e)
        Object.prototype.hasOwnProperty.call(e, n) && (i[n] = e[n]);
    }
    return i;
  }, V.apply(this, arguments);
}
function ti(i, t) {
  if (i == null) return {};
  var e = {}, n = Object.keys(i), o, r;
  for (r = 0; r < n.length; r++)
    o = n[r], !(t.indexOf(o) >= 0) && (e[o] = i[o]);
  return e;
}
function ei(i, t) {
  if (i == null) return {};
  var e = ti(i, t), n, o;
  if (Object.getOwnPropertySymbols) {
    var r = Object.getOwnPropertySymbols(i);
    for (o = 0; o < r.length; o++)
      n = r[o], !(t.indexOf(n) >= 0) && Object.prototype.propertyIsEnumerable.call(i, n) && (e[n] = i[n]);
  }
  return e;
}
var ni = "1.15.6";
function q(i) {
  if (typeof window < "u" && window.navigator)
    return !!/* @__PURE__ */ navigator.userAgent.match(i);
}
var K = q(/(?:Trident.*rv[ :]?11\.|msie|iemobile|Windows Phone)/i), Lt = q(/Edge/i), Ve = q(/firefox/i), Nt = q(/safari/i) && !q(/chrome/i) && !q(/android/i), Ie = q(/iP(ad|od|hone)/i), hn = q(/chrome/i) && q(/android/i), un = {
  capture: !1,
  passive: !1
};
function y(i, t, e) {
  i.addEventListener(t, e, !K && un);
}
function b(i, t, e) {
  i.removeEventListener(t, e, !K && un);
}
function oe(i, t) {
  if (t) {
    if (t[0] === ">" && (t = t.substring(1)), i)
      try {
        if (i.matches)
          return i.matches(t);
        if (i.msMatchesSelector)
          return i.msMatchesSelector(t);
        if (i.webkitMatchesSelector)
          return i.webkitMatchesSelector(t);
      } catch {
        return !1;
      }
    return !1;
  }
}
function cn(i) {
  return i.host && i !== document && i.host.nodeType ? i.host : i.parentNode;
}
function L(i, t, e, n) {
  if (i) {
    e = e || document;
    do {
      if (t != null && (t[0] === ">" ? i.parentNode === e && oe(i, t) : oe(i, t)) || n && i === e)
        return i;
      if (i === e) break;
    } while (i = cn(i));
  }
  return null;
}
var Ke = /\s+/g;
function H(i, t, e) {
  if (i && t)
    if (i.classList)
      i.classList[e ? "add" : "remove"](t);
    else {
      var n = (" " + i.className + " ").replace(Ke, " ").replace(" " + t + " ", " ");
      i.className = (n + (e ? " " + t : "")).replace(Ke, " ");
    }
}
function f(i, t, e) {
  var n = i && i.style;
  if (n) {
    if (e === void 0)
      return document.defaultView && document.defaultView.getComputedStyle ? e = document.defaultView.getComputedStyle(i, "") : i.currentStyle && (e = i.currentStyle), t === void 0 ? e : e[t];
    !(t in n) && t.indexOf("webkit") === -1 && (t = "-webkit-" + t), n[t] = e + (typeof e == "string" ? "" : "px");
  }
}
function vt(i, t) {
  var e = "";
  if (typeof i == "string")
    e = i;
  else
    do {
      var n = f(i, "transform");
      n && n !== "none" && (e = n + " " + e);
    } while (!t && (i = i.parentNode));
  var o = window.DOMMatrix || window.WebKitCSSMatrix || window.CSSMatrix || window.MSCSSMatrix;
  return o && new o(e);
}
function dn(i, t, e) {
  if (i) {
    var n = i.getElementsByTagName(t), o = 0, r = n.length;
    if (e)
      for (; o < r; o++)
        e(n[o], o);
    return n;
  }
  return [];
}
function Y() {
  var i = document.scrollingElement;
  return i || document.documentElement;
}
function D(i, t, e, n, o) {
  if (!(!i.getBoundingClientRect && i !== window)) {
    var r, a, l, s, h, d, u;
    if (i !== window && i.parentNode && i !== Y() ? (r = i.getBoundingClientRect(), a = r.top, l = r.left, s = r.bottom, h = r.right, d = r.height, u = r.width) : (a = 0, l = 0, s = window.innerHeight, h = window.innerWidth, d = window.innerHeight, u = window.innerWidth), (t || e) && i !== window && (o = o || i.parentNode, !K))
      do
        if (o && o.getBoundingClientRect && (f(o, "transform") !== "none" || e && f(o, "position") !== "static")) {
          var g = o.getBoundingClientRect();
          a -= g.top + parseInt(f(o, "border-top-width")), l -= g.left + parseInt(f(o, "border-left-width")), s = a + r.height, h = l + r.width;
          break;
        }
      while (o = o.parentNode);
    if (n && i !== window) {
      var _ = vt(o || i), v = _ && _.a, E = _ && _.d;
      _ && (a /= E, l /= v, u /= v, d /= E, s = a + d, h = l + u);
    }
    return {
      top: a,
      left: l,
      bottom: s,
      right: h,
      width: u,
      height: d
    };
  }
}
function Ze(i, t, e) {
  for (var n = nt(i, !0), o = D(i)[t]; n; ) {
    var r = D(n)[e], a = void 0;
    if (a = o >= r, !a) return n;
    if (n === Y()) break;
    n = nt(n, !1);
  }
  return !1;
}
function yt(i, t, e, n) {
  for (var o = 0, r = 0, a = i.children; r < a.length; ) {
    if (a[r].style.display !== "none" && a[r] !== p.ghost && (n || a[r] !== p.dragged) && L(a[r], e.draggable, i, !1)) {
      if (o === t)
        return a[r];
      o++;
    }
    r++;
  }
  return null;
}
function Ne(i, t) {
  for (var e = i.lastElementChild; e && (e === p.ghost || f(e, "display") === "none" || t && !oe(e, t)); )
    e = e.previousElementSibling;
  return e || null;
}
function F(i, t) {
  var e = 0;
  if (!i || !i.parentNode)
    return -1;
  for (; i = i.previousElementSibling; )
    i.nodeName.toUpperCase() !== "TEMPLATE" && i !== p.clone && (!t || oe(i, t)) && e++;
  return e;
}
function Je(i) {
  var t = 0, e = 0, n = Y();
  if (i)
    do {
      var o = vt(i), r = o.a, a = o.d;
      t += i.scrollLeft * r, e += i.scrollTop * a;
    } while (i !== n && (i = i.parentNode));
  return [t, e];
}
function ii(i, t) {
  for (var e in i)
    if (i.hasOwnProperty(e)) {
      for (var n in t)
        if (t.hasOwnProperty(n) && t[n] === i[e][n]) return Number(e);
    }
  return -1;
}
function nt(i, t) {
  if (!i || !i.getBoundingClientRect) return Y();
  var e = i, n = !1;
  do
    if (e.clientWidth < e.scrollWidth || e.clientHeight < e.scrollHeight) {
      var o = f(e);
      if (e.clientWidth < e.scrollWidth && (o.overflowX == "auto" || o.overflowX == "scroll") || e.clientHeight < e.scrollHeight && (o.overflowY == "auto" || o.overflowY == "scroll")) {
        if (!e.getBoundingClientRect || e === document.body) return Y();
        if (n || t) return e;
        n = !0;
      }
    }
  while (e = e.parentNode);
  return Y();
}
function oi(i, t) {
  if (i && t)
    for (var e in t)
      t.hasOwnProperty(e) && (i[e] = t[e]);
  return i;
}
function fe(i, t) {
  return Math.round(i.top) === Math.round(t.top) && Math.round(i.left) === Math.round(t.left) && Math.round(i.height) === Math.round(t.height) && Math.round(i.width) === Math.round(t.width);
}
var Mt;
function fn(i, t) {
  return function() {
    if (!Mt) {
      var e = arguments, n = this;
      e.length === 1 ? i.call(n, e[0]) : i.apply(n, e), Mt = setTimeout(function() {
        Mt = void 0;
      }, t);
    }
  };
}
function ri() {
  clearTimeout(Mt), Mt = void 0;
}
function pn(i, t, e) {
  i.scrollLeft += t, i.scrollTop += e;
}
function gn(i) {
  var t = window.Polymer, e = window.jQuery || window.Zepto;
  return t && t.dom ? t.dom(i).cloneNode(!0) : e ? e(i).clone(!0)[0] : i.cloneNode(!0);
}
function mn(i, t, e) {
  var n = {};
  return Array.from(i.children).forEach(function(o) {
    var r, a, l, s;
    if (!(!L(o, t.draggable, i, !1) || o.animated || o === e)) {
      var h = D(o);
      n.left = Math.min((r = n.left) !== null && r !== void 0 ? r : 1 / 0, h.left), n.top = Math.min((a = n.top) !== null && a !== void 0 ? a : 1 / 0, h.top), n.right = Math.max((l = n.right) !== null && l !== void 0 ? l : -1 / 0, h.right), n.bottom = Math.max((s = n.bottom) !== null && s !== void 0 ? s : -1 / 0, h.bottom);
    }
  }), n.width = n.right - n.left, n.height = n.bottom - n.top, n.x = n.left, n.y = n.top, n;
}
var N = "Sortable" + (/* @__PURE__ */ new Date()).getTime();
function ai() {
  var i = [], t;
  return {
    captureAnimationState: function() {
      if (i = [], !!this.options.animation) {
        var n = [].slice.call(this.el.children);
        n.forEach(function(o) {
          if (!(f(o, "display") === "none" || o === p.ghost)) {
            i.push({
              target: o,
              rect: D(o)
            });
            var r = z({}, i[i.length - 1].rect);
            if (o.thisAnimationDuration) {
              var a = vt(o, !0);
              a && (r.top -= a.f, r.left -= a.e);
            }
            o.fromRect = r;
          }
        });
      }
    },
    addAnimationState: function(n) {
      i.push(n);
    },
    removeAnimationState: function(n) {
      i.splice(ii(i, {
        target: n
      }), 1);
    },
    animateAll: function(n) {
      var o = this;
      if (!this.options.animation) {
        clearTimeout(t), typeof n == "function" && n();
        return;
      }
      var r = !1, a = 0;
      i.forEach(function(l) {
        var s = 0, h = l.target, d = h.fromRect, u = D(h), g = h.prevFromRect, _ = h.prevToRect, v = l.rect, E = vt(h, !0);
        E && (u.top -= E.f, u.left -= E.e), h.toRect = u, h.thisAnimationDuration && fe(g, u) && !fe(d, u) && // Make sure animatingRect is on line between toRect & fromRect
        (v.top - u.top) / (v.left - u.left) === (d.top - u.top) / (d.left - u.left) && (s = li(v, g, _, o.options)), fe(u, d) || (h.prevFromRect = d, h.prevToRect = u, s || (s = o.options.animation), o.animate(h, v, u, s)), s && (r = !0, a = Math.max(a, s), clearTimeout(h.animationResetTimer), h.animationResetTimer = setTimeout(function() {
          h.animationTime = 0, h.prevFromRect = null, h.fromRect = null, h.prevToRect = null, h.thisAnimationDuration = null;
        }, s), h.thisAnimationDuration = s);
      }), clearTimeout(t), r ? t = setTimeout(function() {
        typeof n == "function" && n();
      }, a) : typeof n == "function" && n(), i = [];
    },
    animate: function(n, o, r, a) {
      if (a) {
        f(n, "transition", ""), f(n, "transform", "");
        var l = vt(this.el), s = l && l.a, h = l && l.d, d = (o.left - r.left) / (s || 1), u = (o.top - r.top) / (h || 1);
        n.animatingX = !!d, n.animatingY = !!u, f(n, "transform", "translate3d(" + d + "px," + u + "px,0)"), this.forRepaintDummy = si(n), f(n, "transition", "transform " + a + "ms" + (this.options.easing ? " " + this.options.easing : "")), f(n, "transform", "translate3d(0,0,0)"), typeof n.animated == "number" && clearTimeout(n.animated), n.animated = setTimeout(function() {
          f(n, "transition", ""), f(n, "transform", ""), n.animated = !1, n.animatingX = !1, n.animatingY = !1;
        }, a);
      }
    }
  };
}
function si(i) {
  return i.offsetWidth;
}
function li(i, t, e, n) {
  return Math.sqrt(Math.pow(t.top - i.top, 2) + Math.pow(t.left - i.left, 2)) / Math.sqrt(Math.pow(t.top - e.top, 2) + Math.pow(t.left - e.left, 2)) * n.animation;
}
var dt = [], pe = {
  initializeByDefault: !0
}, Xt = {
  mount: function(t) {
    for (var e in pe)
      pe.hasOwnProperty(e) && !(e in t) && (t[e] = pe[e]);
    dt.forEach(function(n) {
      if (n.pluginName === t.pluginName)
        throw "Sortable: Cannot mount plugin ".concat(t.pluginName, " more than once");
    }), dt.push(t);
  },
  pluginEvent: function(t, e, n) {
    var o = this;
    this.eventCanceled = !1, n.cancel = function() {
      o.eventCanceled = !0;
    };
    var r = t + "Global";
    dt.forEach(function(a) {
      e[a.pluginName] && (e[a.pluginName][r] && e[a.pluginName][r](z({
        sortable: e
      }, n)), e.options[a.pluginName] && e[a.pluginName][t] && e[a.pluginName][t](z({
        sortable: e
      }, n)));
    });
  },
  initializePlugins: function(t, e, n, o) {
    dt.forEach(function(l) {
      var s = l.pluginName;
      if (!(!t.options[s] && !l.initializeByDefault)) {
        var h = new l(t, e, t.options);
        h.sortable = t, h.options = t.options, t[s] = h, V(n, h.defaults);
      }
    });
    for (var r in t.options)
      if (t.options.hasOwnProperty(r)) {
        var a = this.modifyOption(t, r, t.options[r]);
        typeof a < "u" && (t.options[r] = a);
      }
  },
  getEventProperties: function(t, e) {
    var n = {};
    return dt.forEach(function(o) {
      typeof o.eventProperties == "function" && V(n, o.eventProperties.call(e[o.pluginName], t));
    }), n;
  },
  modifyOption: function(t, e, n) {
    var o;
    return dt.forEach(function(r) {
      t[r.pluginName] && r.optionListeners && typeof r.optionListeners[e] == "function" && (o = r.optionListeners[e].call(t[r.pluginName], n));
    }), o;
  }
};
function hi(i) {
  var t = i.sortable, e = i.rootEl, n = i.name, o = i.targetEl, r = i.cloneEl, a = i.toEl, l = i.fromEl, s = i.oldIndex, h = i.newIndex, d = i.oldDraggableIndex, u = i.newDraggableIndex, g = i.originalEvent, _ = i.putSortable, v = i.extraEventProperties;
  if (t = t || e && e[N], !!t) {
    var E, k = t.options, W = "on" + n.charAt(0).toUpperCase() + n.substr(1);
    window.CustomEvent && !K && !Lt ? E = new CustomEvent(n, {
      bubbles: !0,
      cancelable: !0
    }) : (E = document.createEvent("Event"), E.initEvent(n, !0, !0)), E.to = a || e, E.from = l || e, E.item = o || e, E.clone = r, E.oldIndex = s, E.newIndex = h, E.oldDraggableIndex = d, E.newDraggableIndex = u, E.originalEvent = g, E.pullMode = _ ? _.lastPutMode : void 0;
    var P = z(z({}, v), Xt.getEventProperties(n, t));
    for (var B in P)
      E[B] = P[B];
    e && e.dispatchEvent(E), k[W] && k[W].call(t, E);
  }
}
var ui = ["evt"], I = function(t, e) {
  var n = arguments.length > 2 && arguments[2] !== void 0 ? arguments[2] : {}, o = n.evt, r = ei(n, ui);
  Xt.pluginEvent.bind(p)(t, e, z({
    dragEl: c,
    parentEl: S,
    ghostEl: m,
    rootEl: w,
    nextEl: ht,
    lastDownEl: Jt,
    cloneEl: $,
    cloneHidden: et,
    dragStarted: Tt,
    putSortable: T,
    activeSortable: p.active,
    originalEvent: o,
    oldIndex: mt,
    oldDraggableIndex: Rt,
    newIndex: U,
    newDraggableIndex: Q,
    hideGhostForTarget: yn,
    unhideGhostForTarget: En,
    cloneNowHidden: function() {
      et = !0;
    },
    cloneNowShown: function() {
      et = !1;
    },
    dispatchSortableEvent: function(l) {
      x({
        sortable: e,
        name: l,
        originalEvent: o
      });
    }
  }, r));
};
function x(i) {
  hi(z({
    putSortable: T,
    cloneEl: $,
    targetEl: c,
    rootEl: w,
    oldIndex: mt,
    oldDraggableIndex: Rt,
    newIndex: U,
    newDraggableIndex: Q
  }, i));
}
var c, S, m, w, ht, Jt, $, et, mt, U, Rt, Q, Wt, T, gt = !1, re = !1, ae = [], st, j, ge, me, Qe, tn, Tt, ft, Ht, Ut = !1, Gt = !1, Qt, O, ve = [], we = !1, se = [], ce = typeof document < "u", qt = Ie, en = Lt || K ? "cssFloat" : "float", ci = ce && !hn && !Ie && "draggable" in document.createElement("div"), vn = function() {
  if (ce) {
    if (K)
      return !1;
    var i = document.createElement("x");
    return i.style.cssText = "pointer-events:auto", i.style.pointerEvents === "auto";
  }
}(), _n = function(t, e) {
  var n = f(t), o = parseInt(n.width) - parseInt(n.paddingLeft) - parseInt(n.paddingRight) - parseInt(n.borderLeftWidth) - parseInt(n.borderRightWidth), r = yt(t, 0, e), a = yt(t, 1, e), l = r && f(r), s = a && f(a), h = l && parseInt(l.marginLeft) + parseInt(l.marginRight) + D(r).width, d = s && parseInt(s.marginLeft) + parseInt(s.marginRight) + D(a).width;
  if (n.display === "flex")
    return n.flexDirection === "column" || n.flexDirection === "column-reverse" ? "vertical" : "horizontal";
  if (n.display === "grid")
    return n.gridTemplateColumns.split(" ").length <= 1 ? "vertical" : "horizontal";
  if (r && l.float && l.float !== "none") {
    var u = l.float === "left" ? "left" : "right";
    return a && (s.clear === "both" || s.clear === u) ? "vertical" : "horizontal";
  }
  return r && (l.display === "block" || l.display === "flex" || l.display === "table" || l.display === "grid" || h >= o && n[en] === "none" || a && n[en] === "none" && h + d > o) ? "vertical" : "horizontal";
}, di = function(t, e, n) {
  var o = n ? t.left : t.top, r = n ? t.right : t.bottom, a = n ? t.width : t.height, l = n ? e.left : e.top, s = n ? e.right : e.bottom, h = n ? e.width : e.height;
  return o === l || r === s || o + a / 2 === l + h / 2;
}, fi = function(t, e) {
  var n;
  return ae.some(function(o) {
    var r = o[N].options.emptyInsertThreshold;
    if (!(!r || Ne(o))) {
      var a = D(o), l = t >= a.left - r && t <= a.right + r, s = e >= a.top - r && e <= a.bottom + r;
      if (l && s)
        return n = o;
    }
  }), n;
}, bn = function(t) {
  function e(r, a) {
    return function(l, s, h, d) {
      var u = l.options.group.name && s.options.group.name && l.options.group.name === s.options.group.name;
      if (r == null && (a || u))
        return !0;
      if (r == null || r === !1)
        return !1;
      if (a && r === "clone")
        return r;
      if (typeof r == "function")
        return e(r(l, s, h, d), a)(l, s, h, d);
      var g = (a ? l : s).options.group.name;
      return r === !0 || typeof r == "string" && r === g || r.join && r.indexOf(g) > -1;
    };
  }
  var n = {}, o = t.group;
  (!o || Zt(o) != "object") && (o = {
    name: o
  }), n.name = o.name, n.checkPull = e(o.pull, !0), n.checkPut = e(o.put), n.revertClone = o.revertClone, t.group = n;
}, yn = function() {
  !vn && m && f(m, "display", "none");
}, En = function() {
  !vn && m && f(m, "display", "");
};
ce && !hn && document.addEventListener("click", function(i) {
  if (re)
    return i.preventDefault(), i.stopPropagation && i.stopPropagation(), i.stopImmediatePropagation && i.stopImmediatePropagation(), re = !1, !1;
}, !0);
var lt = function(t) {
  if (c) {
    t = t.touches ? t.touches[0] : t;
    var e = fi(t.clientX, t.clientY);
    if (e) {
      var n = {};
      for (var o in t)
        t.hasOwnProperty(o) && (n[o] = t[o]);
      n.target = n.rootEl = e, n.preventDefault = void 0, n.stopPropagation = void 0, e[N]._onDragOver(n);
    }
  }
}, pi = function(t) {
  c && c.parentNode[N]._isOutsideThisEl(t.target);
};
function p(i, t) {
  if (!(i && i.nodeType && i.nodeType === 1))
    throw "Sortable: `el` must be an HTMLElement, not ".concat({}.toString.call(i));
  this.el = i, this.options = t = V({}, t), i[N] = this;
  var e = {
    group: null,
    sort: !0,
    disabled: !1,
    store: null,
    handle: null,
    draggable: /^[uo]l$/i.test(i.nodeName) ? ">li" : ">*",
    swapThreshold: 1,
    // percentage; 0 <= x <= 1
    invertSwap: !1,
    // invert always
    invertedSwapThreshold: null,
    // will be set to same as swapThreshold if default
    removeCloneOnHide: !0,
    direction: function() {
      return _n(i, this.options);
    },
    ghostClass: "sortable-ghost",
    chosenClass: "sortable-chosen",
    dragClass: "sortable-drag",
    ignore: "a, img",
    filter: null,
    preventOnFilter: !0,
    animation: 0,
    easing: null,
    setData: function(a, l) {
      a.setData("Text", l.textContent);
    },
    dropBubble: !1,
    dragoverBubble: !1,
    dataIdAttr: "data-id",
    delay: 0,
    delayOnTouchOnly: !1,
    touchStartThreshold: (Number.parseInt ? Number : window).parseInt(window.devicePixelRatio, 10) || 1,
    forceFallback: !1,
    fallbackClass: "sortable-fallback",
    fallbackOnBody: !1,
    fallbackTolerance: 0,
    fallbackOffset: {
      x: 0,
      y: 0
    },
    // Disabled on Safari: #1571; Enabled on Safari IOS: #2244
    supportPointer: p.supportPointer !== !1 && "PointerEvent" in window && (!Nt || Ie),
    emptyInsertThreshold: 5
  };
  Xt.initializePlugins(this, i, e);
  for (var n in e)
    !(n in t) && (t[n] = e[n]);
  bn(t);
  for (var o in this)
    o.charAt(0) === "_" && typeof this[o] == "function" && (this[o] = this[o].bind(this));
  this.nativeDraggable = t.forceFallback ? !1 : ci, this.nativeDraggable && (this.options.touchStartThreshold = 1), t.supportPointer ? y(i, "pointerdown", this._onTapStart) : (y(i, "mousedown", this._onTapStart), y(i, "touchstart", this._onTapStart)), this.nativeDraggable && (y(i, "dragover", this), y(i, "dragenter", this)), ae.push(this.el), t.store && t.store.get && this.sort(t.store.get(this) || []), V(this, ai());
}
p.prototype = /** @lends Sortable.prototype */
{
  constructor: p,
  _isOutsideThisEl: function(t) {
    !this.el.contains(t) && t !== this.el && (ft = null);
  },
  _getDirection: function(t, e) {
    return typeof this.options.direction == "function" ? this.options.direction.call(this, t, e, c) : this.options.direction;
  },
  _onTapStart: function(t) {
    if (t.cancelable) {
      var e = this, n = this.el, o = this.options, r = o.preventOnFilter, a = t.type, l = t.touches && t.touches[0] || t.pointerType && t.pointerType === "touch" && t, s = (l || t).target, h = t.target.shadowRoot && (t.path && t.path[0] || t.composedPath && t.composedPath()[0]) || s, d = o.filter;
      if (wi(n), !c && !(/mousedown|pointerdown/.test(a) && t.button !== 0 || o.disabled) && !h.isContentEditable && !(!this.nativeDraggable && Nt && s && s.tagName.toUpperCase() === "SELECT") && (s = L(s, o.draggable, n, !1), !(s && s.animated) && Jt !== s)) {
        if (mt = F(s), Rt = F(s, o.draggable), typeof d == "function") {
          if (d.call(this, t, s, this)) {
            x({
              sortable: e,
              rootEl: h,
              name: "filter",
              targetEl: s,
              toEl: n,
              fromEl: n
            }), I("filter", e, {
              evt: t
            }), r && t.preventDefault();
            return;
          }
        } else if (d && (d = d.split(",").some(function(u) {
          if (u = L(h, u.trim(), n, !1), u)
            return x({
              sortable: e,
              rootEl: u,
              name: "filter",
              targetEl: s,
              fromEl: n,
              toEl: n
            }), I("filter", e, {
              evt: t
            }), !0;
        }), d)) {
          r && t.preventDefault();
          return;
        }
        o.handle && !L(h, o.handle, n, !1) || this._prepareDragStart(t, l, s);
      }
    }
  },
  _prepareDragStart: function(t, e, n) {
    var o = this, r = o.el, a = o.options, l = r.ownerDocument, s;
    if (n && !c && n.parentNode === r) {
      var h = D(n);
      if (w = r, c = n, S = c.parentNode, ht = c.nextSibling, Jt = n, Wt = a.group, p.dragged = c, st = {
        target: c,
        clientX: (e || t).clientX,
        clientY: (e || t).clientY
      }, Qe = st.clientX - h.left, tn = st.clientY - h.top, this._lastX = (e || t).clientX, this._lastY = (e || t).clientY, c.style["will-change"] = "all", s = function() {
        if (I("delayEnded", o, {
          evt: t
        }), p.eventCanceled) {
          o._onDrop();
          return;
        }
        o._disableDelayedDragEvents(), !Ve && o.nativeDraggable && (c.draggable = !0), o._triggerDragStart(t, e), x({
          sortable: o,
          name: "choose",
          originalEvent: t
        }), H(c, a.chosenClass, !0);
      }, a.ignore.split(",").forEach(function(d) {
        dn(c, d.trim(), _e);
      }), y(l, "dragover", lt), y(l, "mousemove", lt), y(l, "touchmove", lt), a.supportPointer ? (y(l, "pointerup", o._onDrop), !this.nativeDraggable && y(l, "pointercancel", o._onDrop)) : (y(l, "mouseup", o._onDrop), y(l, "touchend", o._onDrop), y(l, "touchcancel", o._onDrop)), Ve && this.nativeDraggable && (this.options.touchStartThreshold = 4, c.draggable = !0), I("delayStart", this, {
        evt: t
      }), a.delay && (!a.delayOnTouchOnly || e) && (!this.nativeDraggable || !(Lt || K))) {
        if (p.eventCanceled) {
          this._onDrop();
          return;
        }
        a.supportPointer ? (y(l, "pointerup", o._disableDelayedDrag), y(l, "pointercancel", o._disableDelayedDrag)) : (y(l, "mouseup", o._disableDelayedDrag), y(l, "touchend", o._disableDelayedDrag), y(l, "touchcancel", o._disableDelayedDrag)), y(l, "mousemove", o._delayedDragTouchMoveHandler), y(l, "touchmove", o._delayedDragTouchMoveHandler), a.supportPointer && y(l, "pointermove", o._delayedDragTouchMoveHandler), o._dragStartTimer = setTimeout(s, a.delay);
      } else
        s();
    }
  },
  _delayedDragTouchMoveHandler: function(t) {
    var e = t.touches ? t.touches[0] : t;
    Math.max(Math.abs(e.clientX - this._lastX), Math.abs(e.clientY - this._lastY)) >= Math.floor(this.options.touchStartThreshold / (this.nativeDraggable && window.devicePixelRatio || 1)) && this._disableDelayedDrag();
  },
  _disableDelayedDrag: function() {
    c && _e(c), clearTimeout(this._dragStartTimer), this._disableDelayedDragEvents();
  },
  _disableDelayedDragEvents: function() {
    var t = this.el.ownerDocument;
    b(t, "mouseup", this._disableDelayedDrag), b(t, "touchend", this._disableDelayedDrag), b(t, "touchcancel", this._disableDelayedDrag), b(t, "pointerup", this._disableDelayedDrag), b(t, "pointercancel", this._disableDelayedDrag), b(t, "mousemove", this._delayedDragTouchMoveHandler), b(t, "touchmove", this._delayedDragTouchMoveHandler), b(t, "pointermove", this._delayedDragTouchMoveHandler);
  },
  _triggerDragStart: function(t, e) {
    e = e || t.pointerType == "touch" && t, !this.nativeDraggable || e ? this.options.supportPointer ? y(document, "pointermove", this._onTouchMove) : e ? y(document, "touchmove", this._onTouchMove) : y(document, "mousemove", this._onTouchMove) : (y(c, "dragend", this), y(w, "dragstart", this._onDragStart));
    try {
      document.selection ? te(function() {
        document.selection.empty();
      }) : window.getSelection().removeAllRanges();
    } catch {
    }
  },
  _dragStarted: function(t, e) {
    if (gt = !1, w && c) {
      I("dragStarted", this, {
        evt: e
      }), this.nativeDraggable && y(document, "dragover", pi);
      var n = this.options;
      !t && H(c, n.dragClass, !1), H(c, n.ghostClass, !0), p.active = this, t && this._appendGhost(), x({
        sortable: this,
        name: "start",
        originalEvent: e
      });
    } else
      this._nulling();
  },
  _emulateDragOver: function() {
    if (j) {
      this._lastX = j.clientX, this._lastY = j.clientY, yn();
      for (var t = document.elementFromPoint(j.clientX, j.clientY), e = t; t && t.shadowRoot && (t = t.shadowRoot.elementFromPoint(j.clientX, j.clientY), t !== e); )
        e = t;
      if (c.parentNode[N]._isOutsideThisEl(t), e)
        do {
          if (e[N]) {
            var n = void 0;
            if (n = e[N]._onDragOver({
              clientX: j.clientX,
              clientY: j.clientY,
              target: t,
              rootEl: e
            }), n && !this.options.dragoverBubble)
              break;
          }
          t = e;
        } while (e = cn(e));
      En();
    }
  },
  _onTouchMove: function(t) {
    if (st) {
      var e = this.options, n = e.fallbackTolerance, o = e.fallbackOffset, r = t.touches ? t.touches[0] : t, a = m && vt(m, !0), l = m && a && a.a, s = m && a && a.d, h = qt && O && Je(O), d = (r.clientX - st.clientX + o.x) / (l || 1) + (h ? h[0] - ve[0] : 0) / (l || 1), u = (r.clientY - st.clientY + o.y) / (s || 1) + (h ? h[1] - ve[1] : 0) / (s || 1);
      if (!p.active && !gt) {
        if (n && Math.max(Math.abs(r.clientX - this._lastX), Math.abs(r.clientY - this._lastY)) < n)
          return;
        this._onDragStart(t, !0);
      }
      if (m) {
        a ? (a.e += d - (ge || 0), a.f += u - (me || 0)) : a = {
          a: 1,
          b: 0,
          c: 0,
          d: 1,
          e: d,
          f: u
        };
        var g = "matrix(".concat(a.a, ",").concat(a.b, ",").concat(a.c, ",").concat(a.d, ",").concat(a.e, ",").concat(a.f, ")");
        f(m, "webkitTransform", g), f(m, "mozTransform", g), f(m, "msTransform", g), f(m, "transform", g), ge = d, me = u, j = r;
      }
      t.cancelable && t.preventDefault();
    }
  },
  _appendGhost: function() {
    if (!m) {
      var t = this.options.fallbackOnBody ? document.body : w, e = D(c, !0, qt, !0, t), n = this.options;
      if (qt) {
        for (O = t; f(O, "position") === "static" && f(O, "transform") === "none" && O !== document; )
          O = O.parentNode;
        O !== document.body && O !== document.documentElement ? (O === document && (O = Y()), e.top += O.scrollTop, e.left += O.scrollLeft) : O = Y(), ve = Je(O);
      }
      m = c.cloneNode(!0), H(m, n.ghostClass, !1), H(m, n.fallbackClass, !0), H(m, n.dragClass, !0), f(m, "transition", ""), f(m, "transform", ""), f(m, "box-sizing", "border-box"), f(m, "margin", 0), f(m, "top", e.top), f(m, "left", e.left), f(m, "width", e.width), f(m, "height", e.height), f(m, "opacity", "0.8"), f(m, "position", qt ? "absolute" : "fixed"), f(m, "zIndex", "100000"), f(m, "pointerEvents", "none"), p.ghost = m, t.appendChild(m), f(m, "transform-origin", Qe / parseInt(m.style.width) * 100 + "% " + tn / parseInt(m.style.height) * 100 + "%");
    }
  },
  _onDragStart: function(t, e) {
    var n = this, o = t.dataTransfer, r = n.options;
    if (I("dragStart", this, {
      evt: t
    }), p.eventCanceled) {
      this._onDrop();
      return;
    }
    I("setupClone", this), p.eventCanceled || ($ = gn(c), $.removeAttribute("id"), $.draggable = !1, $.style["will-change"] = "", this._hideClone(), H($, this.options.chosenClass, !1), p.clone = $), n.cloneId = te(function() {
      I("clone", n), !p.eventCanceled && (n.options.removeCloneOnHide || w.insertBefore($, c), n._hideClone(), x({
        sortable: n,
        name: "clone"
      }));
    }), !e && H(c, r.dragClass, !0), e ? (re = !0, n._loopId = setInterval(n._emulateDragOver, 50)) : (b(document, "mouseup", n._onDrop), b(document, "touchend", n._onDrop), b(document, "touchcancel", n._onDrop), o && (o.effectAllowed = "move", r.setData && r.setData.call(n, o, c)), y(document, "drop", n), f(c, "transform", "translateZ(0)")), gt = !0, n._dragStartId = te(n._dragStarted.bind(n, e, t)), y(document, "selectstart", n), Tt = !0, window.getSelection().removeAllRanges(), Nt && f(document.body, "user-select", "none");
  },
  // Returns true - if no further action is needed (either inserted or another condition)
  _onDragOver: function(t) {
    var e = this.el, n = t.target, o, r, a, l = this.options, s = l.group, h = p.active, d = Wt === s, u = l.sort, g = T || h, _, v = this, E = !1;
    if (we) return;
    function k(Dt, $n) {
      I(Dt, v, z({
        evt: t,
        isOwner: d,
        axis: _ ? "vertical" : "horizontal",
        revert: a,
        dragRect: o,
        targetRect: r,
        canSort: u,
        fromSortable: g,
        target: n,
        completed: P,
        onMove: function(He, Sn) {
          return Vt(w, e, c, o, He, D(He), t, Sn);
        },
        changed: B
      }, $n));
    }
    function W() {
      k("dragOverAnimationCapture"), v.captureAnimationState(), v !== g && g.captureAnimationState();
    }
    function P(Dt) {
      return k("dragOverCompleted", {
        insertion: Dt
      }), Dt && (d ? h._hideClone() : h._showClone(v), v !== g && (H(c, T ? T.options.ghostClass : h.options.ghostClass, !1), H(c, l.ghostClass, !0)), T !== v && v !== p.active ? T = v : v === p.active && T && (T = null), g === v && (v._ignoreWhileAnimating = n), v.animateAll(function() {
        k("dragOverAnimationComplete"), v._ignoreWhileAnimating = null;
      }), v !== g && (g.animateAll(), g._ignoreWhileAnimating = null)), (n === c && !c.animated || n === e && !n.animated) && (ft = null), !l.dragoverBubble && !t.rootEl && n !== document && (c.parentNode[N]._isOutsideThisEl(t.target), !Dt && lt(t)), !l.dragoverBubble && t.stopPropagation && t.stopPropagation(), E = !0;
    }
    function B() {
      U = F(c), Q = F(c, l.draggable), x({
        sortable: v,
        name: "change",
        toEl: e,
        newIndex: U,
        newDraggableIndex: Q,
        originalEvent: t
      });
    }
    if (t.preventDefault !== void 0 && t.cancelable && t.preventDefault(), n = L(n, l.draggable, e, !0), k("dragOver"), p.eventCanceled) return E;
    if (c.contains(t.target) || n.animated && n.animatingX && n.animatingY || v._ignoreWhileAnimating === n)
      return P(!1);
    if (re = !1, h && !l.disabled && (d ? u || (a = S !== w) : T === this || (this.lastPutMode = Wt.checkPull(this, h, c, t)) && s.checkPut(this, h, c, t))) {
      if (_ = this._getDirection(t, n) === "vertical", o = D(c), k("dragOverValid"), p.eventCanceled) return E;
      if (a)
        return S = w, W(), this._hideClone(), k("revert"), p.eventCanceled || (ht ? w.insertBefore(c, ht) : w.appendChild(c)), P(!0);
      var M = Ne(e, l.draggable);
      if (!M || _i(t, _, this) && !M.animated) {
        if (M === c)
          return P(!1);
        if (M && e === t.target && (n = M), n && (r = D(n)), Vt(w, e, c, o, n, r, t, !!n) !== !1)
          return W(), M && M.nextSibling ? e.insertBefore(c, M.nextSibling) : e.appendChild(c), S = e, B(), P(!0);
      } else if (M && vi(t, _, this)) {
        var it = yt(e, 0, l, !0);
        if (it === c)
          return P(!1);
        if (n = it, r = D(n), Vt(w, e, c, o, n, r, t, !1) !== !1)
          return W(), e.insertBefore(c, it), S = e, B(), P(!0);
      } else if (n.parentNode === e) {
        r = D(n);
        var X = 0, ot, wt = c.parentNode !== e, R = !di(c.animated && c.toRect || o, n.animated && n.toRect || r, _), $t = _ ? "top" : "left", Z = Ze(n, "top", "top") || Ze(c, "top", "top"), St = Z ? Z.scrollTop : void 0;
        ft !== n && (ot = r[$t], Ut = !1, Gt = !R && l.invertSwap || wt), X = bi(t, n, r, _, R ? 1 : l.swapThreshold, l.invertedSwapThreshold == null ? l.swapThreshold : l.invertedSwapThreshold, Gt, ft === n);
        var G;
        if (X !== 0) {
          var rt = F(c);
          do
            rt -= X, G = S.children[rt];
          while (G && (f(G, "display") === "none" || G === m));
        }
        if (X === 0 || G === n)
          return P(!1);
        ft = n, Ht = X;
        var At = n.nextElementSibling, J = !1;
        J = X === 1;
        var zt = Vt(w, e, c, o, n, r, t, J);
        if (zt !== !1)
          return (zt === 1 || zt === -1) && (J = zt === 1), we = !0, setTimeout(mi, 30), W(), J && !At ? e.appendChild(c) : n.parentNode.insertBefore(c, J ? At : n), Z && pn(Z, 0, St - Z.scrollTop), S = c.parentNode, ot !== void 0 && !Gt && (Qt = Math.abs(ot - D(n)[$t])), B(), P(!0);
      }
      if (e.contains(c))
        return P(!1);
    }
    return !1;
  },
  _ignoreWhileAnimating: null,
  _offMoveEvents: function() {
    b(document, "mousemove", this._onTouchMove), b(document, "touchmove", this._onTouchMove), b(document, "pointermove", this._onTouchMove), b(document, "dragover", lt), b(document, "mousemove", lt), b(document, "touchmove", lt);
  },
  _offUpEvents: function() {
    var t = this.el.ownerDocument;
    b(t, "mouseup", this._onDrop), b(t, "touchend", this._onDrop), b(t, "pointerup", this._onDrop), b(t, "pointercancel", this._onDrop), b(t, "touchcancel", this._onDrop), b(document, "selectstart", this);
  },
  _onDrop: function(t) {
    var e = this.el, n = this.options;
    if (U = F(c), Q = F(c, n.draggable), I("drop", this, {
      evt: t
    }), S = c && c.parentNode, U = F(c), Q = F(c, n.draggable), p.eventCanceled) {
      this._nulling();
      return;
    }
    gt = !1, Gt = !1, Ut = !1, clearInterval(this._loopId), clearTimeout(this._dragStartTimer), $e(this.cloneId), $e(this._dragStartId), this.nativeDraggable && (b(document, "drop", this), b(e, "dragstart", this._onDragStart)), this._offMoveEvents(), this._offUpEvents(), Nt && f(document.body, "user-select", ""), f(c, "transform", ""), t && (Tt && (t.cancelable && t.preventDefault(), !n.dropBubble && t.stopPropagation()), m && m.parentNode && m.parentNode.removeChild(m), (w === S || T && T.lastPutMode !== "clone") && $ && $.parentNode && $.parentNode.removeChild($), c && (this.nativeDraggable && b(c, "dragend", this), _e(c), c.style["will-change"] = "", Tt && !gt && H(c, T ? T.options.ghostClass : this.options.ghostClass, !1), H(c, this.options.chosenClass, !1), x({
      sortable: this,
      name: "unchoose",
      toEl: S,
      newIndex: null,
      newDraggableIndex: null,
      originalEvent: t
    }), w !== S ? (U >= 0 && (x({
      rootEl: S,
      name: "add",
      toEl: S,
      fromEl: w,
      originalEvent: t
    }), x({
      sortable: this,
      name: "remove",
      toEl: S,
      originalEvent: t
    }), x({
      rootEl: S,
      name: "sort",
      toEl: S,
      fromEl: w,
      originalEvent: t
    }), x({
      sortable: this,
      name: "sort",
      toEl: S,
      originalEvent: t
    })), T && T.save()) : U !== mt && U >= 0 && (x({
      sortable: this,
      name: "update",
      toEl: S,
      originalEvent: t
    }), x({
      sortable: this,
      name: "sort",
      toEl: S,
      originalEvent: t
    })), p.active && ((U == null || U === -1) && (U = mt, Q = Rt), x({
      sortable: this,
      name: "end",
      toEl: S,
      originalEvent: t
    }), this.save()))), this._nulling();
  },
  _nulling: function() {
    I("nulling", this), w = c = S = m = ht = $ = Jt = et = st = j = Tt = U = Q = mt = Rt = ft = Ht = T = Wt = p.dragged = p.ghost = p.clone = p.active = null, se.forEach(function(t) {
      t.checked = !0;
    }), se.length = ge = me = 0;
  },
  handleEvent: function(t) {
    switch (t.type) {
      case "drop":
      case "dragend":
        this._onDrop(t);
        break;
      case "dragenter":
      case "dragover":
        c && (this._onDragOver(t), gi(t));
        break;
      case "selectstart":
        t.preventDefault();
        break;
    }
  },
  /**
   * Serializes the item into an array of string.
   * @returns {String[]}
   */
  toArray: function() {
    for (var t = [], e, n = this.el.children, o = 0, r = n.length, a = this.options; o < r; o++)
      e = n[o], L(e, a.draggable, this.el, !1) && t.push(e.getAttribute(a.dataIdAttr) || Ei(e));
    return t;
  },
  /**
   * Sorts the elements according to the array.
   * @param  {String[]}  order  order of the items
   */
  sort: function(t, e) {
    var n = {}, o = this.el;
    this.toArray().forEach(function(r, a) {
      var l = o.children[a];
      L(l, this.options.draggable, o, !1) && (n[r] = l);
    }, this), e && this.captureAnimationState(), t.forEach(function(r) {
      n[r] && (o.removeChild(n[r]), o.appendChild(n[r]));
    }), e && this.animateAll();
  },
  /**
   * Save the current sorting
   */
  save: function() {
    var t = this.options.store;
    t && t.set && t.set(this);
  },
  /**
   * For each element in the set, get the first element that matches the selector by testing the element itself and traversing up through its ancestors in the DOM tree.
   * @param   {HTMLElement}  el
   * @param   {String}       [selector]  default: `options.draggable`
   * @returns {HTMLElement|null}
   */
  closest: function(t, e) {
    return L(t, e || this.options.draggable, this.el, !1);
  },
  /**
   * Set/get option
   * @param   {string} name
   * @param   {*}      [value]
   * @returns {*}
   */
  option: function(t, e) {
    var n = this.options;
    if (e === void 0)
      return n[t];
    var o = Xt.modifyOption(this, t, e);
    typeof o < "u" ? n[t] = o : n[t] = e, t === "group" && bn(n);
  },
  /**
   * Destroy
   */
  destroy: function() {
    I("destroy", this);
    var t = this.el;
    t[N] = null, b(t, "mousedown", this._onTapStart), b(t, "touchstart", this._onTapStart), b(t, "pointerdown", this._onTapStart), this.nativeDraggable && (b(t, "dragover", this), b(t, "dragenter", this)), Array.prototype.forEach.call(t.querySelectorAll("[draggable]"), function(e) {
      e.removeAttribute("draggable");
    }), this._onDrop(), this._disableDelayedDragEvents(), ae.splice(ae.indexOf(this.el), 1), this.el = t = null;
  },
  _hideClone: function() {
    if (!et) {
      if (I("hideClone", this), p.eventCanceled) return;
      f($, "display", "none"), this.options.removeCloneOnHide && $.parentNode && $.parentNode.removeChild($), et = !0;
    }
  },
  _showClone: function(t) {
    if (t.lastPutMode !== "clone") {
      this._hideClone();
      return;
    }
    if (et) {
      if (I("showClone", this), p.eventCanceled) return;
      c.parentNode == w && !this.options.group.revertClone ? w.insertBefore($, c) : ht ? w.insertBefore($, ht) : w.appendChild($), this.options.group.revertClone && this.animate(c, $), f($, "display", ""), et = !1;
    }
  }
};
function gi(i) {
  i.dataTransfer && (i.dataTransfer.dropEffect = "move"), i.cancelable && i.preventDefault();
}
function Vt(i, t, e, n, o, r, a, l) {
  var s, h = i[N], d = h.options.onMove, u;
  return window.CustomEvent && !K && !Lt ? s = new CustomEvent("move", {
    bubbles: !0,
    cancelable: !0
  }) : (s = document.createEvent("Event"), s.initEvent("move", !0, !0)), s.to = t, s.from = i, s.dragged = e, s.draggedRect = n, s.related = o || t, s.relatedRect = r || D(t), s.willInsertAfter = l, s.originalEvent = a, i.dispatchEvent(s), d && (u = d.call(h, s, a)), u;
}
function _e(i) {
  i.draggable = !1;
}
function mi() {
  we = !1;
}
function vi(i, t, e) {
  var n = D(yt(e.el, 0, e.options, !0)), o = mn(e.el, e.options, m), r = 10;
  return t ? i.clientX < o.left - r || i.clientY < n.top && i.clientX < n.right : i.clientY < o.top - r || i.clientY < n.bottom && i.clientX < n.left;
}
function _i(i, t, e) {
  var n = D(Ne(e.el, e.options.draggable)), o = mn(e.el, e.options, m), r = 10;
  return t ? i.clientX > o.right + r || i.clientY > n.bottom && i.clientX > n.left : i.clientY > o.bottom + r || i.clientX > n.right && i.clientY > n.top;
}
function bi(i, t, e, n, o, r, a, l) {
  var s = n ? i.clientY : i.clientX, h = n ? e.height : e.width, d = n ? e.top : e.left, u = n ? e.bottom : e.right, g = !1;
  if (!a) {
    if (l && Qt < h * o) {
      if (!Ut && (Ht === 1 ? s > d + h * r / 2 : s < u - h * r / 2) && (Ut = !0), Ut)
        g = !0;
      else if (Ht === 1 ? s < d + Qt : s > u - Qt)
        return -Ht;
    } else if (s > d + h * (1 - o) / 2 && s < u - h * (1 - o) / 2)
      return yi(t);
  }
  return g = g || a, g && (s < d + h * r / 2 || s > u - h * r / 2) ? s > d + h / 2 ? 1 : -1 : 0;
}
function yi(i) {
  return F(c) < F(i) ? 1 : -1;
}
function Ei(i) {
  for (var t = i.tagName + i.className + i.src + i.href + i.textContent, e = t.length, n = 0; e--; )
    n += t.charCodeAt(e);
  return n.toString(36);
}
function wi(i) {
  se.length = 0;
  for (var t = i.getElementsByTagName("input"), e = t.length; e--; ) {
    var n = t[e];
    n.checked && se.push(n);
  }
}
function te(i) {
  return setTimeout(i, 0);
}
function $e(i) {
  return clearTimeout(i);
}
ce && y(document, "touchmove", function(i) {
  (p.active || gt) && i.cancelable && i.preventDefault();
});
p.utils = {
  on: y,
  off: b,
  css: f,
  find: dn,
  is: function(t, e) {
    return !!L(t, e, t, !1);
  },
  extend: oi,
  throttle: fn,
  closest: L,
  toggleClass: H,
  clone: gn,
  index: F,
  nextTick: te,
  cancelNextTick: $e,
  detectDirection: _n,
  getChild: yt,
  expando: N
};
p.get = function(i) {
  return i[N];
};
p.mount = function() {
  for (var i = arguments.length, t = new Array(i), e = 0; e < i; e++)
    t[e] = arguments[e];
  t[0].constructor === Array && (t = t[0]), t.forEach(function(n) {
    if (!n.prototype || !n.prototype.constructor)
      throw "Sortable: Mounted plugin must be a constructor function, not ".concat({}.toString.call(n));
    n.utils && (p.utils = z(z({}, p.utils), n.utils)), Xt.mount(n);
  });
};
p.create = function(i, t) {
  return new p(i, t);
};
p.version = ni;
var A = [], Ot, Se, Ae = !1, be, ye, le, Pt;
function $i() {
  function i() {
    this.defaults = {
      scroll: !0,
      forceAutoScrollFallback: !1,
      scrollSensitivity: 30,
      scrollSpeed: 10,
      bubbleScroll: !0
    };
    for (var t in this)
      t.charAt(0) === "_" && typeof this[t] == "function" && (this[t] = this[t].bind(this));
  }
  return i.prototype = {
    dragStarted: function(e) {
      var n = e.originalEvent;
      this.sortable.nativeDraggable ? y(document, "dragover", this._handleAutoScroll) : this.options.supportPointer ? y(document, "pointermove", this._handleFallbackAutoScroll) : n.touches ? y(document, "touchmove", this._handleFallbackAutoScroll) : y(document, "mousemove", this._handleFallbackAutoScroll);
    },
    dragOverCompleted: function(e) {
      var n = e.originalEvent;
      !this.options.dragOverBubble && !n.rootEl && this._handleAutoScroll(n);
    },
    drop: function() {
      this.sortable.nativeDraggable ? b(document, "dragover", this._handleAutoScroll) : (b(document, "pointermove", this._handleFallbackAutoScroll), b(document, "touchmove", this._handleFallbackAutoScroll), b(document, "mousemove", this._handleFallbackAutoScroll)), nn(), ee(), ri();
    },
    nulling: function() {
      le = Se = Ot = Ae = Pt = be = ye = null, A.length = 0;
    },
    _handleFallbackAutoScroll: function(e) {
      this._handleAutoScroll(e, !0);
    },
    _handleAutoScroll: function(e, n) {
      var o = this, r = (e.touches ? e.touches[0] : e).clientX, a = (e.touches ? e.touches[0] : e).clientY, l = document.elementFromPoint(r, a);
      if (le = e, n || this.options.forceAutoScrollFallback || Lt || K || Nt) {
        Ee(e, this.options, l, n);
        var s = nt(l, !0);
        Ae && (!Pt || r !== be || a !== ye) && (Pt && nn(), Pt = setInterval(function() {
          var h = nt(document.elementFromPoint(r, a), !0);
          h !== s && (s = h, ee()), Ee(e, o.options, h, n);
        }, 10), be = r, ye = a);
      } else {
        if (!this.options.bubbleScroll || nt(l, !0) === Y()) {
          ee();
          return;
        }
        Ee(e, this.options, nt(l, !1), !1);
      }
    }
  }, V(i, {
    pluginName: "scroll",
    initializeByDefault: !0
  });
}
function ee() {
  A.forEach(function(i) {
    clearInterval(i.pid);
  }), A = [];
}
function nn() {
  clearInterval(Pt);
}
var Ee = fn(function(i, t, e, n) {
  if (t.scroll) {
    var o = (i.touches ? i.touches[0] : i).clientX, r = (i.touches ? i.touches[0] : i).clientY, a = t.scrollSensitivity, l = t.scrollSpeed, s = Y(), h = !1, d;
    Se !== e && (Se = e, ee(), Ot = t.scroll, d = t.scrollFn, Ot === !0 && (Ot = nt(e, !0)));
    var u = 0, g = Ot;
    do {
      var _ = g, v = D(_), E = v.top, k = v.bottom, W = v.left, P = v.right, B = v.width, M = v.height, it = void 0, X = void 0, ot = _.scrollWidth, wt = _.scrollHeight, R = f(_), $t = _.scrollLeft, Z = _.scrollTop;
      _ === s ? (it = B < ot && (R.overflowX === "auto" || R.overflowX === "scroll" || R.overflowX === "visible"), X = M < wt && (R.overflowY === "auto" || R.overflowY === "scroll" || R.overflowY === "visible")) : (it = B < ot && (R.overflowX === "auto" || R.overflowX === "scroll"), X = M < wt && (R.overflowY === "auto" || R.overflowY === "scroll"));
      var St = it && (Math.abs(P - o) <= a && $t + B < ot) - (Math.abs(W - o) <= a && !!$t), G = X && (Math.abs(k - r) <= a && Z + M < wt) - (Math.abs(E - r) <= a && !!Z);
      if (!A[u])
        for (var rt = 0; rt <= u; rt++)
          A[rt] || (A[rt] = {});
      (A[u].vx != St || A[u].vy != G || A[u].el !== _) && (A[u].el = _, A[u].vx = St, A[u].vy = G, clearInterval(A[u].pid), (St != 0 || G != 0) && (h = !0, A[u].pid = setInterval(function() {
        n && this.layer === 0 && p.active._onTouchMove(le);
        var At = A[this.layer].vy ? A[this.layer].vy * l : 0, J = A[this.layer].vx ? A[this.layer].vx * l : 0;
        typeof d == "function" && d.call(p.dragged.parentNode[N], J, At, i, le, A[this.layer].el) !== "continue" || pn(A[this.layer].el, J, At);
      }.bind({
        layer: u
      }), 24))), u++;
    } while (t.bubbleScroll && g !== s && (g = nt(g, !1)));
    Ae = h;
  }
}, 30), wn = function(t) {
  var e = t.originalEvent, n = t.putSortable, o = t.dragEl, r = t.activeSortable, a = t.dispatchSortableEvent, l = t.hideGhostForTarget, s = t.unhideGhostForTarget;
  if (e) {
    var h = n || r;
    l();
    var d = e.changedTouches && e.changedTouches.length ? e.changedTouches[0] : e, u = document.elementFromPoint(d.clientX, d.clientY);
    s(), h && !h.el.contains(u) && (a("spill"), this.onSpill({
      dragEl: o,
      putSortable: n
    }));
  }
};
function Me() {
}
Me.prototype = {
  startIndex: null,
  dragStart: function(t) {
    var e = t.oldDraggableIndex;
    this.startIndex = e;
  },
  onSpill: function(t) {
    var e = t.dragEl, n = t.putSortable;
    this.sortable.captureAnimationState(), n && n.captureAnimationState();
    var o = yt(this.sortable.el, this.startIndex, this.options);
    o ? this.sortable.el.insertBefore(e, o) : this.sortable.el.appendChild(e), this.sortable.animateAll(), n && n.animateAll();
  },
  drop: wn
};
V(Me, {
  pluginName: "revertOnSpill"
});
function Re() {
}
Re.prototype = {
  onSpill: function(t) {
    var e = t.dragEl, n = t.putSortable, o = n || this.sortable;
    o.captureAnimationState(), e.parentNode && e.parentNode.removeChild(e), o.animateAll();
  },
  drop: wn
};
V(Re, {
  pluginName: "removeOnSpill"
});
p.mount(new $i());
p.mount(Re, Me);
var Si = Object.defineProperty, Ai = Object.getOwnPropertyDescriptor, Yt = (i, t, e, n) => {
  for (var o = n > 1 ? void 0 : n ? Ai(t, e) : t, r = i.length - 1, a; r >= 0; r--)
    (a = i[r]) && (o = (n ? a(t, e, o) : a(o)) || o);
  return n && o && Si(t, e, o), o;
};
let Et = class extends It {
  constructor() {
    super(...arguments), this.title = "", this.value = "", this.items = [];
  }
  firstUpdated() {
    new p(this.sortContainer, {
      animation: 150,
      ghostClass: "opacity-25",
      onEnd: (i) => {
        this.value = `Moved from ${i.oldIndex} to ${i.newIndex}`, this.dispatchEvent(
          new CustomEvent("change", {
            detail: `Moved from ${i.oldIndex} to ${i.newIndex}`
          })
        );
      }
    });
  }
  createRenderRoot() {
    return this;
  }
  render() {
    return console.log(this), We`
      <div class="flex p-10 border border-primary rounded">
        <div class="flex flex-col gap-8">
          <div class="text-lg">${this.title}: <span class="font-bold">${this.value}</span></div>
          <div>Open your console to see event results</div>
          <div id="sortable-container" class="flex flex-col gap-4">
            ${this.items.map(
      (i) => We` <div class="bg-primary text-primary-content p-4 rounded-box">${i.name}</div> `
    )}
          </div>
        </div>
      </div>
    `;
  }
};
Yt([
  Jn("#sortable-container")
], Et.prototype, "sortContainer", 2);
Yt([
  xe({ type: String })
], Et.prototype, "title", 2);
Yt([
  xe({ type: String })
], Et.prototype, "value", 2);
Yt([
  xe({ type: Array })
], Et.prototype, "items", 2);
Et = Yt([
  qn("sortable-example")
], Et);
export {
  Et as SortableExample
};
//# sourceMappingURL=lit.js.map
