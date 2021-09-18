import { Inject, Injectable } from '@nestjs/common';
import { EntityManager, EntityRepository } from '@mikro-orm/postgresql';
import { UserCreationRequest, UserResponse } from './users.dto';
import KSUID from 'ksuid';
import { User } from '../entities/User';
import { InjectRepository } from '@mikro-orm/nestjs';
import { QueryOrder, Reference } from '@mikro-orm/core';
import { PG_CONNECTION } from '../const';
import { Pool } from 'pg';
import {
  findUserByUid,
  insertNewUser,
  stringArray,
} from '../queries/user.queries';

@Injectable()
export class UsersService {
  constructor(
    private readonly em: EntityManager,
    @InjectRepository(User)
    private readonly userRepository: EntityRepository<User>,
    @Inject(PG_CONNECTION) private readonly pg: Pool,
  ) {}

  async getUserResponse(userUid: string): Promise<UserResponse> {
    const user = await findUserByUid.run(
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
      await this.userRepository.find(
        {},
        {
          orderBy: { id: QueryOrder.ASC },
          offset: offset,
          limit: limit,
        },
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
    let tags: Array<string>;
    const insertedIds = await insertNewUser.run(
      {
        users: [
          {
            uid: userKsuid.string,
            name: newUser.name,
            email: newUser.email,
            tags: newUser.tags,
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
