import template from "./template.js";

export default class ColorRandom extends HTMLElement {
    constructor() {
        super();

        this.attachShadow({ mode: "open" });
        this.shadowRoot.innerHTML = template.render();

        this.dom = template.mapDOM(this.shadowRoot);

        this.observer = new MutationObserver((event) =>
            this.onMutationChange(event),
        );
        this.observer.observe(this.shadowRoot, {
            attributes: true,
            subtree: true,
        });
    }

    onMutationChange(events) {
        events.forEach((event) => {
            if (event.attributeName === "hexcolor") {
                const updatedColor = event.target.getAttribute("hexcolor");

                this.dom.colorTemplate.colorBackground = updatedColor;
            }
        });
    }
}

if (!customElements.get("color-random")) {
    customElements.define("color-random", ColorRandom);
}
