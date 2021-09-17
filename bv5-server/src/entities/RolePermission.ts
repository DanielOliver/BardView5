import { Entity, ManyToOne, PrimaryKey, Property } from '@mikro-orm/core';
import { Role } from './Role';

@Entity()
export class RolePermission {

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

  @Property({ columnType: 'text[]' })
  actions!: string[];

  @Property({ columnType: 'text' })
  subject!: string;

  @Property({ nullable: true })
  conditions?: object;

  @Property({ columnType: 'text[]', nullable: true })
  fields?: string[];

}
