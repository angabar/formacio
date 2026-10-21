import "../branch1child1/branch1child1.js";
import "../branch2child1/branch2child1.js";

export default {
    render(props) {
        return `
            ${this.css(props)}
            ${this.html(props)}
        `;
    },
    html(props) {
        return `
            <h1>Event calls experiment</h1>
            <branch1-child1></branch1-child1>
            <branch2-child1></branch2-child1>
        `;
    },
    css(props) {
        return `
            <style>

            </style>
        `;
    },
};
