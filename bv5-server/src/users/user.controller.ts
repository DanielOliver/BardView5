import { Body, Controller, Get, Param, Post, Res } from '@nestjs/common';
import { EntityManager } from '@mikro-orm/postgresql';
import { User } from '../entities/User';
import { ApiOkResponse, ApiOperation, ApiTags } from '@nestjs/swagger';
import { ApiPrefixV1 } from '../globals';
import { UserResponse, UserCreationRequest } from './user.dto';
import { Response } from 'express';

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

  @Get(':id/roleassignments')
  @ApiOperation({ summary: 'Fetches all roles assigned to user' })
  @ApiOkResponse({
    isArray: true,
    type: UserResponse,
  })
  async getUserRoleAssignments(@Param('id') userId: string) {
    return (
      await this.em.findOne(
        User,
        { id: userId },
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
    const newUser = new User();
    newUser.name = user.name;
    newUser.email = user.email;
    await this.em.persistAndFlush([newUser]);
    res.set('Location', '/' + ApiPrefixV1 + UserPrefix + '/' + newUser.id);
  }
}
