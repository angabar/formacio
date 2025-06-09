import test, { expect } from "@playwright/test";

test("More validations", async ({ page }) => {
    await page.goto("https://rahulshettyacademy.com/AutomationPractice/");

    await expect(page.locator("#displayed-text")).toBeVisible();
    await page.locator("#hide-textbox").click();
    await expect(page.locator("#displayed-text")).toBeHidden();

    page.on("dialog", async (dialog) => {
        await dialog.accept();
    });

    await page.locator("#confirmbtn").click();
    await page.locator("#mousehover").hover();

    const iframePage = page.frameLocator("#courses-iframe");

    await iframePage.locator("li a[href*='lifetime-access']:visible").click();
    const textCheck = await iframePage.locator(".text h2").textContent();
    console.log(textCheck.split(" ")[1]);
});
