const { test, expect } = require("@playwright/test");

test("Fist playwright test", async ({ page }) => {
    // page nos proporciona por defecto estos dos pasos

    // const context = await browser.newContext();
    // const page = await context.newPage();

    await page.goto("https://google.com");

    const title = await page.title();
    await expect(page).toHaveTitle(title);
});
