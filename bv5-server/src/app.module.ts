import { Module } from '@nestjs/common';
import { UsersController } from './users/users.controller';
import { ConfigModule } from '@nestjs/config';
import { UsersService } from './users/users.service';
import { Pool } from 'pg';
import { PG_CONNECTION } from './const';

@Module({
  imports: [ConfigModule.forRoot()],
  controllers: [UsersController],
  providers: [
    UsersService,
    {
      provide: PG_CONNECTION,
      useValue: new Pool({
        user: process.env.postgres_user,
        host: process.env.postgres_host,
        database: process.env.postgres_database,
        password: process.env.postgres_password,
        port: 5432,
      }),
    },
  ],
})
export class AppModule {}
