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
    errorMessage: Boom.tooManyRequests(),
    id: (ctx) => ctx.ip,
    max: 10,
    whitelist: (): boolean => false,
    blacklist: (): boolean => false,
  })
);

const routerErrors = {
  throw: true,
  notImplemented: () => Boom.notImplemented(),
  methodNotAllowed: () => Boom.methodNotAllowed(),
};

app.use(resourcesRouter.routes());
app.use(resourcesRouter.allowedMethods(routerErrors));

app.listen(3000).on('listening', () => logger.success('Listening on port 3000'));
