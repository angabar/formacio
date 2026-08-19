import Template from "./template.js";

export default class WorkoutExercise extends HTMLElement {
    constructor() {
        super();

        this.attachShadow({ mode: "open" });

        const params = {
            label: this.getAttribute("label"),
            type: this.getAttribute("type"),
            thumb: this.getAttribute("thumb"),
            time: this.getAttribute("time"),
            count: this.getAttribute("count"),
            estimatedTimePerCount: this.getAttribute("estimatedtimepercount"),
            sets: this.getAttribute("sets"),
        };

        this.shadowRoot.innerHTML = Template.render(params);
    }

    get label() {
        return this.getAttribute("label");
    }

    set label(newVal) {
        this.setAttribute("label", newVal);
    }

    get type() {
        return this.getAttribute("type");
    }

    set type(newVal) {
        this.setAttribute("type", newVal);
    }

    get thumb() {
        return this.getAttribute("thumb");
    }

    set thumb(newVal) {
        this.setAttribute("thumb", newVal);
    }

    get time() {
        return this.getAttribute("time");
    }

    set time(newVal) {
        this.setAttribute("time", newVal);
    }

    get count() {
        return this.getAttribute("count");
    }

    set count(newVal) {
        this.setAttribute("count", newVal);
    }

    get estimatedTimePerCount() {
        return this.getAttribute("estimatedtimepercount");
    }

    set estimatedTimePerCount(newVal) {
        this.setAttribute("estimatedtimepercount", newVal);
    }

    get sets() {
        return this.getAttribute("sets");
    }

    set sets(newVal) {
        this.setAttribute("sets", newVal);
    }

    serialize() {
        return {
            label: this.label,
            type: this.type,
            thumb: this.thumb,
            time: this.time,
            count: this.count,
            estimatedTimePerCount: this.estimatedTimePerCount,
            sets: this.sets,
        };
    }

    static toAttributeString(object) {
        let attributes = "";

        Object.keys(object).forEach((key) => {
            if (object[key]) {
                attributes += ` ${key}="${object[key]}"`;
            }
        });

        return attributes;
    }
}

if (!customElements.get("workout-exercise")) {
    customElements.define("workout-exercise", WorkoutExercise);
}
