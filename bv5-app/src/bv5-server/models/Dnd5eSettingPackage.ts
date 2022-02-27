/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { Dnd5eMonster } from './Dnd5eMonster';
import type { Dnd5eSetting } from './Dnd5eSetting';

export type Dnd5eSettingPackage = {
    setting: Dnd5eSetting;
    monsters: Array<Dnd5eMonster>;
}