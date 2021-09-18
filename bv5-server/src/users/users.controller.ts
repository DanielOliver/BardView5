import { Body, Controller, Get, Param, Post, Query, Res } from '@nestjs/common';
import { ApiOkResponse, ApiOperation, ApiTags } from '@nestjs/swagger';
import { ApiPrefixV1 } from '../globals';
import { UserResponse, UserCreationRequest } from './users.dto';
import { Response } from 'express';
import { UsersService } from './users.service';
import { RoleAssignmentResponse } from '../roles/roles.dto';
import { RolesService } from '../roles/roles.service';

const UserPrefix = '/users';

@Controller(ApiPrefixV1 + UserPrefix)
@ApiTags('Users')
export class UsersController {
  constructor(
    private readonly usersService: UsersService,
    private readonly rolesService: RolesService,
  ) {}

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
    return await this.usersService.getUserResponses(offset, limit);
  }

  @Get(':uid/roleassignments')
  @ApiOperation({ summary: "Fetches a user's role assignments" })
  @ApiOkResponse({
    type: RoleAssignmentResponse,
    isArray: true,
  })
  async getUserRoleAssignments(
    @Param('uid') userUid: string,
  ): Promise<RoleAssignmentResponse[]> {
    return await this.rolesService.getRoleAssignmentResponseByUser(userUid);
  }

  @Get(':uid')
  @ApiOperation({ summary: 'Fetches a user' })
  @ApiOkResponse({
    type: UserResponse,
  })
  async getUser(@Param('uid') userUid: string): Promise<UserResponse> {
    return await this.usersService.getUserResponse(userUid);
  }

  @Post()
  @ApiOperation({ summary: 'Creates a new user' })
  async createUser(
    @Body() user: UserCreationRequest,
    @Res({ passthrough: true }) res: Response,
  ) {
    const newUserUid = await this.usersService.createUser(user);
    res.set('Location', '/' + ApiPrefixV1 + UserPrefix + '/' + newUserUid);
  }
}
