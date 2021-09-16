import { Controller, Get, Param } from '@nestjs/common';
import { UserService } from './user.service';
import { User } from '@prisma/client';

@Controller()
export class AppController {
  constructor(private readonly userService: UserService) {}

  @Get()
  getHello(): Promise<User[]> {
    return this.userService.users({});
  }

  @Get(':email')
  create(@Param('email') email: string) {
    this.userService.createUser({ email: email, name: 'name2' });
  }
}
