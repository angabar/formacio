import "../set-color/set-color.js";
import "../show-color/show-color.js";

export default {
    render(props) {
        return `
            ${this.css(props)}
            ${this.html(props)}
        `;
    },
    mapDOM(scope) {
        return {
            colorTemplate: scope.querySelector("show-color"),
        };
    },
    html(props) {
        return `
            <set-color></set-color>
            <show-color></show-color>
        `;
    },
    css(props) {
        return `
            <style>
            </style>
        `;
    },
};
