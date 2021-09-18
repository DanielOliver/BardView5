import { Injectable } from '@nestjs/common';
import { RoleAssignmentResponse } from './roles.dto';

@Injectable()
export class RolesService {
  constructor() {}

  async getRoleAssignmentResponseByUid(
    roleAssignmentUid: string,
  ): Promise<RoleAssignmentResponse> {
    throw new Error('FAIL!');
  }

  async getRoleAssignmentResponseByUser(
    userUid: string,
  ): Promise<RoleAssignmentResponse[]> {
    throw new Error('FAIL!');
  }
}
