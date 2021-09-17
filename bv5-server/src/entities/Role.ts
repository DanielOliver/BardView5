import { Entity, ManyToOne, PrimaryKey, Property } from '@mikro-orm/core';
import { RoleType } from './RoleType';

@Entity()
export class Role {

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

  @Property({ columnType: 'text' })
  name!: string;

  @ManyToOne({ entity: () => RoleType, nullable: true })
  roleType?: RoleType;

}
