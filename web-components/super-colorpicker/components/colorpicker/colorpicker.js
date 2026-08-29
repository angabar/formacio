import Template from "./template.js";

export default class ColorPicker extends HTMLElement {
    constructor() {
        super();

        this.attachShadow({ mode: "open" });
        this.shadowRoot.innerHTML = Template.render();
        this.dom = Template.mapDOM(this.shadowRoot);

        // Con MutationObserver podemos estar pendientes de los cambios que
        // ocurran en todos los componentes y etiquetas HTML que estén incluidas
        // en el componente raíz que definamos, en este caso, this.shadowRoot
        const observer = new MutationObserver((event) =>
            this.onMutationChange(event),
        );
        observer.observe(this.shadowRoot, { attributes: true, subtree: true });
    }

    onMutationChange(records) {
        console.log("records__ ", records);

        records.forEach((record) => {
            this.data = Handlers.update({
                model: this.data,
                dom: this.dom,
                component: this,
                element: record.target,
                attribute: record.attributeName,
            });
        });
    }
}

if (!customElements.get("color-picker")) {
    customElements.define("color-picker", ColorPicker);
}
