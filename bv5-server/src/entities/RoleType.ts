import { Entity, PrimaryKey, Property } from '@mikro-orm/core';

@Entity()
export class RoleType {

  @PrimaryKey({ columnType: 'int8' })
  id!: string;

  @Property({ columnType: 'text' })
  name!: string;

  @Property()
  multipleAllowed!: boolean;

}
