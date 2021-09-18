import { Module } from '@nestjs/common';
import { UsersController } from './users/users.controller';
import { ConfigModule } from '@nestjs/config';
import { MikroOrmModule } from '@mikro-orm/nestjs';
import { User } from './entities/User';
import { RoleAssignment } from './entities/RoleAssignment';
import { RolePermission } from './entities/RolePermission';
import { Role } from './entities/Role';
import { RoleType } from './entities/RoleType';
import { RoleAction } from './entities/RoleAction';
import { RoleSubject } from './entities/RoleSubject';
import { UsersService } from './users/users.service';
import { RolesService } from './roles/roles.service';
import { Pool } from 'pg';
import { PG_CONNECTION } from './const';

@Module({
  imports: [
    ConfigModule.forRoot(),
    MikroOrmModule.forRoot({
      entities: [
        User,
        RoleAssignment,
        RolePermission,
        Role,
        RoleType,
        RoleAction,
        RoleSubject,
      ],
      dbName: 'bardview5',
      type: 'postgresql',
      clientUrl: process.env.DATABASE_URL,
    }),
    MikroOrmModule.forFeature([User, RoleAssignment, RolePermission, Role]),
  ],
  controllers: [UsersController],
  providers: [
    UsersService,
    RolesService,
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
