import Template from "./template.js";

class WorkoutPlan extends HTMLElement {
    constructor() {
        super();

        this.attachShadow({ mode: "open" });
        this.shadowRoot.innerHTML = Template.render();
    }
}

if (!customElements.get("workout-plan")) {
    customElements.define("workout-plan", WorkoutPlan);
}
