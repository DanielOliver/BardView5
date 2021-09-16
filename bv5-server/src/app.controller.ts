import { Controller, Get, Param } from '@nestjs/common';
import { UserService } from './user.service';
import { EntityManager } from '@mikro-orm/postgresql';
import { User } from './entities/User';
import {
  ApiOkResponse,
  ApiOperation,
  ApiResponse,
  ApiTags,
} from '@nestjs/swagger';
import { ApiPrefix } from './globals';
import { UserModel } from './api-models/usermodel.dto';
import { Loaded } from '@mikro-orm/core';

@Controller(ApiPrefix)
@ApiTags('Users')
export class AppController {
  constructor(
    private readonly userService: UserService,
    private readonly em: EntityManager,
  ) {}

  @Get()
  @ApiOperation({ summary: 'Fetches all users' })
  async getUsers(): Promise<UserModel[]> {
    return (await this.em.find(User, {})).map((user) => new UserModel(user));
  }

  @Get(':email')
  create(@Param('email') email: string) {
    this.userService.createUser({ email: email, name: 'name2' });
  }
}
