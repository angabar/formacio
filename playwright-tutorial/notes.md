## Básico

Los tests en playwright no son ejecutados de manera secuencial, se lanzan de manera simultánea, sin ningún tipo de orden por ello es importante usar `await` en cada acción.

Cada invocación de `test` recibe un objeto útil para nuestros tests, uno de los atributos es `browser`

```js
test("First playwright test", async ({ browser }) => {
    ...
});
```

El tipo de navegador que tendremos en `browser` serà el que hayamos especificado previamente en el archivo `playwright.config.js` y nos sirve para definir instancias especificas del navegador, como cookies, gateways... si lo que queremos es solo usa páginas, podemos simplemente usar ese atributo del objeto global directamente.

```js
test("First playwright test", async ({ page }) => {
    await page.goto("htps://rahulshettyacademy.com/loginpagePractise/");
});
```

## Flags

`--headed` - Lanza los tests mostrando el navegador, es muy rápido.
`--ui` - Lanza los tests en modo visual, puedes ir viendo que pasa en cada fase, como en Cypress.
`--debug` - Sirve para lanzar los tests en modo _debug_ y así poder ir viendo que pasa en cada paso.
`-g "Nombre del test"` - Sirve para lanzar solo un test en concreto.

## Configuración

Toda la configuración se incluye en el archivo mencionado antes `playwright.config.js` en el veremos atributos como:

`testDir` - Define donde se ubican los tests que han de lanzarse.
`timeout` - El tiempo en milisegundos que esperaran los tests para buscar componentes, este afecta a todos los tests.
`expect.timeout` - El tiempo en milisegundos que esperan los tests para hacer aserciones.
`reporter` - Define como se va a hacer el reporte de los resultados de los tests.
`use.browserName` - Define el navegador que vamos a utilizar.
`use.headless` - Especifica si los tests han de lanzarse visualizando el navegador con `false` se verán los navegadores con `true` (el valor por defecto) no.
`use.screenshot` - Determina si han de generarse capturas de pantalla en cada test pasandole el valor "on" o si han de guardarse capturas de pantalla solo en los casos en los que el test falle pasandole "only-on-failure".
`use.trace` - Determina si han de generarse _traces_ en cada test pasandole el valor "on" o si han de guardarse _traces_ solo en los casos en los que el test falle pasandole "retain-on-failure".

## Lanzar los tests

Por defecto Playwright lanza los tests sin mostrar el navegador, si queremos que los lance mostrando el navegador tenemos que usar el flag `--headed`

Para lanzar los tests tenemos que ejecutar `npx playwright test`

Los archivos de tests son lanzados de manera paralela, pero los tests en el mismo archivo son lanzados de manera secuencial.

Para ejecutar un único test dentro de un archivo usaremos el atributo `.only` después de `test`

```js
test.only("First playwright test", async ({ page }) => {
    await page.goto("https://rahulshettyacademy.com/loginpagePractise/");
});
```

## Aserciones

Las aserciones se hacen utilizando el objeto `expect` que importamos desde Playwright directamente, como en todas las acciones de este, tendremos que usar `await` para asegurarnos que accedemos al elemento en el momento correcto.

```js
test("Page test", async ({ page }) => {
    await page.goto("https://www.google.com/");

    await expect(page).toHaveTitle("Google");
});
```

## Selección

Para seleccionar elementos y poder hacer los tests se utiliza el método `locator` Playwright utiliza selectores de CSS para hacer las selecciones.

```js
test("Can fill fields", async ({ page }) => {
    await page.goto("https://rahulshettyacademy.com/loginpagePractise/");

    const element = page.locator("#username");
});
```

Una vez seleccionado podemos usar `fill` para escribir o `click` para clickar, hay muchos más, referir a la documentación oficial para más información.

Playwright espera automáticamente a que los elementos que hayamos especificado aparezcan en pantalla sin necesidad de definirlo como tal (ejemplo de `waitUntil` en otras herrameintas).

Como se espera Playwright depende del método de acceso que estemos utilizando, por ejemplo `textContent` esperara a que el elemento esté conectado en el HTML, pero `allTextContents` no, este último se lanza sin esperar a que los elementos estén en el HTML, por lo que sino esperamos antes, el test puede fallar.

Una estrategia para poder acceder a los elementos que todavía no han cargado, es esperar a que las peticiones http hayan terminado para ello usaremos el método `waitFor` actuando sobre un único elemento, la idea es decirle a Playwright, espera a que el primer elemento de lo selectores que te he especificado a que esté disponible en la pantalla.

```js
test("Page test", async ({ page }) => {
    // Espera que al menos haya un elemento seleccionable en la pantalla
    await page.locator(".card-body b").first().waitFor();

    // Podemos continuar con los tests
    const titles = await page.locator(".card-body b").allTextContents();

    console.log(titles);
});
```

Los checkbox no tienen aserción para comprobar si han sido des-seleccionados, pero podemos usar la propiedad `isChecked` del `locator` para dereminar si esta es `true` o `false` en función de si está marcado o no.

```js
expect(await page.locator("#terms").isChecked()).toBeFalsy();
```

Observa bien la posición del `await` es importante observar antes de terminar los tests, cuales son los métodos asíncronos y cuales no.

### Aserción de atributos

Si queremos comprobar si un elemento tiene un atributo con un determinado valor, usaremos el método `toHaveAttribute` el cual recibe dos valores, el nombre del atributo y el valor que debería tener.

```js
await expect(documentLink).toHaveAttribute("class", "blinkingText");
```

## Trabajando en otras ventanas

Para trabajar con otras ventanas tenemos disponible el método `waitForEvent("page")` un método que se queda escuchando hasta que alguien lance el evento de creación de página, puede ser un click o cualquier otra cosa, pero hay algo que debemos tener en cuenta.

Hay que tener en cuenta el punto de que al ser un _listenner_ ha de colocarse antes del evento que ejecuta esta acción.

Pero esto nos lleva a otro problema, cuando invocamos `waitForEvent` no podemos ponerle un `await` porque al ser un _listenner_ este se queda escuchando indefinidamente, lo que puede bloquear la ejecución normal del código, lo que tenemos que hacer es usar `Promise.all` quitando los `await`

```js
const [newPage] = await Promise.all([
    context.waitForEvent("page"),
    documentLink.click(),
]);

const text = await newPage.locator(".red").textContent();
```

Observa que necesitamos el objeto `context` del navegador, el cual nos devolverá el objeto de la página abierta para que podamos realizar los tests o comprobaciones en esta.

## Generación de código

Playwright tiene la opción de generar el código del test a medida que vayamos usando la aplicación, para ello lanzamos el comando de ejecución de tests con el _flag_ `codegen` seguido de la URL que queremos instanciar en primer lugar.

```
npx playwright codegen http://www.google.es
```

A partir de ese momento cada acción que realicemos en la web que se nos abra quedará registrada en el código como acciones o aserciones de Playwright.

Es importante tener en cuenta que esto es una herramienta avanzada y que ha de usarse solo cuando se tengan conocimientos avanzados del entorno de testing de Playwright.

## Scope del evento `locator`

El método `locator` actua sobre el _scope_ que tenga la variable que lo invoca, si es `page` entonces tendrá el _scope_ de toda la página, si es un `locator` previo, entonces el _scope_ será el de ese `locator` previo.

```js
// scope = página entera
const products = page.locator(".card-body");

// scope = producto concreto
await products.nth(i).locator("text= Add To Cart").click();
```

Observa la peculiaridad en el ejemplo anterior de como podemos utilizar el `locator` también para seleccionar textos, este ha de ponerse tal y como aparece en el DOM, con espacios o sin ellos.

## Comprobar que el elemento es visible

Playwright tiene un método para determinar si el elemento es visible en pantalla o no `isVisible` pero hay que tener en cuenta que no espera a los elementos a buscar de manera nativa, si estos tardan en aparecer, tendremos que especificar un `waitFor` de otros elementos antes.

```js
await page.locator("[routerLink*='cart']").click();

await page.locator("div li").waitFor();
const isVisible = await page.locator("h3:has-text('ZARA COAT 3')").isVisible();
```

## `getByLabel`

El selector `getByLabel` selecciona los elemento usando el `label` del DOM, a este le tenemos que pasar el texto incluido entre las etiquetas para que haga la selección.

```js
await page.getByLabel(/icecreams/i).check();
await page.getByLabel("Employed").check();
await page.getByLabel("Gender").selectOption("Female");
```

Como vemos en el ejemplo anterior `check` funciona para seleccionar elementos de `checkbox` y `radiobuttons` hace lo mismo que `click`

Para seleccionar elementos de una etiqueta `select` utilizaremos `selectOption`

Podemos también seleccionar por texto usando `getByText`

Así como usando el role especifico de un elemento HTML con `getByRole`

```js
await page.getByPlaceholder("Password").fill("abc123");
await page.getByRole("button", { name: "Submit" }).click();
await page.getByText(/submitted successfully/i).isVisible();
await page.getByRole("link", { name: "Shop" }).click();
```

Otra propiedad muy potente de Playwright es la de filtrar grupos encontrados con el método `filter` imaginemos que tenemos un `page.locator` que devuelve 4 elementos, pues bien, con `filter` podemos filtrar estos 4 elementos y obtener el que cumpla los criterios que especifiquemos para poder así seguir aplicando selectores ya dentro del resultado obtenido, por ejemplo:

```js
await page
    .locator("app-card")
    .filter({
        hasText: "Nokia Edge", // Filtramos por texto
    })
    .getByRole("button") // Aplicamos el getByRole solo sobre los resultados obtenidos del filtro
    .click();
```
