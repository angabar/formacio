const { defineConfig, devices } = require("@playwright/test");

module.exports = defineConfig({
    testDir: "./tests",
    fullyParallel: true,
    forbidOnly: !!process.env.CI,
    retries: process.env.CI ? 2 : 0,
    workers: process.env.CI ? 1 : undefined,
    reporter: "html",
    timeout: 30000,
    expect: {
        timeout: 5000,
    },
    use: {
        browserName: "chromium",
        trace: "on-first-retry",
        headless: false,
    },
    projects: [
        {
            name: "chromium",
            use: { ...devices["Desktop Chrome"] },
        },
    ],
});
