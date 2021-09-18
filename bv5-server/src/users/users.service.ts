import { Injectable } from '@nestjs/common';
import { EntityManager, EntityRepository } from '@mikro-orm/postgresql';
import { UserCreationRequest, UserResponse } from './users.dto';
import KSUID from 'ksuid';
import { User } from '../entities/User';
import { InjectRepository } from '@mikro-orm/nestjs';
import { QueryOrder, Reference } from '@mikro-orm/core';

@Injectable()
export class UsersService {
  constructor(
    private readonly em: EntityManager,
    @InjectRepository(User)
    private readonly userRepository: EntityRepository<User>,
  ) {}

  async getUserResponse(userUid: string): Promise<UserResponse> {
    return new UserResponse(
      await this.userRepository.findOne({ uid: userUid }),
    );
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
    const dbUser = new User();
    dbUser.name = newUser.name;
    dbUser.email = newUser.email;
    dbUser.uid = userKsuid.string;
    dbUser.isActive = true;
    dbUser.tags = newUser.tags ?? [];
    if (optional?.creatingUserId) {
      dbUser.createdBy = Reference.create(
        this.em.getReference(User, optional.creatingUserId),
      );
    }
    await this.em.persistAndFlush([dbUser]);
    return dbUser.uid;
  }
}
