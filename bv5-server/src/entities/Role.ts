import {
  Collection,
  Entity,
  ManyToOne,
  OneToMany,
  PrimaryKey,
  Property,
} from '@mikro-orm/core';
import { RoleType } from './RoleType';
import { User } from './User';
import { RoleAssignment } from './RoleAssignment';
import { RolePermission } from './RolePermission';

@Entity()
export class Role {
  @PrimaryKey({ columnType: 'int8' })
  id!: string;

  @Property({ length: 27, columnType: 'bpchar' })
  uid!: string;

  @ManyToOne({ entity: () => User, fieldName: 'created_by', nullable: true })
  createdBy?: User;

  @Property({
    columnType: 'timestamp',
    length: 6,
    defaultRaw: `timezone('utc'::text, now())`,
  })
  createdAt!: Date;

  @Property({
    columnType: 'timestamp',
    length: 6,
    defaultRaw: `timezone('utc'::text, now())`,
  })
  effectiveDate!: Date;

  @Property({ columnType: 'timestamp', length: 6, nullable: true })
  endDate?: Date;

  @Property()
  isActive = true;

  @Property({ columnType: 'text' })
  name!: string;

  @ManyToOne({ entity: () => RoleType, nullable: true })
  roleType?: RoleType;

  @Property({ columnType: 'text[]' })
  tags!: string[];

  @OneToMany({
    entity: () => RoleAssignment,
    mappedBy: (roleAssignment) => roleAssignment.role,
  })
  roleAssignments: Collection<RoleAssignment>;

  @OneToMany({
    entity: () => RolePermission,
    mappedBy: (rolePermission) => rolePermission.role,
  })
  rolePermissions: Collection<RolePermission>;
}
