import { Body, Controller, Get, Param, Post } from '@nestjs/common';
import { EntityManager } from '@mikro-orm/postgresql';
import { User } from '../entities/User';
import { ApiOkResponse, ApiOperation, ApiTags } from '@nestjs/swagger';
import { ApiPrefixV1 } from '../globals';
import { UserResponse, UserCreationRequest } from './user.dto';

@Controller(ApiPrefixV1 + '/users')
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

  @Post()
  @ApiOperation({ summary: 'Creates a new user' })
  createUser(@Body() user: UserCreationRequest) {
    return;
  }
}
