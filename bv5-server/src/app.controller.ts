import { Controller, Get, Param } from '@nestjs/common';
import { UserService } from './user.service';
import { EntityManager } from '@mikro-orm/postgresql';
import { User } from './entities/User';
import { ApiOperation, ApiTags } from '@nestjs/swagger';
import { ApiPrefixV1 } from './globals';
import { UserModel } from './api-models/usermodel.dto';

@Controller(ApiPrefixV1 + '/user')
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
