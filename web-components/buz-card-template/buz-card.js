import template from "./template.js";

class BuzCard extends HTMLElement {
    connectedCallback() {
        this.innerHTML = template.render({
            title: this.getAttribute("title"),
            subtitle: this.getAttribute("subtitle"),
        });

        this.dom = template.mapDOM(this);

        this.dom.title.style.backgroundColor = "red";
    }

    static get observedAttributes() {
        return ["layout"];
    }

    attributeChangedCallback(name, oldVal, newVal) {
        this.innerHTML = "";

        const template = document.getElementById(newVal);

        if (template) {
            const clone = template.content.cloneNode(true);
            this.appendChild(clone);
        }
    }
}

if (!customElements.get("buz-card")) {
    customElements.define("buz-card", BuzCard);
}
