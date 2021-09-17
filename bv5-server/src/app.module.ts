import { Module } from '@nestjs/common';
import { UserController } from './users/user.controller';
import { ConfigModule } from '@nestjs/config';
import { MikroOrmModule } from '@mikro-orm/nestjs';
import { User } from './entities/User';

@Module({
  imports: [
    ConfigModule.forRoot(),
    MikroOrmModule.forRoot({
      entities: [User],
      dbName: 'bardview5',
      type: 'postgresql',
      clientUrl: process.env.DATABASE_URL,
      // autoLoadEntities: true,
    }),
  ],
  controllers: [UserController],
})
export class AppModule {}
