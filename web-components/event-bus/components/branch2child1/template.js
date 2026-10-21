export default {
    render(props) {
        return `
            ${this.css(props)}
            ${this.html(props)}
        `;
    },
    html(props) {
        return `
            <h1>brnach 2 child 1</h1>
        `;
    },
    css(props) {
        return `
            <style>

            </style>
        `;
    },
};
