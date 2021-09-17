import { Body, Controller, Get, Param, Post, Res } from '@nestjs/common';
import { EntityManager } from '@mikro-orm/postgresql';
import { User } from '../entities/User';
import { ApiOkResponse, ApiOperation, ApiTags } from '@nestjs/swagger';
import { ApiPrefixV1 } from '../globals';
import { UserResponse, UserCreationRequest } from './user.dto';
import { Response } from 'express';
import KSUID from 'ksuid';
import { RoleAssignment } from '../entities/RoleAssignment';

const UserPrefix = '/users';

@Controller(ApiPrefixV1 + UserPrefix)
@ApiTags('Users')
export class UserController {
  constructor(private readonly em: EntityManager) {}

  @Get()
  @ApiOperation({ summary: 'Fetches all users' })
  @ApiOkResponse({
    isArray: true,
    type: UserResponse,
  })
  async getUsers(): Promise<UserResponse[]> {
    return (await this.em.find(User, {})).map((user) => new UserResponse(user));
  }

  @Get(':uid')
  @ApiOperation({ summary: 'Fetches a user' })
  @ApiOkResponse({
    type: UserResponse,
  })
  async getUser(@Param('uid') userId: string) {
    return new UserResponse(await this.em.findOne(User, { uid: userId }));
  }

  @Get(':uid/roleassignments')
  @ApiOperation({ summary: 'Fetches all roles assigned to user' })
  @ApiOkResponse({
    isArray: true,
    type: RoleAssignment,
  })
  async getUserRoleAssignments(@Param('uid') userId: string) {
    return (
      await this.em.findOne(
        User,
        { uid: userId },
        {
          populate: {
            roleAssignments: true,
          },
        },
      )
    ).roleAssignments;
  }

  @Post()
  @ApiOperation({ summary: 'Creates a new user' })
  async createUser(
    @Body() user: UserCreationRequest,
    @Res({ passthrough: true }) res: Response,
  ) {
    const userKsuid = await KSUID.random();
    const newUser = new User();
    newUser.name = user.name;
    newUser.email = user.email;
    newUser.uid = userKsuid.string;
    newUser.isActive = true;
    newUser.tags = user.tags ?? [];

    await this.em.persistAndFlush([newUser]);
    res.set('Location', '/' + ApiPrefixV1 + UserPrefix + '/' + newUser.uid);
  }
}
