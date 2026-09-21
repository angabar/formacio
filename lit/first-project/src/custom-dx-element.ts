import { LitElement, css, html } from "lit";
import { customElement } from "lit/decorators.js";
import Tooltip from "devextreme/ui/tooltip";
import "devextreme/dist/css/dx.light.css";

@customElement("custom-dx-element")
export class MyElement extends LitElement {
    static styles = css`
        button {
            color: black;
            background: white;
            padding: 8px 16px;
            border: 10px solid red;
        }
    `;

    private tooltip?: Tooltip;

    protected createRenderRoot() {
        return this;
    }

    protected firstUpdated() {
        const target = this.querySelector("#target");
        const tooltipHost = this.querySelector("#tooltip");

        if (!target || !tooltipHost) {
            return;
        }

        this.tooltip = new Tooltip(tooltipHost as HTMLElement, {
            target,
            showEvent: "mouseenter",
            hideEvent: "mouseleave",
            contentTemplate: () => {
                const div = document.createElement("div");
                div.textContent = "Hola desde DevExtreme";
                return div;
            },
        });
    }

    disconnectedCallback() {
        super.disconnectedCallback();

        this.tooltip?.dispose();
    }

    protected render() {
        return html`
            <style>
                ${MyElement.styles}
            </style>
            <button id="target">Pasa el ratón</button>
            <div id="tooltip"></div>
        `;
    }
}
