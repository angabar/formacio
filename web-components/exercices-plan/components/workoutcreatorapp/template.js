import "../workout-plan/workout-plan.js";
import "../workout-exercise-lib/workout-exercise-lib.js";

export default {
    render(props) {
        return `
            ${this.css(props)}
            ${this.html(props)}
        `;
    },
    html(props) {
        return `
            <workout-exercise-lib></workout-exercise-lib>
            <div id="divider-line"></div>
            <workout-plan></workout-plan>
        `;
    },
    css(props) {
        return `
            <style>
                :host {
                    display: flex;
                }

                workout-exercise-lib,
                workout-plan {
                    flex: 1;
                    height: 100%;
                    background-color: #eaeaea;
                }

                #divider-line {
                    width: 1px;
                    height: 100%;
                    margin-right: 25px;
                    background-color: black;
                }
            </style>
        `;
    },
};
