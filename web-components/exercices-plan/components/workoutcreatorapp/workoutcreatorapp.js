import Template from "./template.js";

class WorkoutCreatorApp extends HTMLElement {
    constructor() {
        super();

        this.attachShadow({ mode: "open" });
        this.shadowRoot.innerHTML = Template.render();
    }
}

if (!customElements.get("workout-creator-app")) {
    customElements.define("workout-creator-app", WorkoutCreatorApp);
}
