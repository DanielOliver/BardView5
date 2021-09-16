import { Module } from '@nestjs/common';
import { UserController } from './users/user.controller';
import { ConfigModule } from '@nestjs/config';
import { OrmModule } from './orm/orm.module';

@Module({
  imports: [ConfigModule.forRoot(), OrmModule],
  controllers: [UserController],
})
export class AppModule {}
