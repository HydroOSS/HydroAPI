import * as Router from '@koa/router';
import * as r from 'rethinkdb';
import * as Boom from '@hapi/boom';
import { Context } from 'koa';
import { db } from '..';

const router = new Router({ prefix: '/api/v1/resources' });

['servers', 'users', 'strikes', 'merits', 'giveaways'].forEach((resource) => {
  router.get(`/${resource}/:key`, async (ctx: Context) => {
    const server = await r.db('Hydro').table(resource).get(ctx.params.key).run(db);
    if (!server) {
      ctx.body = Boom.notFound(`There isn't a document in the ${resource} table for the specified ID.`);
      ctx.status = 404;
      return;
    }
    ctx.body = { data: server };
  });
});

export default router;
