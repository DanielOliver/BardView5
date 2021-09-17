import { MikroORM } from '@mikro-orm/core';
import * as dotenv from 'dotenv';

(async () => {
  dotenv.config({ path: __dirname + '/.env' });
  const orm = await MikroORM.init({
    discovery: {
      // we need to disable validation for no entities
      warnWhenNoEntities: false,
    },
    entitiesTs: ['./src/entities'],
    dbName: 'bardview5',
    type: 'postgresql',
    clientUrl: process.env.DATABASE_URL,
    migrations: {
      path: './migrations',
    },
  });

  const migrator = orm.getMigrator();
  // await migrator.createInitialMigration();
  // await migrator.createMigration('./migrations'); // creates file Migration20191019195930.ts
  await migrator.createMigration();

  await orm.close(true);
})();
