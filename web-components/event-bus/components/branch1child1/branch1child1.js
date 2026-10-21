import template from "./template.js";

class Branch1Child1 extends HTMLElement {
    connectedCallback() {
        this.innerHTML = template.render();
        this.dom = template.mapDOM(this);

        this.addEventListener("set-title", (event) => {
            this.dom.titleSection.innerHTML = event?.detail?.text;
        });
    }
}

if (!customElements.get("branch1-child1")) {
    customElements.define("branch1-child1", Branch1Child1);
}
