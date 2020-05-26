import * as Koa from 'koa';
import * as ratelimit from 'koa-ratelimit';
import * as Boom from '@hapi/boom';
import * as Journl from 'journl';
import * as r from 'rethinkdb';

import resourcesRouter from './routes/resources';

export const logger = new Journl({});
export let db;
r.connect({ host: 'localhost', port: 28015 }, (err, conn) => {
  if (err) throw err;
  logger.success('Established database connection on port 28015');
  db = conn;
});

const app = new Koa();
const ratelimitStore = new Map();

app.use(
  ratelimit({
    driver: 'memory',
    db: ratelimitStore,
    duration: 10000,
    errorMessage: Boom.tooManyRequests('>30 requests in 10s; ratelimited'),
    id: (ctx) => ctx.ip,
    max: 30,
    whitelist: (): boolean => false,
    blacklist: (): boolean => false,
  })
);

// This is separate to the middleware declarations as it'll be used elsewhere.
const routerErrors = {
  throw: true,
  notImplemented: () => Boom.notImplemented(),
  methodNotAllowed: () => Boom.methodNotAllowed(),
};

app.use(resourcesRouter.routes());
app.use(resourcesRouter.allowedMethods(routerErrors));

app.listen(3000).on('listening', () => logger.success('Listening on port 3000'));
