import template from "./template.js";

class AppEvents extends HTMLElement {
    connectedCallback() {
        this.innerHTML = template.render();
    }
}

if (!customElements.get("app-events")) {
    customElements.define("app-events", AppEvents);
}
