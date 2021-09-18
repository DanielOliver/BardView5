import {
  Body,
  Controller,
  Get,
  NotFoundException,
  Param,
  Post,
  Query,
  Res,
} from '@nestjs/common';
import { ApiOkResponse, ApiOperation, ApiTags } from '@nestjs/swagger';
import { ApiPrefixV1 } from '../globals';
import { UserResponse, UserCreationRequest } from './users.dto';
import { Response } from 'express';
import { UsersService } from './users.service';
import { SessionScope } from '../session/session';

const UserPrefix = '/users';

@Controller(ApiPrefixV1 + UserPrefix)
@ApiTags('Users')
export class UsersController {
  constructor(private readonly usersService: UsersService) {}

  @Get()
  @ApiOperation({ summary: 'Fetches some users' })
  @ApiOkResponse({
    isArray: true,
    type: UserResponse,
  })
  async getUsers(
    @Query('offset') offset: number,
    @Query('limit') limit: number,
  ): Promise<UserResponse[]> {
    return await this.usersService.getUserResponses(
      offset,
      limit,
      new SessionScope({
        userId: '3',
      }),
    );
  }

  @Get(':uid')
  @ApiOperation({ summary: 'Fetches a user' })
  @ApiOkResponse({
    type: UserResponse,
  })
  async getUser(@Param('uid') userUid: string): Promise<UserResponse> {
    const result = await this.usersService.getUserResponse(
      userUid,
      new SessionScope({
        userId: '3',
      }),
    );
    if (result) {
      return result;
    }
    throw new NotFoundException();
  }

  @Post()
  @ApiOperation({ summary: 'Creates a new user' })
  async createUser(
    @Body() user: UserCreationRequest,
    @Res({ passthrough: true }) res: Response,
  ) {
    const newUserUid = await this.usersService.createUser(
      user,
      new SessionScope({
        userId: null,
      }),
    );
    res.set('Location', '/' + ApiPrefixV1 + UserPrefix + '/' + newUserUid);
  }
}
