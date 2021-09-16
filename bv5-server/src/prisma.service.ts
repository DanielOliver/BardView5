import { INestApplication, Injectable } from '@nestjs/common';
import { PrismaClient } from '@prisma/client';

// https://docs.nestjs.com/recipes/prisma
@Injectable()
export class PrismaService extends PrismaClient {
  async enableShutdownHooks(app: INestApplication) {
    this.$on('beforeExit', async () => {
      await app.close();
    });
  }
}
