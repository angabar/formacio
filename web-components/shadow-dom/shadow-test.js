class ShadowTest extends HTMLElement {
    connectedCallback() {
        this.attachShadow({ mode: "open" });

        this.shadowRoot.innerHTML = `<div class="test">hola shadow</div>`;
    }
}

if (!customElements.get("shadow-test")) {
    customElements.define("shadow-test", ShadowTest);
}
