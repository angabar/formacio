export default {
    render(props) {
        return `
            ${this.css(props)}
            ${this.html(props)}
        `;
    },
    html(props) {
        return `
            <h1>My plan</h1>
            <div id="container"></div>
            <div id="time">Total time:</div>
        `;
    },
    css(props) {
        return `
            <style>
                :host {
                    display: flex;
                    flex-direction: column;
                }

                #time {
                    height: 30px
                }

                #container {
                    background: linear-gradient(90deg, rgba(235, 235, 235, 1) 0%, rgba(208, 208, 208, 1) 100%);
                    height: calc(100% - 60px);
                    overflow-y: scroll;
                }
            </style>
        `;
    },
};
