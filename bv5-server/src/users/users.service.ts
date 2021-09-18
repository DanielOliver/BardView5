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

@Injectable()
export class UsersService {
  constructor(@Inject(PG_CONNECTION) private readonly pg: Pool) {}

  async getUserResponse(userUid: string): Promise<UserResponse> {
    const user = await sqlUsersFindByUid.run(
      {
        uid: userUid,
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
  ): Promise<UserResponse[]> {
    return (
      await sqlUsersGetPaginated.run(
        { offset: offset.toString(), limit: limit.toString() },
        this.pg,
      )
    ).map((x) => new UserResponse(x));
  }

  async createUser(
    newUser: UserCreationRequest,
    optional?: {
      creatingUserId: string | null;
    },
  ): Promise<string> {
    const userKsuid = await KSUID.random();
    const insertedIds = await sqlUsersUpsert.run(
      {
        users: [
          {
            uid: userKsuid.string,
            name: newUser.name,
            email: newUser.email,
            tags: newUser.tags,
            created_by: null,
          },
        ],
      },
      this.pg,
    );
    if (insertedIds.length < 1) {
      return null;
    }
    return insertedIds[0].uid;
  }
}
