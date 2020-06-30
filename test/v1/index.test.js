const request = require('supertest');

test('GET /api/v1', (done) => {
  request("http://localhost:3000")
    .get('/api/v1')
    .set('Accept', 'application/json')
    .expect('Content-Type', /json/)
    .expect(200, done);
});
