import { Entity, PrimaryKey, Property, Unique } from '@mikro-orm/core';

@Entity()
export class RoleType {
  @PrimaryKey({ columnType: 'int8' })
  id!: string;

  @Property({
    columnType: 'timestamp',
    length: 6,
    defaultRaw: `timezone('utc'::text, now())`,
  })
  createdAt!: Date;

  @Unique({ name: 'roletype_name_uindex' })
  @Property({ columnType: 'text' })
  name!: string;

  @Property()
  multipleAssignmentsAllowed!: boolean;

  @Property()
  systemManaged!: boolean;
}
