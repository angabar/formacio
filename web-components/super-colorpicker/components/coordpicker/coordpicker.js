import Template from "./template.js";

export default class CoordPicker extends HTMLElement {
    constructor() {
        super();

        this.attachShadow({ mode: "open" });
        this.shadowRoot.innerHTML = Template.render();
        this.dom = Template.mapDOM(this.shadowRoot);

        document.addEventListener("mousemove", (event) =>
            this.eventHandler(event),
        );
        document.addEventListener("mouseup", (event) =>
            this.eventHandler(event),
        );
        this.addEventListener("mousedown", (event) => this.eventHandler(event));
    }

    eventHandler(event) {
        const bounds = this.getBoundingClientRect();
        const coords = {
            x: event.clientX - bounds.left,
            y: event.clientY - bounds.top,
        };

        switch (event.type) {
            case "mousedown":
                this.isDragging = true;
                this.updateXY(coords.x, coords.y);
                this.refreshCoordinates();
                break;
            case "mouseup":
                this.isDragging = false;
                break;
            case "mousemove":
                if (this.isDragging) {
                    this.updateXY(coords.x, coords.y);
                    this.refreshCoordinates();
                }
                break;
        }
    }

    updateXY(x, y) {
        let hPos = x - this.dom.thumb.offsetWidth / 2;
        let vPos = y - this.dom.thumb.offsetHeight / 2;

        if (hPos > this.offsetWidth) {
            hPos = this.offsetWidth;
        }

        if (hPos < 0) {
            hPos = 0;
        }

        if (vPos > this.offsetHeight) {
            vPos = this.offsetHeight;
        }

        if (vPos < 0) {
            vPos = 0;
        }

        this.x = (hPos / this.offsetWidth) * 100;
        this.y = (vPos / this.offsetHeight) * 100;
    }

    refreshCoordinates() {
        this.dom.thumb.style.left = `${(this.x / 100) * this.offsetWidth - this.dom.thumb.offsetWidth / 2}px`;
        this.dom.thumb.style.top = `${(this.y / 100) * this.offsetHeight - this.dom.thumb.offsetHeight / 2}px`;
    }

    static get observedAttributes() {
        return ["x", "y", "backgroundColor"];
    }

    attributeChangedCallback(name, _, newVal) {
        switch (name) {
            case "x":
            case "y":
                this.refreshCoordinates();
                break;
            case "backgroundColor":
                this.style.backgroundColor = newVal;
                break;
        }
    }

    set x(newVal) {
        this.setAttribute("x", newVal);
    }

    get x() {
        return this.getAttribute("x");
    }

    set y(newVal) {
        this.setAttribute("y", newVal);
    }

    get y() {
        return this.getAttribute("y");
    }

    set backgroundColor(newVal) {
        this.setAttribute("backgroundColor", newVal);
    }

    get backgroundColor() {
        return this.getAttribute("backgroundColor");
    }
}

if (!customElements.get("coord-picker")) {
    customElements.define("coord-picker", CoordPicker);
}
