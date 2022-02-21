/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { UserTags } from './UserTags';

export type Dnd5eMonster = {
    dnd5eSettingId: string;
    name: string;
    description?: string;
    sizeCategory?: string;
    armorClass?: number;
    hitPoints?: number;
    legendary?: boolean;
    unique?: boolean;
    monsterType?: string;
    alignment?: string;
    challengeRating?: string;
    userTags: UserTags;
    languages?: Array<string>;
    environments?: Array<string>;
    sources?: Array<string>;
}