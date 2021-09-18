import { Injectable } from '@nestjs/common';
import { EntityManager, EntityRepository } from '@mikro-orm/postgresql';
import KSUID from 'ksuid';
import { InjectRepository } from '@mikro-orm/nestjs';
import { QueryOrder, Reference } from '@mikro-orm/core';
import { RoleAssignment } from '../entities/RoleAssignment';
import { Role } from '../entities/Role';
import { RoleAssignmentResponse } from './roles.dto';
import { User } from '../entities/User';

@Injectable()
export class RolesService {
  constructor(
    private readonly em: EntityManager,
    @InjectRepository(RoleAssignment)
    private readonly roleAssignmentRepository: EntityRepository<RoleAssignment>,
    @InjectRepository(User)
    private readonly userRepository: EntityRepository<User>,
  ) {}

  async getRoleAssignmentResponseByUid(
    roleAssignmentUid: string,
  ): Promise<RoleAssignmentResponse> {
    return new RoleAssignmentResponse(
      await this.roleAssignmentRepository.findOne({ uid: roleAssignmentUid }),
    );
  }

  async getRoleAssignmentResponseByUser(
    userUid: string,
  ): Promise<RoleAssignmentResponse[]> {
    // const user = await this.userRepository.findOne({ uid: userUid });
    const qb = this.em.createQueryBuilder(RoleAssignment);
    const knex = qb.getKnexQuery();
    const results = await knex.select('*').where({ user: { uid: userUid } });
    const roleAssignments = results.map((roleAssignment) =>
      this.em.map(RoleAssignment, roleAssignment),
    );
    // or use EntityRepository.map()const repo = orm.em.getRepository(User);const users = results.map(user => repo.map(user));

    // const roleAssignments = await this.roleAssignmentRepository.find(
    //   {
    //     user: {
    //       uid: userUid,
    //     },
    //   },
    //   {
    //     populate: {
    //       role: true,
    //     },
    //   },
    // );

    return roleAssignments.map(
      (x) =>
        new RoleAssignmentResponse({
          roleUid: x.role.uid,
          userUid: userUid,
          uid: x.uid,
          isActive: x.isActive,
          endDate: x.endDate,
          effectiveDate: x.effectiveDate,
          createdAt: x.createdAt,
        }),
    );
  }
}
