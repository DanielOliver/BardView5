import { Module } from '@nestjs/common';
import { UserController } from './users/user.controller';
import { ConfigModule } from '@nestjs/config';
import { MikroOrmModule } from '@mikro-orm/nestjs';
import { User } from './entities/User';
import { RoleAssignment } from './entities/RoleAssignment';
import { RolePermission } from './entities/RolePermission';
import { Role } from './entities/Role';
import { RoleType } from './entities/RoleType';

@Module({
  imports: [
    ConfigModule.forRoot(),
    MikroOrmModule.forRoot({
      entities: [User, RoleAssignment, RolePermission, Role, RoleType],
      dbName: 'bardview5',
      type: 'postgresql',
      clientUrl: process.env.DATABASE_URL,
      // autoLoadEntities: true,
    }),
  ],
  controllers: [UserController],
})
export class AppModule {}
