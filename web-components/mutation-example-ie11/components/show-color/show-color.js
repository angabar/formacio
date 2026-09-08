import template from "./template.js";

export default class ShowColor extends HTMLElement {
    static get USE_SHADOWDOM_WHEN_AVAILABLE() {
        return false;
    }

    constructor() {
        super();

        if (ShowColor.USE_SHADOWDOM_WHEN_AVAILABLE && this.attachShadow) {
            this.root = this.attachShadow({ mode: "open" });
        } else {
            this.root = this;
        }
    }

    connectedCallback() {
        if (!this.initialized) {
            this.root.innerHTML = template.render({
                useShadowDOM:
                    ShowColor.USE_SHADOWDOM_WHEN_AVAILABLE && this.attachShadow,
                data: null,
                scope: this,
            });

            this.dom = template.mapDOM(this.root);

            this.initialized = true;
        }
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
