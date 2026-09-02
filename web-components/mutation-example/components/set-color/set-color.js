import template from "./template.js";

export default class SetColor extends HTMLElement {
    constructor() {
        super();

        this.attachShadow({ mode: "open" });
        this.shadowRoot.innerHTML = template.render();

        this.dom = template.mapDOM(this.shadowRoot);

        this.dom.rollButton.addEventListener("click", () => {
            this.hexColor = `#${Math.floor(Math.random() * 16777215)
                .toString(16)
                .padStart(6, "0")}`;
        });
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
