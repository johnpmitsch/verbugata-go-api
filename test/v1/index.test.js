const request = require("supertest");

test("GET /api/v1", async () => {
  // can add headers and body
  const { status } = await request("http://localhost:8080").get("/");
  expect(status).toBe(200);
});
