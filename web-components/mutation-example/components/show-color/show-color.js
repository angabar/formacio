import template from "./template.js";

export default class ShowColor extends HTMLElement {
    constructor() {
        super();

        this.attachShadow({ mode: "open" });
        this.shadowRoot.innerHTML = template.render({
            data: null,
            scope: this,
        });

        this.dom = template.mapDOM(this.shadowRoot);
    }

    static get observedAttributes() {
        return ["colorbakground"];
    }

    attributeChangedCallback(name, _, newVal) {
        if (name === "colorbakground") {
            this.dom.colorTemplateComponent.style.backgroundColor = newVal;
        }
    }

    get colorBackground() {
        return this.getAttribute("colorbakground");
    }

    set colorBackground(newValue) {
        return this.setAttribute("colorbakground", newValue);
    }
}

if (!customElements.get("show-color")) {
    customElements.define("show-color", ShowColor);
}
