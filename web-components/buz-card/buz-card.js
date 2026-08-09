import template from "./template.js";

class BuzCard extends HTMLElement {
    connectedCallback() {
        this.innerHTML = template.render({
            title: this.getAttribute("title"),
            subtitle: this.getAttribute("subtitle"),
        });

        // Definimos el DOM una única vez y ya podemos usarlo las veces que
        // queramos sabiendo que se ha declarado el querySelector una sola vez
        this.dom = template.mapDOM(this);

        this.dom.title.style.backgroundColor = "red";
    }
}

if (!customElements.get("buz-card")) {
    customElements.define("buz-card", BuzCard);
}
