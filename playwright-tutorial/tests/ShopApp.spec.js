import test, { expect } from "@playwright/test";

test("Shop app", async ({ page }) => {
    const products = page.locator(".card-body");
    const productName = "ZARA COAT 3";

    await page.goto("https://rahulshettyacademy.com/client");

    await page.locator("#userEmail").fill("anshika@gmail.com");
    await page.locator("#userPassword").fill("Iamking@000");
    await page.locator("[value='Login']").click();

    const count = await products.count();

    for (let i = 0; i < count; i++) {
        const currentTextContentCard = await products
            .nth(i)
            .locator("b")
            .textContent();

        if (currentTextContentCard === productName) {
            await products.nth(i).locator("text= Add To Cart").click();
            await page.pause();
            break;
        }
    }

    await page.locator("[routerLink*='cart']").click();

    await page.locator("div li").first().waitFor();
    const isVisible = await page
        .locator("h3:has-text('ZARA COAT 3')")
        .isVisible();

    expect(isVisible).toBeTruthy();
});
