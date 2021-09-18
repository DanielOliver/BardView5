import { Inject, Injectable } from '@nestjs/common';
import { UserCreationRequest, UserResponse } from './users.dto';
import KSUID from 'ksuid';
import { PG_CONNECTION } from '../const';
import { Pool } from 'pg';
import {
  sqlUsersFindByUid,
  sqlUsersGetPaginated,
  sqlUsersUpsert,
} from '../queries/user.queries';
import { SessionScope } from '../session/session';
import {
  sqlRolesInsertUserRole,
  sqlRolesLinkUserToGlobalUserRole,
  sqlRolesLinkUserToRole,
} from '../queries/role.queries';

@Injectable()
export class UsersService {
  constructor(@Inject(PG_CONNECTION) private readonly pg: Pool) {}

  async getUserResponse(
    userUid: string,
    scope: SessionScope,
  ): Promise<UserResponse> {
    const user = await sqlUsersFindByUid.run(
      {
        uid: userUid,
        session_id: scope.userId,
      },
      this.pg,
    );
    if (user.length > 0) {
      return new UserResponse(user[0]);
    }
    return null;
  }

  async getUserResponses(
    offset: number,
    limit: number,
    scope: SessionScope,
  ): Promise<UserResponse[]> {
    return (
      await sqlUsersGetPaginated.run(
        {
          offset: offset.toString(),
          limit: limit.toString(),
          session_id: scope.userId,
        },
        this.pg,
      )
    ).map((x) => new UserResponse(x));
  }

  async createUser(
    newUser: UserCreationRequest,
    scope: SessionScope,
  ): Promise<string> {
    const userKsuid = await KSUID.random();
    const client = await this.pg.connect();
    try {
      await client.query('BEGIN');

      const insertedUserIds = await sqlUsersUpsert.run(
        {
          users: [
            {
              uid: userKsuid.string,
              name: newUser.name,
              email: newUser.email,
              tags: newUser.tags,
              created_by: scope.userId,
            },
          ],
        },
        client,
      );
      if (insertedUserIds.length == 0) {
        await client.query('ROLLBACK');
        return null;
      }

      const userRole = await sqlRolesInsertUserRole.run(
        {
          uid: insertedUserIds[0].uid,
          created_by: scope.userId,
          name: newUser.email,
          tags: ['User'],
        },
        client,
      );

      if (userRole.length == 0) {
        await client.query('ROLLBACK');
        return null;
      }

      const userRoleAssignment = await sqlRolesLinkUserToRole.run(
        {
          uid: insertedUserIds[0].uid,
          created_by: scope.userId,
          role_id: userRole[0].id,
          user_id: insertedUserIds[0].id,
          tags: ['User'],
        },
        client,
      );

      const globalRoleUserKsuid = await KSUID.random();
      const userGlobalRoleAssignment =
        await sqlRolesLinkUserToGlobalUserRole.run(
          {
            uid: globalRoleUserKsuid.string,
            created_by: scope.userId,
            user_id: insertedUserIds[0].id,
            tags: ['User', 'Global'],
          },
          client,
        );

      await client.query('COMMIT');
      return insertedUserIds[0].uid;
    } catch (e) {
      await client.query('ROLLBACK');
      throw e;
    } finally {
      client.release();
    }
    return null;
  }
}
