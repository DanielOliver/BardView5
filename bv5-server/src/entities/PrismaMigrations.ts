import { Entity, PrimaryKey, Property } from '@mikro-orm/core';

@Entity({ collection: '_prisma_migrations' })
export class PrismaMigrations {

  @PrimaryKey({ length: 36 })
  id!: string;

  @Property({ length: 64 })
  checksum!: string;

  @Property({ length: 6, nullable: true })
  finishedAt?: Date;

  @Property({ length: 255 })
  migrationName!: string;

  @Property({ columnType: 'text', nullable: true })
  logs?: string;

  @Property({ length: 6, nullable: true })
  rolledBackAt?: Date;

  @Property({ length: 6, defaultRaw: `now()` })
  startedAt!: Date;

  @Property()
  appliedStepsCount: number = 0;

}
