import test, { expect, request } from "@playwright/test";

const loginPayload = {
    userEmail: "anshika@gmail.com",
    userPassword: "Iamking@000",
};

let userToken = null;

test.beforeAll(async () => {
    const apiContext = await request.newContext();

    const logInResponse = await apiContext.post(
        "https://rahulshettyacademy.com/api/ecom/auth/login",
        {
            data: loginPayload,
        },
    );

    expect(logInResponse.ok()).toBeTruthy();

    const parsedLoginResponse = await logInResponse.json();

    userToken = parsedLoginResponse.token;
});

test("Web API Part 1", async ({ page }) => {
    await page.goto("https://rahulshettyacademy.com/client");

    await page.locator(".card-body b").first().waitFor();

    const titles = await page.locator(".card-body b").allTextContents();
});
