export default {
    mapDOM(scope) {
        return {
            title: scope.querySelector("h1"),
            subtitle: scope.querySelector("h3"),
        };
    },
    render(props) {
        return `${this.html(props)}
                ${this.css(props)}`;
    },
    html(props) {
        return `
            <h1>${props.title}</h1>
            <h3>${props.subtitle}</h3>
        `;
    },
    css(props) {
        return ``;
    },
};
