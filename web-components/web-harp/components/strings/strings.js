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

        this.stringElements = this.querySelectorAll("web-harp-string");
    }

    set points(points) {
        if (!this.stringElements) {
            return;
        }

        if (!points || !points.current) {
            return;
        }

        let magnitude = Math.abs(points.current.x - points.last.x);
        let xMin = Math.min(points.current.x, points.last.x);
        let xMax = Math.max(points.current.x, points.last.x);

        for (let i = 0; i < this.stringElements.length; i++) {
            if (
                xMin <= this.stringElements[i].offsetLeft &&
                xMax >= this.stringElements[i].offsetLeft
            ) {
                let strum = {
                    power: magnitude,
                    string: i,
                };

                this.stringElements[i].strum(strum);
            }
        }
    }
}

if (!customElements.get("web-harp-strings")) {
    customElements.define("web-harp-strings", WebHarpStrings);
}
