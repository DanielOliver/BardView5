import { MikroORM } from '@mikro-orm/core';
import * as dotenv from 'dotenv';

(async () => {
  dotenv.config({ path: __dirname + '/.env' });
  const orm = await MikroORM.init({
    discovery: {
      // we need to disable validation for no entities
      warnWhenNoEntities: false,
    },
    dbName: 'bardview5',
    type: 'postgresql',
    clientUrl: process.env.DATABASE_URL,
  });
  const generator = orm.getEntityGenerator();
  const dump = await generator.generate({
    save: true,
    baseDir: process.cwd() + '/src/entities',
  });
  console.log(dump);
  await orm.close(true);
})();
