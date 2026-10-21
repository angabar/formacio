import template from "./template.js";

class Branch1Child2 extends HTMLElement {
    connectedCallback() {
        this.innerHTML = template.render();
        this.dom = template.mapDOM(this);

        this.dom.sendEventButton.addEventListener("click", () => {
            this.dispatchEvent(
                new CustomEvent("set-title", {
                    detail: {
                        text: "there it is!",
                    },
                    bubbles: true,
                }),
            );
        });
    }
}

if (!customElements.get("branch1-child2")) {
    customElements.define("branch1-child2", Branch1Child2);
}
