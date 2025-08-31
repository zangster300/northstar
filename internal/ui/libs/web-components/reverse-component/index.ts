class ReverseComponent extends HTMLElement {
  static get observedAttributes() {
    return ["name"];
  }

  attributeChangedCallback(name: string, oldValue: string, newValue: string) {
    const len = newValue.length;

    let value: string | any[] = Array(len);

    let i = len - 1;

    for (const char of newValue) {
      value[i--] = char.codePointAt(0);
    }

    value = String.fromCodePoint(...value);
    this.dispatchEvent(new CustomEvent("reverse", { detail: { value } }));
  }
}

customElements.define("reverse-component", ReverseComponent);
