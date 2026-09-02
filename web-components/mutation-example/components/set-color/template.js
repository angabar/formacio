export default {
    render(props) {
        return `
            ${this.css(props)}
            ${this.html(props)}
        `;
    },
    mapDOM(dom) {
        return {
            rollButton: dom.getElementById("roll-button"),
        };
    },
    html(props) {
        return `
            <button id="roll-button">Click me!</button>
        `;
    },
    css(props) {
        return `
            <style>
            </style>
        `;
    },
};
