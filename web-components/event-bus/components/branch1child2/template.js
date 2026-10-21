export default {
    mapDOM(scope) {
        return {
            sendEventButton: scope.querySelector("#send-event-button"),
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
            <button id="send-event-button">Sent event to parent</button>
        `;
    },
    css(props) {
        return `
            <style>
                branch1-child2 {
                    display: flex;
                    border: 1px solid black;
                    border-radius: 8px;
                    padding: 8px;
                    margin-top: 20px;
                    background-color: aliceblue;
                }
            </style>
        `;
    },
};
