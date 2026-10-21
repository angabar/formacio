import template from "./template.js";

class Branch2Child1 extends HTMLElement {
    connectedCallback() {
        this.innerHTML = template.render();

        this.addEventListener("set-title", (event) => {
            console.log("event__ ", event);
        });
    }
}

if (!customElements.get("branch2-child1")) {
    customElements.define("branch2-child1", Branch2Child1);
}
