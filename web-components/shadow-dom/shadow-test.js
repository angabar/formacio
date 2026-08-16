class ShadowTest extends HTMLElement {
    // Cuando usamos shadowDOM podemos usar el constructor en lugar de
    // connectedCallback para inicializar el HTML y los listenners, puesto que
    // el DOM, al ser un fragment está disponible de manera inmediata, no es
    // necesario esperar a que esté anclado al DOM principal
    constructor() {
        super();

        this.attachShadow({ mode: "open" });

        this.shadowRoot.innerHTML = `<div class="test">hola shadow</div>`;
    }
}

if (!customElements.get("shadow-test")) {
    customElements.define("shadow-test", ShadowTest);
}
