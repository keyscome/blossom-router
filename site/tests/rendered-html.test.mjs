import assert from "node:assert/strict";
import test from "node:test";

const root = new URL("../", import.meta.url);

async function render(path = "/") {
  const workerUrl = new URL("../dist/server/index.js", import.meta.url);
  workerUrl.searchParams.set("test", `${process.pid}-${Date.now()}-${path}`);
  const { default: worker } = await import(workerUrl.href);
  return worker.fetch(new Request(`http://localhost${path}`, { headers: { accept: "text/html" } }), {
    ASSETS: { fetch: async () => new Response("Not found", { status: 404 }) },
  }, { waitUntil() {}, passThroughOnException() {} });
}

test("renders the English product site", async () => {
  const response = await render("/");
  assert.equal(response.status, 200);
  const html = await response.text();
  assert.match(html, /Blossom Router/);
  assert.match(html, /One prompt/);
  assert.match(html, /Is Blossom Router useful for you/);
  assert.match(html, /Frequently asked questions/);
  assert.doesNotMatch(html, /codex-preview|react-loading-skeleton/i);
});

test("renders the Simplified Chinese product site", async () => {
  const response = await render("/zh");
  assert.equal(response.status, 200);
  const html = await response.text();
  assert.match(html, /一个提示词/);
  assert.match(html, /Blossom Router 适合你吗/);
  assert.match(html, /常见问题/);
});

test("ships the bespoke social card", async () => {
  const { access } = await import("node:fs/promises");
  await access(new URL("public/og.png", root));
});
