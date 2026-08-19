export default {
    render(props) {
        return `
            ${this.css(props)}
            ${this.html(props)}
        `;
    },
    html(props) {
        return `
            <div id="info">
                <span id="label">${props.label}</span>
                <span id="delete">x</span>
            </div>
        `;
    },
    css(props) {
        return `
            <style>
                :host {
                    display: inline-block;
                    background: radial-gradient(circle, rgba(235, 235, 235, 1) 0%, rgba(208, 208, 208, 1) 100%);
                    border-left-style: solid;
                    border-left-width: 5px;
                    width: 200px;
                    height: 200px;
                }

                :host(.cardio) {
                    border-left-color: #28a7ff;
                }

                :host(.strength) {
                    border-left-color: #75af01;
                }

                #info {
                    font-size: small;
                    display: flex;
                    align-items: center;
                    background-color: black;
                    color: white;
                }

                :host #info {
                    padding: 5px;
                }
            </style>
        `;
    },
};
