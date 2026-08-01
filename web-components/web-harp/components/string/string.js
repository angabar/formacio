export default class WebHarpString extends HTMLElement {
    strum(params) {
        console.log("params", params);
    }

    stopStrum() {}

    connectedCallback() {
        this.innerHTML = `
            <div class="line"></div>
            <style>
                web-harp-string > .line {
                    background-color: white;
                    height: 100%;
                    width: 2px;
                }
            </style>
        `;
    }
}

if (!customElements.get("web-harp-string")) {
    customElements.define("web-harp-string", WebHarpString);
}
