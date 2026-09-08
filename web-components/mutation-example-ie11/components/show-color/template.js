export default {
    render(props) {
        return `
            ${this.css(props)}
            ${this.html(props)}
        `;
    },
    mapDOM(scope) {
        return {
            colorTemplateComponent: scope.getElementById("color-template"),
        };
    },
    html(props) {
        return `
            <div id="color-template"></div>
        `;
    },
    css(props) {
        return `
            <style>
                :host {
                    display: flex;
                    height: 200px;
                    width: 200px;
                    border: 1px solid #383333d3;
                    border-radius: 8px;
                    margin-top: 20px;
                }

                #color-template {
                    width: 100%;
                    height: 100%;
                }
            </style>
        `;
    },
};
