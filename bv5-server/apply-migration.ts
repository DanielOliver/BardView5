import { MikroORM } from '@mikro-orm/core';
import * as dotenv from 'dotenv';

(async () => {
  dotenv.config({ path: __dirname + '/.env' });
  const orm = await MikroORM.init({
    dbName: 'bardview5',
    type: 'postgresql',
    clientUrl: process.env.DATABASE_URL,
    migrations: {
      path: './migrations',
    },
    entitiesTs: ['./src/entities/*.ts'],
    entities: ['src/entities/*.ts'],
  });

  const migrator = orm.getMigrator();
  await migrator.up();
  await orm.close(true);
})();
