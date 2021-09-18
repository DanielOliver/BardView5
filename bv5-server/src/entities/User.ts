import {
  Collection,
  Entity,
  IdentifiedReference,
  ManyToOne,
  OneToMany,
  PrimaryKey,
  Property,
  Unique,
} from '@mikro-orm/core';
import { RoleAssignment } from './RoleAssignment';

@Entity()
export class User {
  @PrimaryKey({ columnType: 'int8' })
  id!: string;

  @Property({ length: 27, columnType: 'bpchar' })
  uid!: string;

  @ManyToOne({
    entity: () => User,
    fieldName: 'created_by',
    nullable: true,
    wrappedReference: true,
  })
  createdBy?: IdentifiedReference<User>;

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

  @Unique({ name: 'user_email_uindex' })
  @Property({ columnType: 'text' })
  email!: string;

  @Property({ columnType: 'text' })
  name!: string;

  @Property()
  tags!: string[];

  @OneToMany({
    entity: () => RoleAssignment,
    mappedBy: (roleAssignment) => roleAssignment.user,
  })
  roleAssignments: Collection<RoleAssignment>;
}
