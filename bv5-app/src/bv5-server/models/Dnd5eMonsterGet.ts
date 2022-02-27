/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { Created } from './Created';
import type { Dnd5eMonster } from './Dnd5eMonster';

export type Dnd5eMonsterGet = (Dnd5eMonster & {
dnd5eMonsterId: string;
dnd5eSettingId: string;
created: Created;
version: number;
});