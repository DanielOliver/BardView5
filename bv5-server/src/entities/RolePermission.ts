import { Entity, ManyToOne, PrimaryKey, Property } from '@mikro-orm/core';
import { Role } from './Role';
import { RoleAction } from './RoleAction';
import { RoleSubject } from './RoleSubject';
import { User } from './User';

@Entity()
export class RolePermission {
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

  @ManyToOne({ entity: () => Role })
  role!: Role;

  @ManyToOne({ entity: () => RoleAction, fieldName: 'action' })
  action!: RoleAction;

  @ManyToOne({ entity: () => RoleSubject, fieldName: 'subject' })
  subject!: RoleSubject;

  @Property()
  conditions!: object;

  @Property({ columnType: 'text[]', nullable: true })
  fields?: string[];
}
