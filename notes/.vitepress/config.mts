import { defineConfig } from "vitepress";

export default defineConfig({
    title: "Notes",
    description: "A VitePress Site",
    themeConfig: {
        nav: [
            { text: "Home", link: "/" },
            { text: "Lit", link: "/lit-context" },
        ],
        sidebar: [
            {
                text: "Lit",
                items: [{ text: "Context", link: "/lit-context" }],
            },
        ],
        socialLinks: [
            { icon: "github", link: "https://github.com/vuejs/vitepress" },
        ],
    },
});
