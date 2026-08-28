import template from "./template.js";

class WCIASlider extends HTMLElement {
    #isDragging = false;

    constructor() {
        super();

        this.attachShadow({ mode: "open" });

        this.shadowRoot.innerHTML = template.render();

        this.dom = template.mapDOM(this.shadowRoot);

        this.addEventListener("mousedown", (event) => this.eventHandler(event));

        document.addEventListener("mouseup", (event) =>
            this.eventHandler(event),
        );
        document.addEventListener("mousemove", (event) =>
            this.eventHandler(event),
        );
    }

    static get observedAttributes() {
        return ["value", "backgroundcolor"];
    }

    attributeChangedCallback(name, oldVal, newVal) {
        switch (name) {
            case "value":
                this.refreshSlider(newVal);
                break;

            case "backgroundcolor":
                this.setColor(newVal);
                break;
        }
    }

    get value() {
        return this.getAttribute("value");
    }

    set value(newVal) {
        this.setAttribute("value", newVal);
    }

    get backgroundcolor() {
        return this.getAttribute("backgroundcolor");
    }

    set backgroundcolor(newVal) {
        this.setAttribute("backgroundcolor", newVal);
    }

    refreshSlider(newVal) {
        if (this.dom.thumb) {
            this.dom.thumb.style.left = `${(newVal / 100) * this.offsetWidth - this.dom.thumb.offsetWidth / 2}px`;
        }
    }

    setColor(newVal) {
        if (this.dom.bgOverlay) {
            this.dom.bgOverlay.style.background = newVal;
        }
    }

    updateX(x) {
        let hPos = x - this.dom.thumb.offsetWidth / 2;

        if (hPos > this.offsetWidth) {
            hPos = this.offsetWidth;
        }

        if (hPos < 0) {
            hPos = 0;
        }

        this.value = (hPos / this.offsetWidth) * 100;
    }

    eventHandler(event) {
        const bounds = this.getBoundingClientRect();
        const x = event.clientX - bounds.left;

        switch (event.type) {
            case "mousedown":
                this.#isDragging = true;
                this.updateX(x);
                this.refreshSlider(this.value);
                break;
            case "mouseup":
                this.#isDragging = false;
                break;
            case "mousemove":
                if (this.#isDragging) {
                    this.updateX(x);
                    this.refreshSlider(this.value);
                    break;
                }
        }
    }
}

if (!customElements.get("wcia-slider")) {
    customElements.define("wcia-slider", WCIASlider);
}
