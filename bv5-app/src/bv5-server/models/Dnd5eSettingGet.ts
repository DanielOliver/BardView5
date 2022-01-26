/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { Created } from './Created';
import type { Dnd5eSetting } from './Dnd5eSetting';

export type Dnd5eSettingGet = (Dnd5eSetting & {
dnd5eSettingId: string;
created: Created;
version: number;
});