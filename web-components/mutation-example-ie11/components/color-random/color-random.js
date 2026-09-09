import template from "./template.js";

export default class ColorRandom extends HTMLElement {
    static get USE_SHADOWDOM_WHEN_AVAILABLE() {
        return false;
    }

    constructor() {
        super();

        if (ColorRandom.USE_SHADOWDOM_WHEN_AVAILABLE && this.attachShadow) {
            this.root = this.attachShadow({ mode: "open" });
        } else {
            this.root = this;
        }
    }

    connectedCallback() {
        if (!this.initialized) {
            this.root.innerHTML = template.render({
                useShadowDOM:
                    ColorRandom.USE_SHADOWDOM_WHEN_AVAILABLE &&
                    this.attachShadow,
            });

            this.dom = template.mapDOM(this.root);

            this.observer = new MutationObserver((event) =>
                this.onMutationChange(event),
            );
            this.observer.observe(this.root, {
                attributes: true,
                subtree: true,
            });

            this.initialized = true;
        }
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
