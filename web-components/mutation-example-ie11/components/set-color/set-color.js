import template from "./template.js";

export default class SetColor extends HTMLElement {
    static get USE_SHADOWDOM_WHEN_AVAILABLE() {
        return false;
    }

    constructor() {
        super();

        if (SetColor.USE_SHADOWDOM_WHEN_AVAILABLE && this.attachShadow) {
            this.root = this.attachShadow({ mode: "open" });
        } else {
            this.root = this;
        }
    }

    connectedCallback() {
        if (!this.initialized) {
            this.root.innerHTML = template.render({
                useShadowDOM:
                    SetColor.USE_SHADOWDOM_WHEN_AVAILABLE && this.attachShadow,
            });

            this.dom = template.mapDOM(this.root);

            this.dom.rollButton.addEventListener("click", () => {
                this.hexColor = `#${Math.floor(Math.random() * 16777215)
                    .toString(16)
                    .padStart(6, "0")}`;
            });

            this.initialized = true;
        }
    }

    get hexColor() {
        return this.getAttribute("hexcolor");
    }

    set hexColor(newHexColor) {
        this.setAttribute("hexcolor", newHexColor);
    }
}

if (!customElements.get("set-color")) {
    customElements.define("set-color", SetColor);
}
