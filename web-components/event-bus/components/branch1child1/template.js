import "../branch1child2/branch1child2.js";

export default {
    mapDOM(scope) {
        return {
            titleSection: scope.querySelector("#title-section"),
        };
    },
    render(props) {
        return `
            ${this.css(props)}
            ${this.html(props)}
        `;
    },
    html(props) {
        return `
            <div>Display result here: <span id="title-section"></span></div>
            <branch1-child2></branch1-child2>
        `;
    },
    css(props) {
        return `
            <style>
                branch1-child1 {
                    display: flex;
                    flex-direction: column;
                    border: 1px solid black;
                    border-radius: 8px;
                    padding: 8px;
                    background-color: yellow;
                }
            </style>
        `;
    },
};
