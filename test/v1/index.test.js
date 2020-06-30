const request = require("supertest");

test("GET /api/v1", (done) => {
  request("http://localhost:8080")
    .get("/")
    //    .set('Accept', 'application/json')
    //    .expect('Content-Type', /json/)
    .expect(200, done);
});
