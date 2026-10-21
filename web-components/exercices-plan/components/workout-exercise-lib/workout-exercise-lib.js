import Template from "./template.js";
import exercices from "../../data/exercices.js";

class WorkoutExerciseLib extends HTMLElement {
    constructor() {
        super();

        this.attachShadow({ mode: "open" });
        this.shadowRoot.innerHTML = Template.render({
            exercises: exercices.all,
        });
    }
}

if (!customElements.get("workout-exercise-lib")) {
    customElements.define("workout-exercise-lib", WorkoutExerciseLib);
}
