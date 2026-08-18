import Template from "./template.js";

class WorkoutExerciseLib extends HTMLElement {
    constructor() {
        super();

        this.attachShadow({ mode: "open" });
        this.shadowRoot.innerHTML = Template.render();
    }
}

if (!customElements.get("workout-exercise-lib")) {
    customElements.define("workout-exercise-lib", WorkoutExerciseLib);
}
