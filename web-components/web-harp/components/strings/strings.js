import WebHarpString from "../string/string.js";

export default class WebHarpStrings extends HTMLElement {
    connectedCallback() {
        let strings = `<div class="spacer"></div>`;

        for (let i = 0; i < this.getAttribute("strings").length; i++) {
            strings += `<web-harp-string></web-harp-string>`;
        }

        strings += `
            <style>
                web-harp-strings {
                    height: 100%;
                    display: flex;
                }

                web-harp-strings > web-harp-string, div.spacer {
                    flex: 1;
                }
            </style>
        `;

        this.innerHTML = strings;

        this.stringsElements = this.querySelectorAll("web-harp-string");
    }
}

if (!customElements.get("web-harp-strings")) {
    customElements.define("web-harp-strings", WebHarpStrings);
}
