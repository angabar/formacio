import WebHarpStrings from "../strings/strings.js";

export default class WebHarpApp extends HTMLElement {
    connectedCallback() {
        this.innerHTML = `<web-harp-strings strings="${this.getAttribute("strings")}"></web-harp-strings>`;

        this.stringsElements = this.querySelector("web-harp-strings");

        this.addEventListener("mousemove", (event) => this.onMouseMove(event));
    }

    onMouseMove(event) {
        this.stringsElements.points = {
            last: this.lastPoint,
            current: { x: event.pageX, y: event.pageY },
        };
        this.lastPoint = { x: event.pageX, y: event.pageY };
    }
}

if (!customElements.get("web-harp-app")) {
    customElements.define("web-harp-app", WebHarpApp);
}
