const request = require("supertest");

const server = "http://localhost:8080";

test("GET /api/v1", async () => {
  // can add headers and body
  const { status } = await request(server).get("/");
  expect(status).toBe(200);
});

test("404 when route does not exist", async () => {
  // can add headers and body
  const { status } = await request(server).get("/doesnotexist");
  expect(status).toBe(404);
});
