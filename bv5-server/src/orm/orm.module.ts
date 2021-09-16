import { Module } from '@nestjs/common';
import { MikroOrmModule } from '@mikro-orm/nestjs';
import { User } from '../entities/User';
import { Post } from '../entities/Post';

@Module({
  imports: [
    MikroOrmModule.forRoot({
      entities: [Post, User],
      dbName: 'bardview5',
      type: 'postgresql',
      clientUrl: process.env.DATABASE_URL,
      // autoLoadEntities: true,
    }),
    MikroOrmModule.forFeature({
      entities: [User, Post],
    }),
  ],
  exports: [MikroOrmModule],
})
export class OrmModule {}
