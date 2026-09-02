import template from "./template.js";

export default class ShowColor extends HTMLElement {
    constructor() {
        super();

        this.attachShadow({ mode: "open" });
        this.shadowRoot.innerHTML = template.render({
            data: null,
            scope: this,
        });
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
