import Template from "./template.js";
import Handlers from "./handlers.js";

export default class ColorPicker extends HTMLElement {
    static get DEFAULT_HEX() {
        return "#77aabb";
    }
    static get DEFAULT_ALPHA() {
        return 100;
    }

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

        this.shadowRoot.addEventListener("change", (event) =>
            this.onInputValueChanged(event),
        );
    }

    connectedCallback() {
        if (!this.hex) {
            this.hex = ColorPicker.DEFAULT_HEX;
        }

        if (!this.alpha) {
            this.alpha = ColorPicker.DEFAULT_ALPHA;
        }
    }

    onInputValueChanged(event) {
        this.data = Handlers.update({
            model: this.data,
            dom: this.dom,
            component: this,
            element: event.target,
        });
    }

    onMutationChange(records) {
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
