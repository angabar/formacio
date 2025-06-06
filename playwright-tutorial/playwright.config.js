// @ts-check
import { defineConfig } from "@playwright/test";

export default defineConfig({
    testDir: "./tests",
    timeout: 30_000,
    expect: {
        timeout: 5000,
    },
    reporter: "html",
    use: {
        browserName: "chromium",
        screenshot: "only-on-failure",
        trace: "retain-on-failure",
    },
});
