import { Entity, ManyToOne, PrimaryKey, Property } from '@mikro-orm/core';
import { Role } from './Role';
import { User } from './User';

@Entity()
export class RoleAssignment {

  @PrimaryKey({ columnType: 'int8' })
  id!: string;

  @Property({ columnType: 'timestamp', length: 6, defaultRaw: `timezone('utc'::text, now())` })
  createdAt!: Date;

  @Property({ columnType: 'timestamp', length: 6, defaultRaw: `timezone('utc'::text, now())` })
  effectiveDate!: Date;

  @Property({ columnType: 'timestamp', length: 6, nullable: true })
  endDate?: Date;

  @Property()
  isActive: boolean = true;

  @ManyToOne({ entity: () => Role })
  role!: Role;

  @ManyToOne({ entity: () => User })
  user!: User;

}
